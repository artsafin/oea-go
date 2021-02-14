package tg2fa

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"oea-go/internal/auth/twofa"
	"oea-go/internal/common"
	"oea-go/internal/common/config"
)

var botsrv *botServer

type telegramFlow struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewTelegramTwoFactorAuth(cfg *config.Config, logger *zap.SugaredLogger) twofa.Flow {
	if botsrv == nil {
		botsrv = newBotServer(cfg.BotToken, logger)
	}

	return &telegramFlow{cfg: cfg, logger: logger}
}

func (a *telegramFlow) StartAuthFlow(account config.Account, info common.AuthInfo) (isNewSession bool, err error) {
	botapi, err := botsrv.resurrect()
	if err != nil {
		return true, err
	}
	sess := newAuthSession(account, info, botapi, a.logger, a.cfg.SecretKey)

	err = botsrv.registerSession(sess)
	if err != nil {
		return true, err
	}

	chatID, err := a.obtainChatIdFromCache(account.Email)
	if err != nil {
		a.logger.Warnf("auth flow: obtainChatIdFromCache: %v", err)
	}

	go sess.Flow(chatID)

	return chatID == 0, nil
}

func (a *telegramFlow) GetSession(account config.Account) (session twofa.Session, err error) {
	if sess, found := botsrv.sessions[account.ExternalUsername]; found {
		return sess, nil
	}
	return nil, errors.Errorf("missing session for account %v", account)
}

func (a *telegramFlow) obtainChatIdFromCache(email config.Email) (chatID int64, err error) {
	return 0, errors.New("not implemented")
}
