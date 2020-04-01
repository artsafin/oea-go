package main

//go:generate go-bindata -o "common/bindata.go" -pkg "common" resources/ resources/partials/

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
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
	listenAndServe(func(router *mux.Router) {
		GET := router.Methods("GET").Subrouter()

		globals := &templateGlobals{GET}

		GET.HandleFunc("/office/{invoice:.+}/invoice", officeHandler.DownloadInvoice).Name("OfficeDownloadInvoice")
		GET.HandleFunc("/office/{invoice:.+}", partial([]string{"office_invoice_data"}, globals, officeHandler.ShowInvoiceData)).Name("OfficeShowInvoiceData")
		GET.HandleFunc("/office", partial([]string{"office"}, globals, officeHandler.Home)).Name("OfficeHome")

		GET.HandleFunc("/employees", partial([]string{"employees"}, globals, employeesHandler.Home)).Name("EmployeesHome")
		GET.HandleFunc("/employee/{month}", partial([]string{"employees", "employees_month"}, globals, employeesHandler.Month)).Name("EmployeesMonth")
		GET.HandleFunc("/employee/{month}/invoices", employeesHandler.DownloadAllInvoices).Name("EmployeesDownloadAllInvoices")
		GET.HandleFunc("/employee/{month}/{employee}/invoice", employeesHandler.DownloadInvoice).Name("EmployeesDownloadInvoice")

		GET.HandleFunc("/", partial([]string{"index"}, globals, nilTemplateData))

		router.NotFoundHandler = http.HandlerFunc(partial([]string{"404"}, globals, nilTemplateData))
	})
}
