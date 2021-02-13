package tg2fa

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"oea-go/internal/auth/twofa"
	"oea-go/internal/common"
	"oea-go/internal/common/config"
	"time"
)

const (
	sessionTimeout = time.Second * 30
	allowText      = "🔑 Allow"
	denyText       = "⛔ Deny"
	startText      = "/start"
)

type authSession struct {
	startChan         chan userReplyMeta
	allowButtonChan   chan userReplyMeta
	declineButtonChan chan userReplyMeta
	timeoutChan       chan struct{}
	resultChan        chan twofa.Result
	bot               *tgbotapi.BotAPI
	logger            *zap.SugaredLogger
	account           config.Account
	encKey            []byte
	info              common.AuthInfo
}

func newAuthSession(account config.Account, info common.AuthInfo, api *tgbotapi.BotAPI, logger *zap.SugaredLogger, encKey []byte) *authSession {
	logger.Debugf("session: new session %+v", account)

	return &authSession{
		account:           account,
		info:              info,
		startChan:         make(chan userReplyMeta),
		allowButtonChan:   make(chan userReplyMeta),
		declineButtonChan: make(chan userReplyMeta),
		timeoutChan:       make(chan struct{}),
		resultChan:        make(chan twofa.Result),
		bot:               api,
		logger:            logger,
		encKey:            encKey,
	}
}

func (s *authSession) ResultChan() <-chan twofa.Result {
	return s.resultChan
}

func (s *authSession) flow(chatID int64) {
	s.logger.Debugf("auth flow: start (cached chat id %v)", chatID)
	defer s.logger.Debugf("auth flow: finish")

	defer s.terminate()

	var err error
	if chatID == 0 {
		chatID, err = s.waitForChatIdFromUserStartReply(s.account)
	}
	if err != nil {
		s.resultChan <- twofa.Result{Err: errors.Wrap(err, "cannot obtain chat ID")}
		return
	}

	promptErr := s.sendPrompt(chatID, s.info)
	if promptErr != nil {
		s.resultChan <- twofa.Result{Err: errors.Wrap(promptErr, "cannot send auth prompt")}
		return
	}

	select {
	case denyReply := <-s.declineButtonChan:
		s.resultChan <- twofa.Result{Err: errors.Errorf("user denied: %+v", denyReply)}
	case allowReply := <-s.allowButtonChan:
		if validationErr := allowReply.validate(chatID, s.account); validationErr != nil {
			s.resultChan <- twofa.Result{Err: validationErr}
			break
		}
		fp, fpErr := allowReply.fingerprint(s.encKey)
		if fpErr != nil {
			s.resultChan <- twofa.Result{Err: errors.New("fingerprint encryption failed")}
			break
		}
		s.resultChan <- twofa.Result{Fingerprint: fp}
	case <-s.timeoutChan:
		s.resultChan <- twofa.Result{Err: errors.New("auth session timeout")}
	}
}

func (s *authSession) waitForChatIdFromUserStartReply(account config.Account) (chatID int64, err error) {
	select {
	case startReply := <-s.startChan:
		if err = startReply.validate(0, account); err != nil {
			return 0, err
		}
		return startReply.chatID, nil
	case <-s.timeoutChan:
		return 0, errors.New("timeout waiting for start")
	}
}

func (s *authSession) expireAfterTimeout(expirationsChan accountChannel) {
	<-time.After(sessionTimeout)

	expirationsChan <- s.account
	s.timeoutChan <- struct{}{}
}

func (s *authSession) sendPrompt(chatID int64, info common.AuthInfo) (err error) {
	text := fmt.Sprintf(
		"OEA authentication has been requested on `%v` from IP `%v`.\n\n"+
			"- To allow access please press 🔑 `Allow`\n"+
			"- To deny authentication please press ⛔ `Deny` or ignore this message",
		info.TS.Format(time.UnixDate),
		info.IP,
	)
	msgConfig := tgbotapi.NewMessage(chatID, text)
	kb := tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(allowText),
			tgbotapi.NewKeyboardButton(denyText),
		},
	)
	kb.OneTimeKeyboard = true
	kb.ResizeKeyboard = true
	msgConfig.ReplyMarkup = kb

	_, err = s.bot.Send(msgConfig)

	return err
}

func (s *authSession) terminate() {
	close(s.startChan)
	close(s.declineButtonChan)
	close(s.allowButtonChan)
	close(s.timeoutChan)
	close(s.resultChan)
}
