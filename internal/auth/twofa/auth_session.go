package twofa

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"oea-go/internal/common"
	"time"
)

const (
	sessionTimeout = time.Second * 30
	allowText      = "🔑 Allow"
	denyText       = "⛔ Deny"
	startText      = "/start"
)

type tgAuthSession struct {
	startChan         chan tgUserReplyMeta
	allowButtonChan   chan tgUserReplyMeta
	declineButtonChan chan tgUserReplyMeta
	timeoutChan       chan struct{}
	bot               *tgbotapi.BotAPI
	account           common.Account
}

func newAuthSession(account common.Account, api *tgbotapi.BotAPI) *tgAuthSession {
	return &tgAuthSession{
		account:           account,
		startChan:         make(chan tgUserReplyMeta),
		allowButtonChan:   make(chan tgUserReplyMeta),
		declineButtonChan: make(chan tgUserReplyMeta),
		timeoutChan:       make(chan struct{}),
		bot:               api,
	}
}

func (s *tgAuthSession) runExpirationTimer(expirationsC expirationsChannel) {
	<-time.After(sessionTimeout)

	s.timeoutChan <- struct{}{}
	expirationsC <- s.account
}

func (s *tgAuthSession) sendPrompt(chatID int64, info common.AuthInfo) (err error) {
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

func (s *tgAuthSession) terminate() {
	close(s.startChan)
	close(s.declineButtonChan)
	close(s.allowButtonChan)
	close(s.timeoutChan)
}
