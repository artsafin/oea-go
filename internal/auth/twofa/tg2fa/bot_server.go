package tg2fa

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
	"oea-go/internal/common/config"
	"sync"
)

type accountChannel chan config.Account

type botServer struct {
	token            string
	bot              *tgbotapi.BotAPI
	updChan          tgbotapi.UpdatesChannel
	sessions         map[config.Username]*authSession
	unregisteredChan accountChannel
	shutdownChan     chan bool
	mut              sync.Mutex
	logger           *zap.SugaredLogger
}

func newBotServer(token string, logger *zap.SugaredLogger) *botServer {
	return &botServer{
		bot:          nil,
		updChan:      nil,
		token:        token,
		sessions:     make(map[config.Username]*authSession),
		logger:       logger,
		shutdownChan: make(chan bool),
	}
}

func (b *botServer) resurrect() (*tgbotapi.BotAPI, error) {
	b.logger.Infof("resurrect: start")
	defer b.logger.Infof("resurrect: finish")

	if b.updChan != nil {
		b.logger.Infof("\t↳resurrect: botsrv is alive, no need for resurrection")
		return b.bot, nil
	}

	var err error
	b.bot, err = tgbotapi.NewBotAPI(b.token)
	if err != nil {
		return nil, err
	}

	b.bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	b.updChan, err = b.bot.GetUpdatesChan(updateConfig)

	if err != nil {
		b.bot.Client.CloseIdleConnections()
		b.bot = nil
		return nil, err
	}
	b.unregisteredChan = make(accountChannel)

	go b.CleanupExpiredSessions()
	go b.ListenUpdates()

	return b.bot, nil
}

func (b *botServer) CleanupExpiredSessions() {
	b.logger.Infof("cleanup: beginning expired sessions cleanup routine")
	defer b.logger.Infof("cleanup: exiting")

	for {
		select {
		case expiredAcc := <-b.unregisteredChan:
			b.unregisterSession(expiredAcc)

			b.mut.Lock()
			if len(b.sessions) == 0 {
				b.logger.Infof("cleanup: no active sessions left, stopping botsrv")
				b.stop()
				b.mut.Unlock()
				return
			}
			b.mut.Unlock()
		}
	}
}

func (b *botServer) ListenUpdates() {
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

			b.mut.Lock()

			sess, sessFound := b.sessions[userName]
			if !sessFound {
				continue
			}

			userReply := userReplyMeta{
				chatID:   update.Message.Chat.ID,
				userID:   update.Message.From.ID,
				username: userName,
			}

			switch {
			case update.Message.Text == startText:
				sess.startChan <- userReply
			case update.Message.Text == allowText:
				sess.allowButtonChan <- userReply
			case update.Message.Text == denyText:
				sess.declineButtonChan <- userReply
			}

			b.mut.Unlock()
		}
	}
}

func (b *botServer) stop() {
	b.logger.Infof("stop: stopping botsrv")
	defer b.logger.Infof("stop: stopped botsrv")

	b.shutdownChan <- true
	b.bot.StopReceivingUpdates()
	b.updChan.Clear()
	b.bot.Client.CloseIdleConnections()
	b.bot = nil
	b.updChan = nil
	close(b.unregisteredChan)
}

func (b *botServer) registerSession(sess *authSession) error {
	b.mut.Lock()
	defer b.mut.Unlock()

	go sess.expireAfterTimeout(b.unregisteredChan)

	b.sessions[sess.account.ExternalUsername] = sess

	return nil
}

func (b *botServer) unregisterSession(expiredAcc config.Account) {
	b.mut.Lock()
	defer b.mut.Unlock()

	b.logger.Infof("unregister: %+v has expired, cleaning up...", expiredAcc)

	u := expiredAcc.ExternalUsername
	_, found := b.sessions[u]
	if !found {
		b.logger.Infof("unregister: %v not found in sessions", expiredAcc.ExternalUsername)
		return
	}
	delete(b.sessions, u)
}
