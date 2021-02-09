package auth

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"oea-go/internal/common"
	"oea-go/internal/web"
)

type Middleware struct {
	Router *web.Engine
	Config common.Config
	Logger *zap.SugaredLogger
}

func (auth *Middleware) doAuth(r *http.Request) error {
	auth.Router.AuthToken = nil

	authCookie, cookieErr := r.Cookie(CookieName)
	if cookieErr != nil {
		return fmt.Errorf("missing cookie: %s", cookieErr)
	}

	token, tokenErr := validatedTokenFromSource(auth.Config.AppVersion, auth.Config.SecretKey, authCookie.Value)

	if tokenErr != nil {
		return tokenErr
	}

	auth.Router.AuthToken = token

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

func (auth *Middleware) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := common.NewRequestLogger(r, auth.Logger)
		// Do actual authentication of request
		authErr := auth.doAuth(r)

		// Allow /auth anonymous access
		if isAnonymousAccessAllowed(r.URL.Path) {
			// User is actually authenticated but he is on /auth page => redirect to /
			if authErr == nil {
				logger.Infof("Redirecting already authenticated user from %v", r.RequestURI)
				returnUrl := sanitizeReturnUrl(r.URL.Query().Get("return"))
				web.HttpRedirect(w, returnUrl, http.StatusFound)
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

		logger.Errorf("Forbidden: %v", authErr)

		// If page was requested with GET, redirect to /auth page with ability to return back after successful auth
		if r.Method == http.MethodGet {
			web.HttpRedirect(w, fmt.Sprintf("/auth?return=%s", url.QueryEscape(r.RequestURI)), http.StatusFound)
			return
		}

		// If nothing can be done, just throw 403 page
		auth.Router.Page(web.NilTemplateData, "403")(w, r)
	})
}
