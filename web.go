package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"oea-go/common"
	"time"
)

const AuthCookieName = "a"

type webRouter struct {
	*mux.Router
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
	}
}

func (router *webRouter) parseLayout() *template.Template {
	return template.Must(
		template.
			New("layout").
			Funcs(router.createFuncMap()).
			Parse(string(common.MustAsset("resources/layout.go.html"))))
}

type templateGlobals struct {
	Router *mux.Router
}

func (router *webRouter) partial(templateDataFn func(map[string]string, *http.Request) interface{}, names ...string) func(http.ResponseWriter, *http.Request) {
	tpl := router.parseLayout()
	for _, name := range names {
		tpl = template.Must(tpl.Parse(string(common.MustAsset(fmt.Sprintf("resources/partials/%s.go.html", name)))))
	}

	return func(resp http.ResponseWriter, req *http.Request) {
		tplData := struct {
			Page interface{}
		}{
			Page: templateDataFn(mux.Vars(req), req),
		}

		if err := tpl.Execute(resp, tplData); err != nil {
			panic(err)
		}
	}
}

func listenAndServe(routerConfigurer func(*webRouter)) {
	router := &webRouter{mux.NewRouter()}
	auth := &AuthMiddleware{
		router: router,
	}
	router.Use(faviconMiddleware)
	router.Use(loggerMiddleware)
	router.Use(auth.Middleware)
	router.NotFoundHandler = http.HandlerFunc(router.partial(nilTemplateData, "404"))

	routerConfigurer(router)

	log.Println("Server started")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	err := server.ListenAndServe()

	log.Fatal("Server shutdown", err)
}

func nilTemplateData(vars map[string]string, req *http.Request) interface{} {
	return nil
}

type AuthMiddleware struct {
	router *webRouter
}

func (auth *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, cookieErr := r.Cookie(AuthCookieName)
		if cookieErr != nil {
			log.Println("Forbidden:", r.Method, r.RequestURI)
			auth.router.partial(nilTemplateData, "403")(w, r)
			return
		}

		log.Printf("Auth cookie found: %v\n", authCookie)

		next.ServeHTTP(w, r)
	})
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
		log.Println("Request:", r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())

		next.ServeHTTP(writer, r)
	})
}
