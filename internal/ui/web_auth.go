package ui

import (
	"fmt"
	"github.com/badoux/checkmail"
	"net/http"
	"net/url"
	"oea-go/internal/common"
	"strings"
	"time"
)

const (
	AuthCookieName = "a"
)

var anonUrls = []string{
	"/auth",
	"/auth/success",
	"/auth/set",
}

type authController struct {
	config  common.Config
	partial Partial
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
		Name:     AuthCookieName,
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

func (ctl authController) HandleSendSuccess(resp http.ResponseWriter, req *http.Request) {
	ctl.partial.MustRenderWithData(resp, NewPartialData(authControllerData{IsSent: true}, req))
}

func (ctl authController) HandleTokenSet(resp http.ResponseWriter, req *http.Request) {
	returnUrl := sanitizeReturnUrl(req.URL.Query().Get("return"))
	token := req.URL.Query().Get("t")

	logger := common.NewRequestLogger(req)

	if token == "" {
		logger.Println("Token set: token is empty")
		httpRedirect(resp, "/auth", http.StatusFound)
		return
	}

	tok, tokErr := VerifiedTokenFromSource(ctl.config, token)
	if tokErr != nil {
		logger.Println("Token set: token is invalid:", tokErr)
		httpRedirect(resp, "/auth?err="+url.QueryEscape("your login link has expired"), http.StatusFound)
		return
	}

	validationErr := tok.ValidateClaims()
	if validationErr != nil {
		logger.Printf("Token set: token is invalid: %s, token %v\n", validationErr, tok)
		httpRedirect(resp, "/auth?err="+url.QueryEscape("your login link has expired"), http.StatusFound)
		return
	}

	authCookie := newAuthCookie(token, req.Host, tok.Claims.ExpiresAt.Time())
	logger.Println("Token set: auth cookie set:", authCookie)
	http.SetCookie(resp, authCookie)
	httpRedirect(resp, returnUrl, http.StatusFound)
}

func (ctl authController) HandleLogout(resp http.ResponseWriter, req *http.Request) {
	logger := common.NewRequestLogger(req)
	returnUrl := sanitizeReturnUrl(req.URL.Query().Get("return"))

	authCookie := newAuthCookie("0", req.Host, time.Unix(0, 0))

	logger.Printf("Logout authCookie: %+v", authCookie)

	http.SetCookie(resp, authCookie)
	httpRedirect(resp, returnUrl, http.StatusFound)
}

func (ctl authController) HandleAuthStart(resp http.ResponseWriter, req *http.Request) {
	data := authControllerData{
		ReturnUrl: sanitizeReturnUrl(req.URL.Query().Get("return")),
		Error:     req.URL.Query().Get("err"),
	}

	if req.Method == http.MethodGet {
		ctl.partial.MustRenderWithData(resp, NewPartialData(data, req))
		return
	}

	recipient := req.PostFormValue("email")
	data.PrevEmail = recipient

	if !ctl.config.IsAuthAllowed(recipient) {
		data.Error = "specified email is not permitted"
		ctl.partial.MustRenderWithData(resp, NewPartialData(data, req))
		return
	}

	if emailValidErr := checkmail.ValidateFormat(recipient); emailValidErr != nil {
		data.Error = fmt.Sprintf("submitted email is incorrect: %s", emailValidErr)
		ctl.partial.MustRenderWithData(resp, NewPartialData(data, req))
		return
	}
	if emailValidErr := checkmail.ValidateHost(recipient); emailValidErr != nil {
		data.Error = fmt.Sprintf("submitted email is incorrect: %s", emailValidErr)
		ctl.partial.MustRenderWithData(resp, NewPartialData(data, req))
		return
	}

	newToken, tokErr := GenerateToken(recipient, ctl.config)

	if tokErr != nil {
		data.Error = fmt.Sprintf("error generating token: %v", tokErr)
		ctl.partial.MustRenderWithData(resp, NewPartialData(data, req))
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
	queryString.Set("t", newToken)
	link.RawQuery = queryString.Encode()

	if err := sendMail(req.RemoteAddr, link, recipient, ctl.config); err != nil {
		data.Error = err.Error()
		ctl.partial.MustRenderWithData(resp, NewPartialData(data, req))
		return
	}

	httpRedirect(resp, "/auth/success", http.StatusSeeOther)
}

type authMiddleware struct {
	router *WebRouter
	config common.Config
}

func (auth *authMiddleware) doAuth(r *http.Request) error {
	auth.router.AuthToken = nil

	authCookie, cookieErr := r.Cookie(AuthCookieName)
	if cookieErr != nil {
		return fmt.Errorf("missing cookie: %s", cookieErr)
	}

	token, tokenErr := VerifiedTokenFromSource(auth.config, authCookie.Value)

	if tokenErr != nil {
		return tokenErr
	}

	validationErr := token.ValidateClaims()
	if validationErr != nil {
		return fmt.Errorf("invalid claims: %s, token: %v", validationErr, token)
	}
	auth.router.AuthToken = token

	return nil
}

func isAnonymousAccessAllowed(path string) bool {
	for _, anonUrl := range anonUrls {
		if anonUrl == path {
			return true
		}
	}

	return false
}

func (auth *authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := common.NewRequestLogger(r)
		// Do actual authentication of request
		authErr := auth.doAuth(r)

		// Allow /auth anonymous access
		if isAnonymousAccessAllowed(r.URL.Path) {
			// User is actually authenticated but he is on /auth page => redirect to /
			if authErr == nil {
				logger.Println("Redirecting already authenticated user from", r.RequestURI)
				returnUrl := sanitizeReturnUrl(r.URL.Query().Get("return"))
				httpRedirect(w, returnUrl, http.StatusFound)
				return
			}
			next.ServeHTTP(w, r)
			return
		}

		// Authentication successful, pass to next handler
		if authErr == nil {
			next.ServeHTTP(w, r)
			return
		}

		logger.Println("Forbidden:", authErr)

		// If page was requested with GET, redirect to /auth page with ability to return back after successful auth
		if r.Method == http.MethodGet {
			httpRedirect(w, fmt.Sprintf("/auth?return=%s", url.QueryEscape(r.RequestURI)), http.StatusFound)
			return
		}

		// If nothing can be done, just throw 403 page
		auth.router.Page(NilTemplateData, "403")(w, r)
	})
}
