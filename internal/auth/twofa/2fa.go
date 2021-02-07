package twofa

import "oea-go/internal/common"

type AuthResult struct {
	Err         error
	Fingerprint string
}

type TwoFactorAuthRoutine interface {
	Authenticate(authResult chan AuthResult, account common.Account, info common.AuthInfo) error
}
