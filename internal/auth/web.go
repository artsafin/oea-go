package auth

import (
	"encoding/json"
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"oea-go/internal/auth/authtoken"
	"oea-go/internal/auth/twofa/tg2fa"
	"oea-go/internal/common"
	"oea-go/internal/common/config"
	"oea-go/internal/web"
	"strings"
	"time"
)

const (
	CookieName = "a"
)

var anonUrls = []string{
	"/auth",
	"/auth/sent",
	"/auth/set",
	"/auth/twofa",
}

func newAuthInfo(req *http.Request) common.AuthInfo {
	return common.AuthInfo{IP: req.RemoteAddr, TS: time.Now()}
}

func authErrorRedirect(resp http.ResponseWriter, err string) {
	web.HttpRedirect(resp, "/auth?err="+url.QueryEscape(err), http.StatusFound)
}

func NewHandler(cfg *config.Config, partial web.Partial, logger *zap.SugaredLogger, router *mux.Router) *Controller {
	return &Controller{
		config:  cfg,
		partial: partial,
		logger:  logger,
		router:  router,
	}
}

type Controller struct {
	config  *config.Config
	partial web.Partial
	logger  *zap.SugaredLogger
	router  *mux.Router
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
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	}
}

func sanitizeReturnUrl(returnUrl string) string {
	if returnUrl == "" || strings.HasPrefix(returnUrl, "/auth") || !strings.HasPrefix(returnUrl, "/") {
		return "/"
	}

	return returnUrl
}

func (ctl Controller) HandleFirstFactorSendSuccess(resp http.ResponseWriter, req *http.Request) {
	ctl.partial.MustRenderWithData(resp, web.NewPartialData(authControllerData{}.WithFirstFactor(), req))
}

func extractAccountFromJWT(req *http.Request, config *config.Config) (account *config.Account, token *authtoken.Token, err error) {
	tokenSource := req.URL.Query().Get("t")

	if tokenSource == "" {
		tokenSource = strings.Replace(req.Header.Get("Authorization"), "Bearer ", "", 1)
	}

	token, err = authtoken.CreateFromSourceAndValidate(config.AppVersion, config.SecretKey, tokenSource)
	if err != nil {
		return nil, nil, errors.Wrap(err, "HandleBeginSecondFactor: token parse error")
	}

	email, err := token.Email()
	if err != nil {
		return nil, nil, errors.Wrap(err, "HandleBeginSecondFactor: email retrieve")
	}
	account = config.Accounts.Get(email)
	if account == nil {
		return nil, nil, errors.New("HandleBeginSecondFactor: account not found")
	}

	return account, token, nil
}

func (ctl Controller) HandleBeginSecondFactor(resp http.ResponseWriter, req *http.Request) {
	logger := common.NewRequestLogger(req, ctl.logger)

	account, token, accountFetchErr := extractAccountFromJWT(req, ctl.config)
	if accountFetchErr != nil {
		logger.Errorf("HandleBeginSecondFactor: %v", accountFetchErr)
		authErrorRedirect(resp, "incorrect login link")
		return
	}

	twoFa := tg2fa.NewTelegramTwoFactorAuth(ctl.config, ctl.logger)
	isNewSession, sessExpTs, twoFAErr := twoFa.StartAuthFlow(*account, newAuthInfo(req))
	if twoFAErr != nil {
		logger.Errorf("HandleBeginSecondFactor: StartAuthFlow: %v", twoFAErr)
		authErrorRedirect(resp, "cannot initialize 2fa")
		return
	}

	data := newAuthControllerDataFromRequest(req)
	partialData := data.WithSecondFactor(isNewSession, token.Source, sessExpTs)
	if route := ctl.router.Get("AuthCheck2FA"); route != nil {
		routeUrl, err := route.URL()
		if err != nil {
			partialData.Error = err.Error()
		} else {
			partialData.CheckUrl = routeUrl.String()
		}
	}

	ctl.partial.MustRenderWithData(resp, web.NewPartialData(partialData, req))
}

func (ctl Controller) HandleCheckSecondFactor(resp http.ResponseWriter, req *http.Request) {
	logger := common.NewRequestLogger(req, ctl.logger)

	respondWithError := func(err error, repeatRequest bool) {
		data := new2FaError(err, repeatRequest)

		resp.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(resp).Encode(data); err != nil {
			logger.Errorf("response error: %v", err)
		}
	}

	account, token, accountFetchErr := extractAccountFromJWT(req, ctl.config)
	if accountFetchErr != nil {
		logger.Errorf("HandleCheckSecondFactor: %v", accountFetchErr)
		respondWithError(accountFetchErr, false)
		return
	}

	twoFa := tg2fa.NewTelegramTwoFactorAuth(ctl.config, ctl.logger)
	sess, sessErr := twoFa.GetSession(*account)
	if sessErr != nil {
		logger.Errorf("HandleCheckSecondFactor: GetSession: %v", sessErr)
		respondWithError(sessErr, false)
		return
	}

	timeout := time.After(time.Millisecond * 500)
	select {
	case <-timeout:
		respondWithError(errors.New("timeout"), true)
		return
	case authRes := <-sess.ResultChan():
		if authRes.Err != nil {
			respondWithError(authRes.Err, false)
			return
		}
		newToken, newTokenErr := authtoken.GenerateTokenSecondFactor(token, authRes.Fingerprint, ctl.config.SecretKey)

		if newTokenErr != nil {
			logger.Errorf("HandleCheckSecondFactor: GenerateTokenSecondFactor: %v", newTokenErr)
			respondWithError(newTokenErr, false)
			return
		}

		resp.WriteHeader(http.StatusOK)
		data := new2FaSuccess(newToken.InsecureString(), token.Claims.ReturnUrl)
		if err := json.NewEncoder(resp).Encode(data); err != nil {
			logger.Errorf("response error: %v", err)
		}
	}
}

func (ctl Controller) HandleLogout(resp http.ResponseWriter, req *http.Request) {
	logger := common.NewRequestLogger(req, ctl.logger)
	returnUrl := sanitizeReturnUrl(req.URL.Query().Get("return"))

	authCookie := newAuthCookie("0", req.Host, time.Unix(0, 0))

	logger.Infof("Logout authCookie: %+v", authCookie)

	http.SetCookie(resp, authCookie)
	web.HttpRedirect(resp, returnUrl, http.StatusFound)
}

func (ctl Controller) HandleAuthStart(resp http.ResponseWriter, req *http.Request) {
	data := newAuthControllerDataFromRequest(req)

	if req.Method == http.MethodGet {
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data, req))
		return
	}

	recipient := req.PostFormValue("email")
	data.PrevEmail = recipient

	if !ctl.config.IsAuthAllowed(recipient) {
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data.WithError("specified email is not permitted"), req))
		return
	}

	if emailValidErr := checkmail.ValidateFormat(recipient); emailValidErr != nil {
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data.WithError(fmt.Sprintf("submitted email is incorrect: %s", emailValidErr)), req))
		return
	}
	if emailValidErr := checkmail.ValidateHost(recipient); emailValidErr != nil {
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data.WithError(fmt.Sprintf("submitted email is incorrect: %s", emailValidErr)), req))
		return
	}

	newToken, tokErr := authtoken.GenerateTokenFirstFactor(
		ctl.config.AppVersion,
		recipient,
		ctl.config.SecretKey,
		sanitizeReturnUrl(req.URL.Query().Get("return")),
	)

	if tokErr != nil {
		ctl.partial.MustRenderWithData(resp, web.NewPartialData(data.WithError(fmt.Sprintf("error generating token: %v", tokErr)), req))
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
	queryString.Del("return")
	link.RawQuery = queryString.Encode()

	authInfo := newAuthInfo(req)

	if ctl.config.IsEmailsEnabled() {
		if err := sendMail(authInfo, link, recipient, ctl.config); err != nil {
			ctl.partial.MustRenderWithData(resp, web.NewPartialData(data.WithError(err.Error()), req))
			return
		}
	} else {
		// For debug
		ctl.logger.Infow("Login", "link", link.String(), "authInfo", authInfo)
	}

	web.HttpRedirect(resp, "/auth/sent", http.StatusSeeOther)
}
