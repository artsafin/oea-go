package main

//go:generate go-bindata -o "common/bindata.go" -pkg "common" resources/ resources/partials/

import (
	"github.com/spf13/viper"
	"oea-go/common"
	"oea-go/employee"
	"oea-go/office"
)

func prepareConfig() common.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("out_dir", "generated")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var c common.Config
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	return c
}

func main() {
	cfg := prepareConfig()

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
	listenAndServe(func(router *webRouter) {
		GET := router.Methods("GET").Subrouter()

		GET.HandleFunc("/office/{invoice:.+}/invoice", officeHandler.DownloadInvoice).Name("OfficeDownloadInvoice")
		GET.HandleFunc("/office/{invoice:.+}", router.partial(officeHandler.ShowInvoiceData, "office_invoice_data")).Name("OfficeShowInvoiceData")
		GET.HandleFunc("/office", router.partial(officeHandler.Home, "office")).Name("OfficeHome")

		GET.HandleFunc("/employees", router.partial(employeesHandler.Home, "employees")).Name("EmployeesHome")
		GET.HandleFunc("/employee/{month}", router.partial(employeesHandler.Month, "employees", "employees_month")).Name("EmployeesMonth")
		GET.HandleFunc("/employee/{month}/invoices", employeesHandler.DownloadAllInvoices).Name("EmployeesDownloadAllInvoices")
		GET.HandleFunc("/employee/{month}/{employee}/invoice", employeesHandler.DownloadInvoice).Name("EmployeesDownloadInvoice")

		GET.HandleFunc("/", router.partial(nilTemplateData, "index"))
	})
}
