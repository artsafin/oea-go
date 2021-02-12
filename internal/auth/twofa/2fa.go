package twofa

import (
	"oea-go/internal/common"
	"oea-go/internal/common/config"
)

type AuthResult struct {
	Err         error
	Fingerprint string
}

type TwoFactorAuthRoutine interface {
	Authenticate(authResult chan AuthResult, account config.Account, info common.AuthInfo) error
}
