package main

//go:generate go-bindata -o "common/bindata.go" -pkg "common" resources/ resources/partials/

import (
	"flag"
	"oea-go/employee"
	"oea-go/office"
	"os"
	"strings"
	"time"
)

const configTimeout = 2 * time.Second

func main() {
	etcdAddr := flag.String("etcd", "", "Addresses of etcd cluster, separater by comma")
	flag.Parse()
	if *etcdAddr == "" {
		flag.Usage()
		os.Exit(1)
	}
	cfg := loadEtcdConfig(strings.Split(*etcdAddr, ","))

	/*
		/ - index page
			/office - list of [status=""] office invoices in sidebar
				/2020-02%23023 - approval post
					/invoice - downloads invoice
			/employee - last 2 months with invoices
				/2020-02 - list of payslip and financial posts per each employee
					/invoice - download invoice files in .zip
	*/
	officeHandler := office.NewHandler(&cfg)
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
