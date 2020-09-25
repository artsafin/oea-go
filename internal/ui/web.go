package ui

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"oea-go/internal/common"
	"strings"
	"time"
)

const (
	RequestIdHeader     = "X-Request-Id"
)

type WebRouter struct {
	*mux.Router
	AuthToken *Token
}

// Writes Location header to response writer and sets specified status
func httpRedirect(resp http.ResponseWriter, url string, status int) {
	resp.Header().Add("Location", url)
	resp.WriteHeader(status)
}

func (router *WebRouter) createFuncMap() template.FuncMap {
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

func (router *WebRouter) parseLayout() *template.Template {
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

func (router *WebRouter) createPartial(names ...string) Partial {
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

func (router *WebRouter) Page(templateDataFn func(map[string]string, *http.Request) interface{}, names ...string) func(http.ResponseWriter, *http.Request) {
	tpl := router.createPartial(names...)

	return func(resp http.ResponseWriter, req *http.Request) {
		data := NewPartialData(templateDataFn(mux.Vars(req), req), req)

		tpl.MustRenderWithData(resp, data)
	}
}

func ListenAndServe(cfg common.Config, routerConfigurer func(*WebRouter)) {
	router := &WebRouter{Router: mux.NewRouter()}
	router.Use(requestIdMiddleware)
	router.Use(loggerMiddleware)
	if cfg.UseAuth {
		auth := &authMiddleware{
			router: router,
			config: cfg,
		}
		router.Use(auth.Middleware)

		authHandler := authController{cfg, router.createPartial("auth")}
		router.HandleFunc("/auth/success", authHandler.HandleSendSuccess)
		router.HandleFunc("/auth/set", authHandler.HandleTokenSet)
		router.HandleFunc("/auth/logout", authHandler.HandleLogout).Methods(http.MethodPost).Name("Logout")
		router.HandleFunc("/auth", authHandler.HandleAuthStart)
	}
	router.NotFoundHandler = http.HandlerFunc(router.Page(NilTemplateData, "404"))

	routerConfigurer(router)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.InsecurePort),
		Handler:      faviconMiddleware(cssMapRejectorMiddleware(router)),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	var listenErr error
	if cfg.IsTLS() {
		server.Addr = fmt.Sprintf(":%d", cfg.SecurePort)

		log.Printf("Server started at %s (secure)\n", server.Addr)
		listenErr = server.ListenAndServeTLS(cfg.TlsCert, cfg.TlsKey)
	} else {
		log.Printf("Server started at %s (insecure)\n", server.Addr)
		listenErr = server.ListenAndServe()
	}
	log.Fatalf("Server shutdown: %v\n", listenErr)
}

func NilTemplateData(vars map[string]string, req *http.Request) interface{} {
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

func cssMapRejectorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if strings.HasSuffix(request.RequestURI, ".css.map") {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		logger := common.NewRequestLogger(r)
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

		ctx := context.WithValue(r.Context(), common.RequestIdContextKey, requestId)

		writer.Header().Set(RequestIdHeader, requestId)

		next.ServeHTTP(writer, r.WithContext(ctx))
	})
}
