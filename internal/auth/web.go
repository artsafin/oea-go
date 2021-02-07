package auth

import (
	"fmt"
	"github.com/badoux/checkmail"
	"net/http"
	"oea-go/internal/auth/twofa"
	"oea-go/internal/common"
	"oea-go/internal/web"
	"strings"
	"time"
)

const (
	CookieName = "a"
)

var anonUrls = []string{
	"/auth",
	"/auth/success",
	"/auth/set",
}

func newAuthInfo(req *http.Request) common.AuthInfo {
	return common.AuthInfo{IP: req.RemoteAddr, TS: time.Now()}
}

func authErrorRedirect(resp http.ResponseWriter, err string) {
	authErrorRedirect(resp, err)
}

func NewHandler(cfg *common.Config, partial web.Partial, etcd *common.EtcdService) *Controller {
	return &Controller{config: cfg, etcd: etcd, partial: partial}
}

type Controller struct {
	config  *common.Config
	partial web.Partial
	etcd    *common.EtcdService
}

type authControllerData struct {
	ReturnUrl string
	Error     string
	PrevEmail string
	IsSent    bool
}

func newAuthCookie(value string, cookieDomain string, expiration time.Time) *http.Cookie {
	// Strip port number from hostname (https://stackoverflow.com/questions/1612177/are-http-cookies-port-specific)
	if portColon := strings.Index(cookieDomain, ":"); portColon > 0 {
		cookieDomain = cookieDomain[0:portColon]
	}

	return &http.Cookie{
		Name:     CookieName,
		Value:    value,
		Path:     "/",
		Domain:   cookieDomain,
		Expires:  expiration,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func sanitizeReturnUrl(returnUrl string) string {
	if returnUrl == "" || strings.HasPrefix(returnUrl, "/auth") || !strings.HasPrefix(returnUrl, "/") {
		return "/"
	}

	return returnUrl
}

func (ctl Controller) HandleSendSuccess(resp http.ResponseWriter, req *http.Request) {
	ctl.partial.MustRenderWithData(resp, web.NewPartialData(authControllerData{IsSent: true}, req))
}

func (ctl Controller) HandleBegin2FA(resp http.ResponseWriter, req *http.Request) {
	logger := common.NewRequestLogger(req)

	returnUrl := sanitizeReturnUrl(req.URL.Query().Get("return"))
	tokenSource := req.URL.Query().Get("t")

	token, logErr := validatedTokenFromSource(ctl.config.AppVersion, ctl.config.SecretKey, tokenSource)
	if logErr != nil {
		logger.Printf("HandleBegin2FA: token parse error: %v\n", logErr)
		authErrorRedirect(resp, "incorrect login link")

		return
	}

	email, tokenEmailErr := token.Email()
	if tokenEmailErr != nil {
		logger.Printf("HandleBegin2FA: email retrieve: %v\n", tokenEmailErr)
		authErrorRedirect(resp, "incorrect login link")

		return
	}
	account := ctl.config.Accounts.Get(email)
	if account == nil {
		logger.Printf("HandleBegin2FA: account not found\n")
		authErrorRedirect(resp, "incorrect login link")

		return
	}

	twoFa := twofa.NewTelegramTwoFactorAuth(ctl.etcd, ctl.config)

	authResultChan := make(chan twofa.AuthResult)
	twoFAErr := twoFa.Authenticate(authResultChan, *account, newAuthInfo(req))
	if twoFAErr != nil {
		logger.Println("HandleBegin2FA: Authenticate:", twoFAErr)
		authErrorRedirect(resp, "cannot initialize 2fa")
		return
	}

	timeoutChan := time.After(time.Minute * 10)

	select {
	case authRes, authChanOpen := <-authResultChan:
		if !authChanOpen {
			logger.Println("HandleBegin2FA: 2fa channel closed")
			authErrorRedirect(resp, "2fa error")
			return
		}
		if authRes.Err != nil {
			logger.Println("HandleBegin2FA: 2fa error %v", authRes.Err)
			authErrorRedirect(resp, "2fa error")
			return
		}

		newToken, newTokenErr := GenerateTokenSecondFactor(token, authRes.Fingerprint, ctl.config.SecretKey)

		if newTokenErr != nil {
			logger.Println("HandleBegin2FA: GenerateTokenSecondFactor: %v", newTokenErr)
			authErrorRedirect(resp, "2fa error")
			return
		}

		authCookie := newAuthCookie(newToken.InsecureString(), req.Host, token.ExpiresAt())
		logger.Println("Token set: auth cookie set:", authCookie)
		http.SetCookie(resp, authCookie)
		web.HttpRedirect(resp, returnUrl, http.StatusFound)
	case <-timeoutChan:
		logger.Println("HandleBegin2FA: timeout waiting for auth result")
		authErrorRedirect(resp, "2fa timeout")
	}
}

func (ctl Controller) HandleTokenSet(resp http.ResponseWriter, req *http.Request) {
	returnUrl := sanitizeReturnUrl(req.URL.Query().Get("return"))
	token := req.URL.Query().Get("t")

	logger := common.NewRequestLogger(req)

	if token == "" {
		logger.Println("Token set: token is empty")
		web.HttpRedirect(resp, "/auth", http.StatusFound)
		return
	}

	tok, tokErr := tokenFromSource(ctl.config.AppVersion, ctl.config.SecretKey, token)
	if tokErr != nil {
		logger.Println("Token set: token is invalid:", tokErr)
		authErrorRedirect(resp, "your login link has expired")
		return
	}

	validationErr := tok.ValidateClaims()
	if validationErr != nil {
		logger.Printf("Token set: token is invalid: %s, token %v\n", validationErr, tok)
		authErrorRedirect(resp, "your login link has expired")
		return
	}

	authCookie := newAuthCookie(token, req.Host, tok.Claims.ExpiresAt.Time())
	logger.Println("Token set: auth cookie set:", authCookie)
	http.SetCookie(resp, authCookie)
	web.HttpRedirect(resp, returnUrl, http.StatusFound)
}

func (ctl Controller) HandleLogout(resp http.ResponseWriter, req *http.Request) {
	logger := common.NewRequestLogger(req)
	returnUrl := sanitizeReturnUrl(req.URL.Query().Get("return"))

	authCookie := newAuthCookie("0", req.Host, time.Unix(0, 0))

	logger.Printf("Logout authCookie: %+v", authCookie)

	http.SetCookie(resp, authCookie)
	web.HttpRedirect(resp, returnUrl, http.StatusFound)
}

func (ctl Controller) HandleAuthStart(resp http.ResponseWriter, req *http.Request) {
	data := authControllerData{
		ReturnUrl: sanitizeReturnUrl(req.URL.Query().Get("return")),
		Error:     req.URL.Query().Get("err"),
	}

	if req.Method == http.MethodGet {
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data, req))
		return
	}

	recipient := req.PostFormValue("email")
	data.PrevEmail = recipient

	if !ctl.config.IsAuthAllowed(recipient) {
		data.Error = "specified email is not permitted"
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data, req))
		return
	}

	if emailValidErr := checkmail.ValidateFormat(recipient); emailValidErr != nil {
		data.Error = fmt.Sprintf("submitted email is incorrect: %s", emailValidErr)
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data, req))
		return
	}
	if emailValidErr := checkmail.ValidateHost(recipient); emailValidErr != nil {
		data.Error = fmt.Sprintf("submitted email is incorrect: %s", emailValidErr)
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data, req))
		return
	}

	newToken, tokErr := GenerateTokenFirstFactor(ctl.config.AppVersion, recipient, ctl.config.SecretKey)

	if tokErr != nil {
		data.Error = fmt.Sprintf("error generating token: %v", tokErr)
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data, req))
		return
	}

	link := *req.URL
	link.Host = req.Host
	link.Scheme = "https"
	if req.TLS == nil {
		link.Scheme = "http"
	}
	link.Path = "/auth/set"
	queryString := link.Query()
	queryString.Set("t", newToken.InsecureString())
	link.RawQuery = queryString.Encode()

	if err := sendMail(newAuthInfo(req), link, recipient, ctl.config); err != nil {
		data.Error = err.Error()
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data, req))
		return
	}

	web.HttpRedirect(resp, "/auth/success", http.StatusSeeOther)
}
