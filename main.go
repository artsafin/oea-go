package main

import (
	"flag"
	"go.uber.org/zap"
	"net/http"
	"oea-go/internal/auth"
	"oea-go/internal/common"
	"oea-go/internal/employee"
	"oea-go/internal/office"
	"oea-go/internal/web"
	"os"
	"strings"
)

// Populated by -ldflags "-X main.AppVersion=..."
var AppVersion string

func main() {
	verbose := flag.Bool("v", false, "Be more verbose")
	etcdAddr := flag.String("etcd", "", "Addresses of etcd cluster, separater by comma")
	flag.Parse()
	if *etcdAddr == "" {
		flag.Usage()
		os.Exit(1)
	}
	baseLogger, _ := zap.NewDevelopment()
	logger := baseLogger.Sugar()
	cfg := common.NewConfig(AppVersion, logger)
	etcd := common.NewEtcdService(strings.Split(*etcdAddr, ","), logger)
	configErr := common.FillConfigFromEtcd(&cfg, etcd, logger)
	if configErr != nil {
		logger.Fatalw("error loading config", "error", configErr)
		os.Exit(1)
	}

	if *verbose {
		cfg.DumpNonSecretParameters(os.Stdout)
	}

	officeHandler := office.NewHandler(&cfg, etcd, logger)
	employeesHandler := employee.NewHandler(&cfg, logger)
	web.ListenAndServe(cfg, logger, func(router *web.Engine) {
		if cfg.UseAuth {
			authWare := &auth.Middleware{
				Router: router,
				Config: cfg,
				Logger: logger,
			}
			router.Use(authWare.MiddlewareFunc)

			authHandler := auth.NewHandler(&cfg, router.CreatePartial("auth"), etcd, logger)
			router.HandleFunc("/auth/success", authHandler.HandleSendSuccess)
			//router.HandleFunc("/auth/set", authHandler.HandleTokenSet)
			router.HandleFunc("/auth/set", authHandler.HandleBegin2FA)
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
