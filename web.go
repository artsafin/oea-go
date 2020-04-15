package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"oea-go/common"
	"os"
	"time"
)

const (
	RequestIdHeader     = "X-Request-Id"
	RequestIdContextKey = "requestId"
)

type webRouter struct {
	*mux.Router
	AuthToken *Token
}

// Writes Location header to response writer and sets specified status
func httpRedirect(resp http.ResponseWriter, url string, status int) {
	resp.Header().Add("Location", url)
	resp.WriteHeader(status)
}

func (router *webRouter) createFuncMap() template.FuncMap {
	return template.FuncMap{
		"link": func(routeName string, args ...string) string {
			var route *mux.Route
			var u *url.URL
			var err error
			if route = router.Get(routeName); route == nil {
				return ""
			}
			if u, err = route.URL(args...); err != nil {
				return ""
			}
			return u.String()
		},
		"styles": func() template.CSS {
			return template.CSS(common.MustAsset("resources/bootstrap.min.css"))
		},
		"authToken": func() *Token {
			return router.AuthToken
		},
	}
}

func (router *webRouter) parseLayout() *template.Template {
	return template.Must(
		template.
			New("layout").
			Funcs(router.createFuncMap()).
			Parse(string(common.MustAsset("resources/layout.go.html"))))
}

type Partial struct {
	*template.Template
}

type PartialData struct {
	Page       interface{}
	RequestURI string
}

func NewPartialData(data interface{}, r *http.Request) *PartialData {
	return &PartialData{
		Page:       data,
		RequestURI: r.RequestURI,
	}
}

func (router *webRouter) createPartial(names ...string) Partial {
	tpl := router.parseLayout()
	for _, name := range names {
		tpl = template.Must(tpl.Parse(string(common.MustAsset(fmt.Sprintf("resources/partials/%s.go.html", name)))))
	}

	return Partial{tpl}
}

func (partial Partial) MustRenderWithData(wr io.Writer, data *PartialData) {
	if err := partial.Execute(wr, data); err != nil {
		panic(err)
	}
}

func (router *webRouter) page(templateDataFn func(map[string]string, *http.Request) interface{}, names ...string) func(http.ResponseWriter, *http.Request) {
	tpl := router.createPartial(names...)

	return func(resp http.ResponseWriter, req *http.Request) {
		data := NewPartialData(templateDataFn(mux.Vars(req), req), req)

		tpl.MustRenderWithData(resp, data)
	}
}

func listenAndServe(cfg common.Config, routerConfigurer func(*webRouter)) {
	router := &webRouter{Router: mux.NewRouter()}
	auth := &authMiddleware{
		router: router,
		config: cfg,
	}
	router.Use(faviconMiddleware)
	router.Use(requestIdMiddleware)
	router.Use(loggerMiddleware)
	router.Use(auth.Middleware)
	router.NotFoundHandler = http.HandlerFunc(router.page(nilTemplateData, "404"))

	authHandler := authController{cfg, router.createPartial("auth")}
	router.HandleFunc("/auth/success", authHandler.HandleSendSuccess)
	router.HandleFunc("/auth/set", authHandler.HandleTokenSet)
	router.HandleFunc("/auth/logout", authHandler.HandleLogout).Methods(http.MethodPost).Name("Logout")
	router.HandleFunc("/auth", authHandler.HandleAuthStart)

	routerConfigurer(router)

	var certManager autocert.Manager
	if cfg.IsTLS() {
		certsDirErr := os.MkdirAll(cfg.CertsDir, os.ModePerm)
		if certsDirErr != nil {
			panic(certsDirErr)
		}
		certManager = autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(cfg.LetsEncryptDomain),
			Cache:      autocert.DirCache(cfg.CertsDir),
		}
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.InsecurePort),
		Handler:      router,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	var listenErr error
	if cfg.IsTLS() {
		server.TLSConfig = &tls.Config{
			GetCertificate: certManager.GetCertificate,
		}
		server.Addr = fmt.Sprintf(":%d", cfg.SecurePort)

		go http.ListenAndServe(fmt.Sprintf(":%d", cfg.InsecurePort), certManager.HTTPHandler(nil))
		log.Printf("Server started at %s (secure)\n", server.Addr)
		listenErr = server.ListenAndServeTLS("", "")
	} else {
		log.Printf("Server started at %s (insecure)\n", server.Addr)
		listenErr = server.ListenAndServe()
	}
	log.Fatal("Server shutdown", listenErr)
}

func nilTemplateData(vars map[string]string, req *http.Request) interface{} {
	return nil
}

func faviconMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.RequestURI == "/favicon.ico" {
			writer.Header().Add("Content-Type", "image/png")
			writer.WriteHeader(http.StatusOK)
			writer.Write(common.MustAsset("resources/icon.png"))
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		logger := NewRequestLogger(r)
		logger.Println("Request:", r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())

		next.ServeHTTP(writer, r)
	})
}

func requestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get(RequestIdHeader)
		if requestId == "" {
			requestId = uuid.Must(uuid.NewRandom()).String()
		}

		ctx := context.WithValue(r.Context(), RequestIdContextKey, requestId)

		writer.Header().Set(RequestIdHeader, requestId)

		next.ServeHTTP(writer, r.WithContext(ctx))
	})
}

func getRequestId(r *http.Request) string {
	id, ok := r.Context().Value(RequestIdContextKey).(string)
	if !ok {
		return "no-request-id"
	}
	return id
}
