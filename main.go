package main

//go:generate go-bindata -o "common/bindata.go" -pkg "common" resources/ resources/partials/

import (
	"flag"
	"log"
	"oea-go/common"
	"oea-go/employee"
	"oea-go/office"
	"os"
	"strings"
)

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
	configErr := common.FillConfigFromEtcd(&cfg, etcd)
	if configErr != nil {
		log.Fatalf("error loading config: %v\n", configErr)
	}

	if *verbose {
		cfg.DumpNonSecretParameters(os.Stdout)
	}

	officeHandler := office.NewHandler(&cfg, etcd)
	employeesHandler := employee.NewHandler(&cfg)
	listenAndServe(cfg, func(router *webRouter) {
		GET := router.Methods("GET").Subrouter()

		GET.HandleFunc("/office/{invoice:.+}/invoice", officeHandler.DownloadInvoice).Name("OfficeDownloadInvoice")
		GET.HandleFunc("/office/{invoice:.+}", router.page(officeHandler.ShowInvoiceData, "office_invoice_data")).Name("OfficeShowInvoiceData")
		GET.HandleFunc("/office", router.page(officeHandler.Home, "office")).Name("OfficeHome")

		GET.HandleFunc("/employees", router.page(employeesHandler.Home, "employees")).Name("EmployeesHome")
		GET.HandleFunc("/employee/{month}", router.page(employeesHandler.Month, "employees", "employees_month")).Name("EmployeesMonth")
		GET.HandleFunc("/employee/{month}/invoices", employeesHandler.DownloadAllInvoices).Name("EmployeesDownloadAllInvoices")
		GET.HandleFunc("/employee/{month}/{employee}/invoice", employeesHandler.DownloadInvoice).Name("EmployeesDownloadInvoice")

		GET.HandleFunc("/", router.page(nilTemplateData, "index"))
	})
}
