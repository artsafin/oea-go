package main

import (
	"flag"
	"log"
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
	cfg := common.NewConfig(AppVersion)
	etcd := common.NewEtcdService(strings.Split(*etcdAddr, ","))
	configErr := common.FillConfigFromEtcd(&cfg, etcd)
	if configErr != nil {
		log.Fatalf("error loading config: %v\n", configErr)
	}

	if *verbose {
		cfg.DumpNonSecretParameters(os.Stdout)
	}

	officeHandler := office.NewHandler(&cfg, etcd)
	employeesHandler := employee.NewHandler(&cfg)
	web.ListenAndServe(cfg, func(router *web.Engine) {
		if cfg.UseAuth {
			authWare := &auth.Middleware{
				Router: router,
				Config: cfg,
			}
			router.Use(authWare.MiddlewareFunc)

			authHandler := auth.NewHandler(&cfg, router.CreatePartial("auth"), etcd)
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
