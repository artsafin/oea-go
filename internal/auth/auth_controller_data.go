package auth

import (
	"net/http"
	"oea-go/internal/auth/authtoken"
)

type Status uint8

const (
	StatusNew                    Status = iota
	StatusError                  Status = iota
	StatusFirstFactorSent        Status = iota
	StatusSecondFactorNewSession Status = iota
	StatusSecondFactorCached     Status = iota
)

type authControllerData struct {
	ReturnUrl string
	Error     string
	PrevEmail string
	Status    Status
	Token     string
}

func (d authControllerData) IsNew() bool {
	return d.Status == StatusNew
}

func (d authControllerData) IsError() bool {
	return d.Status == StatusError
}

func (d authControllerData) IsFirstFactorSent() bool {
	return d.Status == StatusFirstFactorSent
}

func (d authControllerData) IsSecondFactorNewSession() bool {
	return d.Status == StatusSecondFactorNewSession
}

func (d authControllerData) IsSecondFactorCached() bool {
	return d.Status == StatusSecondFactorCached
}

func (d authControllerData) IsSecondFactor() bool {
	return d.IsSecondFactorCached() || d.IsSecondFactorNewSession()
}

func newAuthControllerDataFromRequest(req *http.Request) authControllerData {
	errFromRequest := req.URL.Query().Get("err")

	initStatus := StatusNew
	if errFromRequest != "" {
		initStatus = StatusError
	}
	return authControllerData{
		ReturnUrl: sanitizeReturnUrl(req.URL.Query().Get("return")),
		Error:     errFromRequest,
		Status:    initStatus,
	}
}

func (d authControllerData) WithError(err string) authControllerData {
	return authControllerData{
		ReturnUrl: d.ReturnUrl,
		Error:     err,
		PrevEmail: d.PrevEmail,
		Status:    StatusError,
	}
}

func (d authControllerData) WithFirstFactor() authControllerData {
	return authControllerData{
		ReturnUrl: d.ReturnUrl,
		Error:     "",
		PrevEmail: d.PrevEmail,
		Status:    StatusFirstFactorSent,
	}
}

func (d authControllerData) WithSecondFactor(isNewSession bool, token *authtoken.Token) authControllerData {
	status := StatusSecondFactorNewSession

	if !isNewSession {
		status = StatusSecondFactorCached
	}

	return authControllerData{
		ReturnUrl: d.ReturnUrl,
		Error:     "",
		PrevEmail: d.PrevEmail,
		Status:    status,
		Token:     token.String(),
	}
}
