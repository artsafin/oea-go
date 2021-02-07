package twofa

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"oea-go/internal/common"
	"time"
)

const (
	sessionTimeout = time.Second * 30
)

type tgAuthSession struct {
	startChan         chan tgUserReplyMeta
	allowButtonChan   chan tgUserReplyMeta
	declineButtonChan chan tgUserReplyMeta
	timeoutChan       chan int64
	bot               *tgbotapi.BotAPI
	ts                time.Time
}

func newAuthSession(api *tgbotapi.BotAPI) *tgAuthSession {
	return &tgAuthSession{
		startChan:         make(chan tgUserReplyMeta),
		allowButtonChan:   make(chan tgUserReplyMeta),
		declineButtonChan: make(chan tgUserReplyMeta),
		timeoutChan:       make(chan int64),
		bot:               api,
		ts:                time.Now(),
	}
}

func (s tgAuthSession) sendPrompt(chatID int64, info common.AuthInfo) (err error) {
	text := fmt.Sprintf(
		"OEA authentication has been requested on %v from IP %v.\n"+
			"- To allow access please press 🔑 `Allow`\n"+
			"- To deny authentication please press ⛔ `Deny` or ignore this message",
		info.TS.Format(time.UnixDate),
		info.IP,
	)
	msgConfig := tgbotapi.NewMessage(chatID, text)
	kb := tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButtonContact("🔑 Allow"),
			tgbotapi.NewKeyboardButtonContact("⛔ Deny"),
		},
	)
	kb.OneTimeKeyboard = true
	kb.ResizeKeyboard = true
	msgConfig.ReplyMarkup = kb

	_, err = s.bot.Send(msgConfig)

	return
}
