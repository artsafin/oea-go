package twofa

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
	"oea-go/internal/common/config"
	"sync"
)

type expirationsChannel chan config.Account

type tgBot struct {
	token           string
	bot             *tgbotapi.BotAPI
	updChan         tgbotapi.UpdatesChannel
	sessions        map[config.Username]*tgAuthSession
	expirationsChan expirationsChannel
	shutdownChan    chan bool
	mut             sync.Mutex
	logger          *zap.SugaredLogger
}

func newBotSupervisor(token string, logger *zap.SugaredLogger) *tgBot {
	return &tgBot{
		bot:          nil,
		updChan:      nil,
		token:        token,
		sessions:     make(map[config.Username]*tgAuthSession),
		logger:       logger,
		shutdownChan: make(chan bool),
	}
}

func (b *tgBot) resurrect() error {
	b.logger.Infof("resurrect: resurrecting bot service")
	defer b.logger.Infof("resurrect: finish")

	if b.updChan != nil {
		b.logger.Infof("\t↳resurrect: update channel is alive, no need for resurrection")
		return nil
	}

	var err error
	b.bot, err = tgbotapi.NewBotAPI(b.token)
	if err != nil {
		return err
	}

	b.bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	b.updChan, err = b.bot.GetUpdatesChan(updateConfig)

	if err != nil {
		b.bot.Client.CloseIdleConnections()
		b.bot = nil
		return err
	}
	b.expirationsChan = make(expirationsChannel)

	go b.CleanupExpiredSessions()
	go b.ListenUpdates()

	return nil
}

func (b *tgBot) cleanupAccount(expiredAcc config.Account) (shouldTerminate bool) {
	b.mut.Lock()
	defer b.mut.Unlock()

	b.logger.Infof("cleanup: %+v has expired, cleaning up...", expiredAcc)

	u := expiredAcc.ExternalUsername
	sess, found := b.sessions[u]
	if !found {
		b.logger.Infof("cleanup: %v not found in sessions", expiredAcc.ExternalUsername)
		return
	}
	sess.terminate()
	delete(b.sessions, u)

	if len(b.sessions) == 0 {
		b.logger.Infof("cleanup: no active sessions left, stopping bot")
		b.stop()
		return true
	}

	return false
}

func (b *tgBot) CleanupExpiredSessions() {
	b.logger.Infof("cleanup: beginning expired sessions cleanup routine")
	defer b.logger.Infof("cleanup: exiting")
loop:
	for {
		select {
		case expiredAcc := <-b.expirationsChan:
			shouldTerminate := b.cleanupAccount(expiredAcc)
			if shouldTerminate {
				break loop
			}
		}
	}
}

func (b *tgBot) ListenUpdates() {
	b.logger.Infof("listen: begin listening")
	defer b.logger.Infof("listen: end listening")

	for {
		select {
		case <-b.shutdownChan:
			return
		case update := <-b.updChan:
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			userName := config.Username(update.Message.From.UserName)
			if userName == "" {
				continue
			}

			userReply := tgUserReplyMeta{
				chatID:   update.Message.Chat.ID,
				userID:   update.Message.From.ID,
				username: userName,
			}

			b.mut.Lock()

			switch {
			case update.Message.Text == startText:
				b.sessions[userName].startChan <- userReply
			case update.Message.Text == allowText:
				b.sessions[userName].allowButtonChan <- userReply
			case update.Message.Text == denyText:
				b.sessions[userName].declineButtonChan <- userReply
			}

			b.mut.Unlock()
		}
	}
}

func (b *tgBot) stop() {
	b.logger.Infof("stop: stopping bot")
	b.shutdownChan <- true
	b.bot.StopReceivingUpdates()
	b.updChan.Clear()
	b.bot.Client.CloseIdleConnections()
	b.bot = nil
	b.updChan = nil
	close(b.expirationsChan)

	for k, s := range b.sessions {
		b.logger.Infof("stop: terminating session %+v", k)
		s.terminate()

		delete(b.sessions, k)
	}
	b.logger.Infof("stop: stopped bot")
}

func (b *tgBot) BeginAuthSession(account config.Account) (*tgAuthSession, error) {
	b.mut.Lock()
	defer b.mut.Unlock()

	resurrectErr := b.resurrect()
	if resurrectErr != nil {
		return nil, resurrectErr
	}

	sess := newAuthSession(account, b.bot)
	b.logger.Infof("begin_auth: new session %+v", account)

	go sess.runExpirationTimer(b.expirationsChan)

	b.sessions[account.ExternalUsername] = sess

	return sess, nil
}
