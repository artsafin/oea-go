package main

import (
	"flag"
	"go.uber.org/zap"
	"net/http"
	"oea-go/internal/auth"
	"oea-go/internal/common/config"
	"oea-go/internal/employee"
	"oea-go/internal/office"
	"oea-go/internal/web"
	"os"
)

// Populated by -ldflags "-X main.AppVersion=..."
var AppVersion string

func main() {
	verbose := flag.Bool("v", false, "Be more verbose")
	flag.Parse()

	baseLogger, _ := zap.NewDevelopment(zap.WithCaller(false))
	logger := baseLogger.Sugar()
	cfg := config.NewDefaultConfig(AppVersion)

	configErr := cfg.LoadFromEnvAndValidate()
	if configErr != nil {
		logger.Fatalf("error loading config: %v", configErr)
	}

	if *verbose {
		cfg.DumpNonSecretParameters(os.Stdout)
	}

	officeHandler := office.NewHandler(&cfg, logger)
	employeesHandler := employee.NewHandler(&cfg, logger)
	web.ListenAndServe(cfg, logger, func(router *web.Engine) {
		if cfg.UseAuth {
			authWare := &auth.Middleware{
				Router: router,
				Config: cfg,
				Logger: logger,
			}
			router.Use(authWare.MiddlewareFunc)

			authHandler := auth.NewHandler(&cfg, router.CreatePartial("auth"), logger)
			router.HandleFunc("/auth/sent", authHandler.HandleFirstFactorSendSuccess)
			//router.HandleFunc("/auth/set", authHandler.HandleTokenSet)
			router.HandleFunc("/auth/twofa", authHandler.HandleCheckSecondFactor)
			router.HandleFunc("/auth/set", authHandler.HandleBeginSecondFactor)
			router.HandleFunc("/auth/logout", authHandler.HandleLogout).Methods(http.MethodPost).Name("Logout")
			router.HandleFunc("/auth", authHandler.HandleAuthStart)
		}

		GET := router.Methods("GET").Subrouter()

		GET.HandleFunc("/office/{invoice:.+}/invoice", officeHandler.DownloadInvoice).Name("OfficeDownloadInvoice")
		GET.HandleFunc("/office/{invoice:.+}", router.Page(officeHandler.ShowInvoiceData, "office_invoice_data")).Name("OfficeShowInvoiceData")
		GET.HandleFunc("/office", router.Page(officeHandler.Home, "office")).Name("OfficeHome")

		GET.HandleFunc("/employees", router.Page(employeesHandler.Home, "employees")).Name("EmployeesHome")
		GET.HandleFunc("/employee/{month}", router.Page(employeesHandler.Month, "employees", "employees_month")).Name("EmployeesMonth")
		GET.HandleFunc("/employee/{month}/invoices", employeesHandler.DownloadAllInvoices).Name("EmployeesDownloadAllInvoices")
		GET.HandleFunc("/employee/{month}/payrollreport", employeesHandler.DownloadPayrollReport).Name("EmployeesDownloadPayrollReport")
		GET.HandleFunc("/employee/{month}/{employee}/invoice", employeesHandler.DownloadInvoice).Name("EmployeesDownloadInvoice")

		GET.HandleFunc("/", router.Page(web.NilTemplateData, "index"))
	})
}
