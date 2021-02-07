package twofa

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"oea-go/internal/common"
	"sync"
)

type tgBotSupervisor struct {
	token        string
	bot          *tgbotapi.BotAPI
	updChan      tgbotapi.UpdatesChannel
	sessions     map[common.Account]*tgAuthSession
	timeoutsChan chan common.Account
	mut          sync.Mutex
}

func newBotSupervisor(token string) *tgBotSupervisor {
	return &tgBotSupervisor{
		bot:          nil,
		updChan:      nil,
		token:        token,
		sessions:     make(map[common.Account]*tgAuthSession),
		timeoutsChan: make(chan common.Account),
	}
}

func (s *tgBotSupervisor) resurrect() error {
	if s.updChan != nil {
		return nil
	}

	var err error
	s.bot, err = tgbotapi.NewBotAPI(s.token)
	if err != nil {
		return err
	}

	s.bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	s.updChan, err = s.bot.GetUpdatesChan(updateConfig)

	if err != nil {
		s.bot.Client.CloseIdleConnections()
		s.bot = nil
		return err
	}

	return nil
}

func (s *tgBotSupervisor) timeoutKiller() {

}

func (s *tgBotSupervisor) stop() {
	s.bot.StopReceivingUpdates()
	s.bot.Client.CloseIdleConnections()
	s.bot = nil
	s.updChan = nil
}

func (s *tgBotSupervisor) beginAuthSession(account common.Account) (*tgAuthSession, error) {
	s.mut.Lock()
	defer s.mut.Unlock()

	resurrectErr := s.resurrect()
	if resurrectErr != nil {
		return nil, resurrectErr
	}

	sess := newAuthSession(s.bot)
	s.sessions[account] = sess

	return sess, nil
}
