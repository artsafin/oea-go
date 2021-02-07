package twofa

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"oea-go/internal/common"
	"sync"
)

type expirationsChannel chan common.Account

type tgBot struct {
	token           string
	bot             *tgbotapi.BotAPI
	updChan         tgbotapi.UpdatesChannel
	sessions        map[common.Username]*tgAuthSession
	expirationsChan expirationsChannel
	mut             sync.Mutex
}

func newBotSupervisor(token string) *tgBot {
	return &tgBot{
		bot:      nil,
		updChan:  nil,
		token:    token,
		sessions: make(map[common.Username]*tgAuthSession),
	}
}

func (b *tgBot) resurrect() error {
	if b.updChan != nil {
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

func (b *tgBot) CleanupExpiredSessions() {
	for {
		select {
		case expiredAcc := <-b.expirationsChan:
			b.mut.Lock()

			u := expiredAcc.ExternalUsername
			sess, found := b.sessions[u]
			if !found {
				b.mut.Unlock()
				continue
			}
			sess.terminate()
			delete(b.sessions, u)

			if len(b.sessions) == 0 {
				b.stop()
			}

			b.mut.Unlock()
		}
	}
}

func (b *tgBot) ListenUpdates() {
	for update := range b.updChan {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		userName := common.Username(update.Message.From.UserName)
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

func (b *tgBot) stop() {
	b.bot.StopReceivingUpdates()
	b.bot.Client.CloseIdleConnections()
	b.bot = nil
	b.updChan = nil
	close(b.expirationsChan)

	for k, s := range b.sessions {
		s.terminate()

		delete(b.sessions, k)
	}
}

func (b *tgBot) BeginAuthSession(account common.Account) (*tgAuthSession, error) {
	b.mut.Lock()
	defer b.mut.Unlock()

	resurrectErr := b.resurrect()
	if resurrectErr != nil {
		return nil, resurrectErr
	}

	sess := newAuthSession(account, b.bot)

	go sess.runExpirationTimer(b.expirationsChan)

	b.sessions[account.ExternalUsername] = sess

	return sess, nil
}
