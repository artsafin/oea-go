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

func layout(router *mux.Router) *template.Template {
	funcMap := template.FuncMap{
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

	return template.Must(template.New("layout").Funcs(funcMap).Parse(string(common.MustAsset("resources/layout.go.html"))))
}

type templateGlobals struct {
	Router *mux.Router
}

func partial(names []string, globals *templateGlobals, templateDataFn func(map[string]string, *http.Request) interface{}) func(http.ResponseWriter, *http.Request) {
	tpl := layout(globals.Router)
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

func listenAndServe(routerConfigurer func(*mux.Router)) {
	router := mux.NewRouter()
	router.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "image/png")
		writer.WriteHeader(http.StatusOK)
		writer.Write(common.MustAsset("resources/icon.png"))
	})
	routerConfigurer(router)

	log.Println("Server started")

	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Request:", r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())

			router.ServeHTTP(w, r)
		}),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := server.ListenAndServe()

	log.Fatal("Server shutdown", err)
}

func nilTemplateData(vars map[string]string, req *http.Request) interface{} {
	return nil
}
