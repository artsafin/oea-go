package main

import (
	"bytes"
	"context"
	"flag"
	"github.com/artsafin/coda-go-client/codaapi"
	"go.uber.org/zap"
	"net/http"
	"oea-go/internal/auth"
	"oea-go/internal/config"
	"oea-go/internal/db"
	emplschema "oea-go/internal/employee/codaschema"
	employeeweb "oea-go/internal/employee/web"
	officeschema "oea-go/internal/office/codaschema"
	officeweb "oea-go/internal/office/web"
	"oea-go/internal/web"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
)

// AppVersion Populated by -ldflags "-X main.AppVersion=..."
var AppVersion string

func parseStackTrace(stackBs []byte) interface{} {
	parsRem := regexp.MustCompile(`\([^)]+\)$`)
	fileRem := regexp.MustCompile(` \+0x.+$`)

	stackBsLines := bytes.Split(stackBs, []byte{'\n'})
	var reltrace []string
	for _, lineBs := range stackBsLines {
		if bytes.Contains(lineBs, []byte("office")) || bytes.Contains(lineBs, []byte("employee")) {
			lineBs = parsRem.ReplaceAll(lineBs, []byte{})
			lineBs = fileRem.ReplaceAll(lineBs, []byte{})
			reltrace = append(reltrace, string(lineBs))
		}
	}

	return strings.ReplaceAll(strings.Join(reltrace, "\n"), "\n\t", " ")
}

func main() {
	verbose := flag.Bool("v", false, "Be more verbose")
	kvAddr := flag.String("s", "localhost:6379", "Storage address")
	flag.Parse()

	cfg := config.NewDefaultConfig(AppVersion, *kvAddr)

	baseLogger, _ := zap.NewDevelopment(zap.WithCaller(false))
	logger := baseLogger.Sugar()

	configErr := cfg.LoadFromEnvAndValidate()
	if configErr != nil {
		logger.Fatalf("error loading config: %v", configErr)
	}

	if *verbose {
		cfg.DumpNonSecretParameters(os.Stdout)
	}

	storageKeys, storageErr := db.NewStorage(cfg.StorageAddr).Keys()
	if storageErr != nil {
		logger.Fatalf("storage at %v is unavailable: %v", cfg.StorageAddr, storageErr)
	} else {
		logger.Debugf("storage is alive (%v keys)", len(storageKeys))
	}

	loggingMiddleware := codaapi.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		if cfg.IsDebug {
			stack := parseStackTrace(debug.Stack())

			logger.Infof("CODA %v %v\n%v\n", req.Method, req.URL, stack)
		}
		return nil
	})

	officeCoda, err := officeschema.NewCodaDocument(cfg.BaseUri, cfg.ApiTokenOf, cfg.DocIdOf, loggingMiddleware)
	if err != nil {
		logger.Fatalf("could not create a coda client (office): %v", err)
	}
	officeHandler := officeweb.NewHandlers(cfg.FilesDir, officeCoda, logger)

	employeesCoda, err := emplschema.NewCodaDocument(cfg.BaseUri, cfg.ApiTokenEm, cfg.DocIdEm, loggingMiddleware)
	if err != nil {
		logger.Fatalf("could not create a coda client (empl): %v", err)
	}
	employeesHandler := employeeweb.NewHandlers(employeesCoda, logger)

	web.ListenAndServe(cfg, logger, func(router *web.Engine) {
		if cfg.UseAuth {
			authWare := &auth.Middleware{
				Router: router,
				Config: cfg,
				Logger: logger,
			}
			router.Use(authWare.MiddlewareFunc)

			authHandler := auth.NewHandler(&cfg, router.NewTemplateWithLayout("auth"), logger, router.Router)
			router.HandleFunc("/auth/sent", authHandler.HandleFirstFactorSendSuccess)
			//router.HandleFunc("/auth/set", authHandler.HandleTokenSet)
			router.HandleFunc("/auth/twofa", authHandler.HandleCheckSecondFactor).
				Methods(http.MethodPost).
				Name("AuthCheck2FA")
			router.HandleFunc("/auth/set", authHandler.HandleBeginSecondFactor)
			router.HandleFunc("/auth/logout", authHandler.HandleLogout).
				Methods(http.MethodPost).
				Name("Logout")
			router.HandleFunc("/auth", authHandler.HandleAuthStart)
		}

		GET := router.Methods("GET").Subrouter()

		GET.HandleFunc("/office/{invoice:.+}/invoice", officeHandler.DownloadInvoice).Name("OfficeDownloadInvoice")
		GET.HandleFunc("/office/{invoice:.+}", router.Page(officeHandler.ShowInvoiceData, "office_invoice_data")).Name("OfficeShowInvoiceData")
		GET.HandleFunc("/office", router.Page(officeHandler.Home, "office")).Name("OfficeHome")

		GET.HandleFunc("/employees", router.Page(employeesHandler.Home, "employees")).Name("EmployeesHome")
		GET.HandleFunc("/employee/{month}", router.Page(employeesHandler.Month, "employees", "employees_month")).Name("EmployeesMonth")
		GET.HandleFunc("/employee/{month}/invoices", employeesHandler.DownloadAllInvoices).Name("EmployeesDownloadAllInvoices")
		GET.HandleFunc("/employee/{month}/payroll.xlsx", employeesHandler.DownloadPayrollReport).Name("EmployeesDownloadPayrollReport")
		GET.HandleFunc("/employee/{month}/payroll.txt", employeesHandler.DownloadHellenicPayroll).Name("EmployeesDownloadHellenicPayroll")
		GET.HandleFunc("/employee/{month}/{employee}/invoice", employeesHandler.DownloadInvoice).Name("EmployeesDownloadInvoice")

		GET.HandleFunc("/", router.Page(nil, "index"))
	})
}
