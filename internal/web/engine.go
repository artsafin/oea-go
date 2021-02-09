package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"oea-go/internal/auth"
	"oea-go/internal/common"
	"time"
)

const (
	RequestIdHeader = "X-Request-Id"
)

type Engine struct {
	*mux.Router
	AuthToken *auth.Token
}

// Writes Location header to response writer and sets specified status
func HttpRedirect(resp http.ResponseWriter, url string, status int) {
	resp.Header().Add("Location", url)
	resp.WriteHeader(status)
}

func (router *Engine) createFuncMap() template.FuncMap {
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
		"authToken": func() *auth.Token {
			return router.AuthToken
		},
	}
}

func (router *Engine) parseLayout() *template.Template {
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

func (router *Engine) CreatePartial(names ...string) Partial {
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

func (router *Engine) Page(templateDataFn func(map[string]string, *http.Request) interface{}, names ...string) func(http.ResponseWriter, *http.Request) {
	tpl := router.CreatePartial(names...)

	return func(resp http.ResponseWriter, req *http.Request) {
		data := NewPartialData(templateDataFn(mux.Vars(req), req), req)

		tpl.MustRenderWithData(resp, data)
	}
}

func ListenAndServe(cfg common.Config, logger *zap.SugaredLogger, routerConfigurer func(*Engine)) {
	router := &Engine{Router: mux.NewRouter()}
	router.Use(requestIdMiddleware)
	router.Use(loggerMiddleware{logger: logger}.MiddlewareFunc)
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

		logger.Infof("Server started at %s (secure)\n", server.Addr)
		listenErr = server.ListenAndServeTLS(cfg.TlsCert, cfg.TlsKey)
	} else {
		logger.Infof("Server started at %s (insecure)\n", server.Addr)
		listenErr = server.ListenAndServe()
	}
	logger.Fatalf("Server shutdown: %v\n", listenErr)
}

func NilTemplateData(vars map[string]string, req *http.Request) interface{} {
	return nil
}
