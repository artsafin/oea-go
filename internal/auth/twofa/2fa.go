package twofa

import (
	"oea-go/internal/common"
	"oea-go/internal/common/config"
)

type Flow interface {
	StartAuthFlow(account config.Account, info common.AuthInfo) (isNewSession bool, err error)
	GetSession(account config.Account) (session Session, err error)
}

type Result struct {
	Err         error
	Fingerprint string
}

type Session interface {
	ResultChan() <-chan Result
}
