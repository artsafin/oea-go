package tg2fa

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"oea-go/internal/auth/twofa"
	"oea-go/internal/common"
	"oea-go/internal/common/config"
	"oea-go/internal/db"
	"time"
)

const (
	sessionTimeout = time.Second * 40
	allowText      = "🔑 Allow"
	denyText       = "⛔ Deny"
	startText      = "/start"
	replyTimeout   = `⛔ Waited too much, authentication canceled.
Please click on the authentication link or refresh the page again to continue.`
	replyDeclined = `⛔ Authentication canceled.
Please click on the authentication link or refresh the page again to continue.`
	replySuccess = `🔑 Authentication successful.
Please wait for the automatic page redirect.
`
)

type authSession struct {
	startChan         chan userReplyMeta
	allowButtonChan   chan userReplyMeta
	declineButtonChan chan userReplyMeta
	timeoutChan       chan bool
	timerCancelChan   chan bool
	resultChan        chan twofa.Result
	shutdownChan      chan bool
	bot               *tgbotapi.BotAPI
	logger            *zap.SugaredLogger
	account           config.Account
	encKey            []byte
	info              common.AuthInfo
	startTs           time.Time
	chatID            int64
	storage           *db.Storage
}

func newAuthSession(account config.Account, info common.AuthInfo, api *tgbotapi.BotAPI, logger *zap.SugaredLogger, encKey []byte, storage *db.Storage) *authSession {
	logger.Debugf("session: new session %v", account)

	return &authSession{
		account:           account,
		info:              info,
		startChan:         make(chan userReplyMeta),
		allowButtonChan:   make(chan userReplyMeta),
		declineButtonChan: make(chan userReplyMeta),
		timeoutChan:       make(chan bool),
		timerCancelChan:   make(chan bool, 1),
		shutdownChan:      make(chan bool),
		resultChan:        make(chan twofa.Result, 1),
		bot:               api,
		logger:            logger,
		encKey:            encKey,
		startTs:           time.Now(),
		storage:           storage,
	}
}

func (s *authSession) String() string {
	return fmt.Sprintf("Session{%v on %v}", s.account, s.startTs.Format(time.RFC822Z))
}

func (s *authSession) GetExpTs() time.Time {
	return s.startTs.Add(sessionTimeout)
}

func (s *authSession) ResultChan() <-chan twofa.Result {
	return s.resultChan
}

func (s *authSession) isMessageAcceptable(message *tgbotapi.Message) bool {
	return message.Time().After(s.startTs)
}

func (s *authSession) Flow(cachedChatID int64) {
	s.logger.Debugf("auth flow: start (cached chat id %v)", cachedChatID)

	defer func() {
		s.shutdownChan <- true
		s.logger.Debugf("auth flow: finish")
	}()

	var err error
	if cachedChatID == 0 {
		s.chatID, err = s.waitForChatIdFromUserStartReply(s.account)
		storageErr := s.storage.SetChatID(s.account.Email, s.chatID)
		if storageErr != nil {
			s.logger.Warnf("auth flow: error storing chatID: %v", storageErr)
		}
	} else {
		s.chatID = cachedChatID
	}
	if err != nil {
		s.logger.Errorf("auth flow: waitForChatId: %v", err)
		s.resultChan <- twofa.Result{Err: errors.Wrap(err, "cannot obtain chat ID")}
		return
	}

	promptMsg, promptErr := s.sendPrompt(s.info)
	if promptErr != nil {
		s.logger.Errorf("auth flow: sendPrompt: %v", promptErr)
		s.resultChan <- twofa.Result{Err: errors.Wrap(promptErr, "cannot send auth prompt")}
		return
	}

	select {
	case denyReply := <-s.declineButtonChan:
		s.logger.Debugf("auth flow: got decline: %v", denyReply)
		s.sendBotReply(promptMsg, replyDeclined)
		s.resultChan <- twofa.Result{Err: errors.Errorf("user has rejected authentication")}
	case allowReply := <-s.allowButtonChan:
		s.logger.Debugf("auth flow: got allow %v", allowReply)
		if validationErr := allowReply.validate(s.chatID, s.account); validationErr != nil {
			s.resultChan <- twofa.Result{Err: validationErr}
			break
		}
		fp, fpErr := allowReply.fingerprint(s.encKey)
		if fpErr != nil {
			s.resultChan <- twofa.Result{Err: errors.New("fingerprint encryption failed")}
			break
		}
		s.logger.Debugf("Successful auth, fingerprint: %v", fp)
		s.sendBotReply(promptMsg, replySuccess)
		s.resultChan <- twofa.Result{Fingerprint: fp}
	case <-s.timeoutChan:
		s.logger.Debugf("auth flow: timeout")
		s.sendBotReply(promptMsg, replyTimeout)
		s.resultChan <- twofa.Result{Err: errors.New("auth session timeout")}
	}
}

func (s *authSession) waitForChatIdFromUserStartReply(account config.Account) (chatID int64, err error) {
	select {
	case startReply, ok := <-s.startChan:
		if !ok {
			return 0, errors.New("start channel closed")
		}
		if err = startReply.validate(0, account); err != nil {
			return 0, err
		}
		return startReply.chatID, nil
	case <-s.timeoutChan:
		return 0, errors.New("timeout waiting for start")
	}
}

func (s *authSession) waitForExpire() {
	s.logger.Debugf("sess expire: begin waiting")
	defer s.logger.Debugf("sess expire: finish")

	timer := time.NewTimer(s.GetExpTs().Sub(time.Now()))

	select {
	case <-timer.C:
		s.logger.Debugf("sess expire: timeout of session")
		s.timeoutChan <- true
	case <-s.timerCancelChan:
	}
}

func (s *authSession) WaitForShutdown(unregisterChan accountChannel) {
	s.logger.Debugf("WaitForShutdown: start")
	defer s.logger.Debugf("WaitForShutdown: finish")

	go s.waitForExpire()

	<-s.shutdownChan
	s.timerCancelChan <- true

	// Wait 5 seconds more before disposal in order for the client to be able to receive the last update on session status
	<-time.After(time.Second * 5)

	unregisterChan <- s.account
}

func (s *authSession) sendPrompt(info common.AuthInfo) (message tgbotapi.Message, err error) {
	text := fmt.Sprintf(
		"OEA authentication has been requested on <code>%v</code> from IP <code>%v</code>.\n\n"+
			"- To allow access please press <code>🔑 Allow</code>\n"+
			"- To deny authentication please press <code>⛔ Deny</code> or ignore this message",
		info.TS.Format(time.UnixDate),
		info.IP,
	)
	msgConfig := tgbotapi.NewMessage(s.chatID, text)
	msgConfig.ParseMode = "HTML"
	kb := tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(allowText),
			tgbotapi.NewKeyboardButton(denyText),
		},
	)
	kb.OneTimeKeyboard = true
	kb.ResizeKeyboard = true
	msgConfig.ReplyMarkup = kb

	return s.bot.Send(msgConfig)
}

func (s *authSession) sendBotReply(replyTo tgbotapi.Message, err string) error {
	msgConfig := tgbotapi.NewMessage(s.chatID, err)
	msgConfig.ReplyToMessageID = replyTo.MessageID
	msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	_, sendErr := s.bot.Send(msgConfig)

	return sendErr
}

func (s *authSession) Close() {
	s.logger.Debugf("sess cleanup: start")
	defer s.logger.Debugf("sess cleanup: finish")

	close(s.startChan)
	close(s.declineButtonChan)
	close(s.allowButtonChan)
	close(s.timeoutChan)
	close(s.resultChan)
	close(s.shutdownChan)
	close(s.timerCancelChan)
}
