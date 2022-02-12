package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
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

type TemplateData struct {
	Page       interface{}
	RequestURI string
	Version    string
}

func (t *TemplateData) WithData(data interface{}) *TemplateData {
	t.Page = data
	return t
}

func NewTemplateData(version string, data interface{}, r *http.Request) *TemplateData {
	return &TemplateData{
		Page:       data,
		RequestURI: r.RequestURI,
		Version:    version,
	}
}

// HttpRedirect writes Location header to response writer and sets specified status
func HttpRedirect(resp http.ResponseWriter, url string, status int) {
	resp.Header().Add("Location", url)
	resp.WriteHeader(status)
}

type Engine struct {
	*mux.Router
	AuthToken *authtoken.Token
	version   string
	logger    *zap.SugaredLogger
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

func (e *Engine) NewTemplateWithLayout(names ...string) *template.Template {
	tpl := e.parseLayout()
	for _, name := range names {
		resPath := fmt.Sprintf("assets/partials/%s.go.html", name)
		tpl = resources.MustParseTemplate(tpl, resPath)
	}

	return tpl
}

func (e *Engine) Page(pageDataFn func(map[string]string, *http.Request) interface{}, names ...string) func(http.ResponseWriter, *http.Request) {
	tpl := e.NewTemplateWithLayout(names...)

	return func(resp http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		var pageData interface{} = nil
		if pageDataFn != nil {
			pageData = pageDataFn(vars, req)
		}
		data := NewTemplateData(e.version, pageData, req)

		if err := tpl.Execute(resp, data); err != nil {
			e.logger.Errorf("error executing template %v: %v. Vars: %+v", names, err, vars)
		}
	}
}

func ListenAndServe(cfg config.Config, logger *zap.SugaredLogger, routerConfigurer func(*Engine)) {
	router := &Engine{Router: mux.NewRouter(), version: cfg.AppVersion, logger: logger}
	router.Use(requestIdMiddleware)
	router.Use(loggerMiddleware{logger: logger}.MiddlewareFunc)
	router.NotFoundHandler = http.HandlerFunc(router.Page(nil, "404"))

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
