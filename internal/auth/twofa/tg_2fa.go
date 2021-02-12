package twofa

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"oea-go/internal/auth/enc"
	"oea-go/internal/common"
	"oea-go/internal/common/config"
	"time"
)

var bots *tgBot

type tgUserReplyMeta struct {
	chatID   int64
	userID   int
	username config.Username
}

func (r *tgUserReplyMeta) validate(expectedChatID int64, expectedAcc config.Account) error {
	if r.username == "" || r.username != expectedAcc.ExternalUsername {
		return errors.New("username mismatch")
	}
	if expectedChatID != 0 && r.chatID != expectedChatID {
		return errors.New("chat ID mismatch")
	}

	return nil
}

func (r *tgUserReplyMeta) fingerprint(encSecretKey []byte) (string, error) {
	var key [32]byte
	copy(key[:], encSecretKey)

	plain := fmt.Sprintf("%v:%v:%v:%v", r.userID, r.chatID, r.username, time.Now().Unix())

	cipher, encErr := enc.Encrypt([]byte(plain), key)

	if encErr == nil && cipher != nil {
		return fmt.Sprintf("%x", cipher), nil
	}

	return "", encErr
}

type telegram2FA struct {
	cfg            *config.Config
	chatIdWaitChan chan int64
}

func NewTelegramTwoFactorAuth(cfg *config.Config, logger *zap.SugaredLogger) TwoFactorAuthRoutine {
	if bots == nil {
		bots = newBotSupervisor(cfg.BotToken, logger)
	}

	return &telegram2FA{cfg: cfg}
}

func (a *telegram2FA) Authenticate(authResult chan AuthResult, account config.Account, info common.AuthInfo) error {
	sess, beginErr := bots.BeginAuthSession(account)
	if beginErr != nil {
		return beginErr
	}

	go func() {
		defer close(authResult)

		chatID, chatIDErr := a.obtainChatId(account, sess)
		if chatIDErr != nil {
			authResult <- AuthResult{Err: errors.Wrap(chatIDErr, "cannot obtain chat ID")}
			return
		}

		promptErr := sess.sendPrompt(chatID, info)
		if promptErr != nil {
			authResult <- AuthResult{Err: errors.Wrap(promptErr, "cannot send auth prompt")}
			return
		}

		select {
		case <-sess.declineButtonChan:
			authResult <- AuthResult{Err: errors.New("user denied")}
		case allowReply := <-sess.allowButtonChan:
			if validationErr := allowReply.validate(chatID, account); validationErr != nil {
				authResult <- AuthResult{Err: validationErr}
				break
			}
			fp, fpErr := allowReply.fingerprint(a.cfg.SecretKey)
			if fpErr != nil {
				authResult <- AuthResult{Err: errors.New("fingerprint encryption failed")}
				break
			}
			authResult <- AuthResult{Fingerprint: fp}
		case <-sess.timeoutChan:
			authResult <- AuthResult{Err: errors.New("auth session timeout")}
		}
	}()

	return nil
}

func (a *telegram2FA) obtainChatId(account config.Account, sess *tgAuthSession) (chatID int64, err error) {
	chatID, err = a.obtainChatIdFromCache(account.Email)
	if err == nil {
		return chatID, nil
	}

	return a.obtainChatIdFromUserStartReply(account, sess)
}

func (a *telegram2FA) obtainChatIdFromUserStartReply(account config.Account, sess *tgAuthSession) (chatID int64, err error) {
	select {
	case startReply := <-sess.startChan:
		if err = startReply.validate(0, account); err != nil {
			return 0, err
		}
		return startReply.chatID, nil
	case <-sess.timeoutChan:
		return 0, errors.New("timeout waiting for start")
	}
}

func (a *telegram2FA) obtainChatIdFromCache(email config.Email) (chatID int64, err error) {
	return 0, errors.New("not implemented")
}
