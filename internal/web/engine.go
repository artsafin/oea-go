package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"oea-go/internal/auth/authtoken"
	"oea-go/internal/common/config"
	"oea-go/resources"
	"time"
)

const (
	RequestIdHeader = "X-Request-Id"
)

type Engine struct {
	*mux.Router
	AuthToken *authtoken.Token
	Version   string
}

// Writes Location header to response writer and sets specified status
func HttpRedirect(resp http.ResponseWriter, url string, status int) {
	resp.Header().Add("Location", url)
	resp.WriteHeader(status)
}

func (e *Engine) createFuncMap() template.FuncMap {
	return template.FuncMap{
		"link": func(routeName string, args ...string) string {
			var route *mux.Route
			var u *url.URL
			var err error
			if route = e.Get(routeName); route == nil {
				return ""
			}
			if u, err = route.URL(args...); err != nil {
				return ""
			}
			return u.String()
		},
		"styles": func() template.CSS {
			return template.CSS(resources.MustReadBytes("assets/bootstrap.min.css"))
		},
		"authToken": func() *authtoken.Token {
			return e.AuthToken
		},
	}
}

func (e *Engine) parseLayout() *template.Template {
	t := template.New("layout.go.html").Funcs(e.createFuncMap())
	return resources.MustParseTemplate(t, "assets/layout.go.html")
}

type Partial struct {
	*template.Template
}

type PartialData struct {
	Page       interface{}
	RequestURI string
	Version    string
}

func NewPartialData(version string, data interface{}, r *http.Request) *PartialData {
	return &PartialData{
		Page:       data,
		RequestURI: r.RequestURI,
		Version:    version,
	}
}

func (e *Engine) CreatePartial(names ...string) Partial {
	tpl := e.parseLayout()
	for _, name := range names {
		resPath := fmt.Sprintf("assets/partials/%s.go.html", name)
		tpl = resources.MustParseTemplate(tpl, resPath)
	}

	return Partial{tpl}
}

func (partial Partial) MustRenderWithData(wr io.Writer, data *PartialData) {
	if err := partial.Execute(wr, data); err != nil {
		panic(err)
	}
}

func (e *Engine) Page(templateDataFn func(map[string]string, *http.Request) interface{}, names ...string) func(http.ResponseWriter, *http.Request) {
	tpl := e.CreatePartial(names...)

	return func(resp http.ResponseWriter, req *http.Request) {
		data := NewPartialData(e.Version, templateDataFn(mux.Vars(req), req), req)

		tpl.MustRenderWithData(resp, data)
	}
}

func ListenAndServe(cfg config.Config, logger *zap.SugaredLogger, routerConfigurer func(*Engine)) {
	router := &Engine{Router: mux.NewRouter(), Version: cfg.AppVersion}
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
