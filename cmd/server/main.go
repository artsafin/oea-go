package main

import (
	"flag"
	"log"
	"oea-go/internal/common"
	"oea-go/internal/employee"
	"oea-go/internal/office"
	"oea-go/internal/ui"
	"os"
	"strings"
)

// invalidate all tokens

func main() {
	verbose := flag.Bool("v", false, "Be more verbose")
	etcdAddr := flag.String("etcd", "", "Addresses of etcd cluster, separater by comma")
	flag.Parse()
	if *etcdAddr == "" {
		flag.Usage()
		os.Exit(1)
	}
	cfg := common.NewConfig()
	etcd := common.NewEtcdConnection(strings.Split(*etcdAddr, ","))
	configErr := common.FillConfigFromEtcd(cfg, etcd)
	if configErr != nil {
		log.Fatalf("error loading config: %v\n", configErr)
	}

	if *verbose {
		cfg.DumpNonSecretParameters(os.Stdout)
	}

	officeHandler := office.NewHandler(cfg, etcd)
	employeesHandler, err := employee.NewHandler(cfg)
	if err != nil {
		log.Fatalf("error creating coda client: %v\n", err)
	}
	ui.ListenAndServe(cfg, func(router *ui.WebRouter) {
		GET := router.Methods("GET").Subrouter()

		GET.HandleFunc("/office/{invoice:.+}/invoice", officeHandler.DownloadInvoice).Name("OfficeDownloadInvoice")
		GET.HandleFunc("/office/{invoice:.+}", router.Page(officeHandler.ShowInvoiceData, "office_invoice_data")).Name("OfficeShowInvoiceData")
		GET.HandleFunc("/office", router.Page(officeHandler.Home, "office")).Name("OfficeHome")

		GET.HandleFunc("/employees", router.Page(employeesHandler.Home, "employees")).Name("EmployeesHome")
		GET.HandleFunc("/employee/{month}", router.Page(employeesHandler.Month, "employees", "employees_month")).Name("EmployeesMonth")
		GET.HandleFunc("/employee/{month}/invoices", employeesHandler.DownloadAllInvoices).Name("EmployeesDownloadAllInvoices")
		GET.HandleFunc("/employee/{month}/{employee}/invoice", employeesHandler.DownloadInvoice).Name("EmployeesDownloadInvoice")

		GET.HandleFunc("/", router.Page(ui.NilTemplateData, "index"))
	})
}
