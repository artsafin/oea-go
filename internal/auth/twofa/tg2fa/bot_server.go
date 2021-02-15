package tg2fa

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
	"oea-go/internal/common/config"
	"sync"
)

type accountChannel chan config.Account

type botServer struct {
	unregisterChan accountChannel
	shutdownChan   chan bool

	token    string
	bot      *tgbotapi.BotAPI
	updChan  tgbotapi.UpdatesChannel
	sessions map[config.Username]*authSession
	mut      sync.Mutex
	logger   *zap.SugaredLogger
}

func newBotServer(token string, logger *zap.SugaredLogger) *botServer {
	return &botServer{
		token:    token,
		sessions: make(map[config.Username]*authSession),
		logger:   logger,
	}
}

func (b *botServer) resurrect() (*tgbotapi.BotAPI, error) {
	b.logger.Infof("resurrect: start")
	defer b.logger.Infof("resurrect: finish")

	if b.updChan != nil {
		b.logger.Infof("resurrect: botsrv is alive, no need for resurrection")
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
	b.shutdownChan = make(chan bool)
	b.unregisterChan = make(accountChannel)

	go b.CleanupExpiredSessions()
	go b.ListenUpdates()

	return b.bot, nil
}

func (b *botServer) CleanupExpiredSessions() {
	b.logger.Infof("cleanup: beginning expired sessions cleanup routine")
	defer b.logger.Infof("cleanup: exiting")

	for {
		select {
		case unregAcc := <-b.unregisterChan:
			b.logger.Infof("cleanup: start check for %v", unregAcc)
			b.unregisterSession(unregAcc)

			wasShutDown := b.checkIfNeedToShutdown()

			if wasShutDown {
				return
			}
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

			b.processUpdateForUsername(update.Message, userName)
		}
	}
}

func (b *botServer) processUpdateForUsername(msg *tgbotapi.Message, userName config.Username) {
	b.logger.Debugf("update: processing update %v", userName)
	defer b.logger.Debugf("update: finished processing update %v", userName)

	b.mut.Lock()
	defer b.mut.Unlock()

	sess, sessFound := b.sessions[userName]
	if !sessFound {
		b.logger.Debugf("update: no session for %v", userName)
		return
	}

	if !sess.isMessageAcceptable(msg) {
		b.logger.Debugf("update: skipping past message %v for %v", msg.MessageID, userName)
		return
	}

	userReply := userReplyMeta{
		chatID:   msg.Chat.ID,
		userID:   msg.From.ID,
		username: userName,
	}

	switch {
	case msg.Text == startText:
		sess.startChan <- userReply
	case msg.Text == allowText:
		sess.allowButtonChan <- userReply
	case msg.Text == denyText:
		sess.declineButtonChan <- userReply
	}
}

func (b *botServer) checkIfNeedToShutdown() bool {
	b.mut.Lock()
	defer b.mut.Unlock()

	if len(b.sessions) != 0 {
		b.logger.Debugf("shutdown: %v session(s) still active, not stopping", len(b.sessions))
		return false
	}

	b.logger.Infof("shutdown: no active sessions left, stopping botsrv")
	defer b.logger.Infof("shutdown: stopped botsrv")

	b.shutdownChan <- true
	b.bot.StopReceivingUpdates()
	b.updChan.Clear()
	b.bot.Client.CloseIdleConnections()
	b.bot = nil
	b.updChan = nil
	close(b.unregisterChan)
	close(b.shutdownChan)

	return true
}

func (b *botServer) registerSession(sess *authSession) error {
	b.mut.Lock()
	defer b.mut.Unlock()

	b.logger.Debugf("registerSession: start after lock!")
	defer b.logger.Debugf("registerSession: finish")

	go sess.WaitForShutdown(b.unregisterChan)

	b.sessions[sess.account.ExternalUsername] = sess

	return nil
}

func (b *botServer) unregisterSession(expiredAcc config.Account) {
	b.mut.Lock()
	defer b.mut.Unlock()

	b.logger.Infof("unregister: %+v has expired, cleaning up...", expiredAcc)

	u := expiredAcc.ExternalUsername
	sess, found := b.sessions[u]
	if !found {
		b.logger.Infof("unregister: %v not found in sessions", expiredAcc.ExternalUsername)
		return
	}

	sess.Close()

	delete(b.sessions, u)
	b.logger.Debugf("unregister: %+v cleanup finish", expiredAcc)
}
