package tg2fa

import (
	"fmt"
	"github.com/pkg/errors"
	"oea-go/internal/auth/enc"
	"oea-go/internal/common/config"
	"time"
)

type userReplyMeta struct {
	chatID   int64
	userID   int
	username config.Username
}

func (r *userReplyMeta) validate(expectedChatID int64, expectedAcc config.Account) error {
	if r.username == "" || r.username != expectedAcc.ExternalUsername {
		return errors.Errorf("username mismatch: %v != %v", r.username, expectedAcc.ExternalUsername)
	}
	if expectedChatID != 0 && r.chatID != expectedChatID {
		return errors.New("chat ID mismatch")
	}

	return nil
}

func (r *userReplyMeta) fingerprint(encSecretKey []byte) (string, error) {
	var key [32]byte
	copy(key[:], encSecretKey)

	plain := fmt.Sprintf("%v:%v:%v:%v", r.userID, r.chatID, r.username, time.Now().Unix())

	cipher, encErr := enc.Encrypt([]byte(plain), key)

	if encErr == nil && cipher != nil {
		return fmt.Sprintf("%x", cipher), nil
	}

	return "", encErr
}