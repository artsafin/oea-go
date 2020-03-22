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
	officeHandler := office.Handler{Config: &cfg}
	employeesHandler := employee.Handler{Config: &cfg}
	listenAndServe(func(router *mux.Router) {
		GET := router.Methods("GET").Subrouter()

		globals := &templateGlobals{GET}

		GET.HandleFunc("/office/{invoice:.+}/invoice", officeHandler.DownloadInvoice).Name("OfficeDownloadInvoice")
		GET.HandleFunc("/office/{invoice:.+}", partial("office_invoice_data", globals, officeHandler.ShowInvoiceData)).Name("OfficeShowInvoiceData")
		GET.HandleFunc("/office", partial("office", globals, officeHandler.Home)).Name("OfficeHome")

		GET.HandleFunc("/employees", partial("employees", globals, employeesHandler.Home)).Name("EmployeesHome")
		router.HandleFunc("/employee/{month}", partial("employees_month", globals, nilTemplateData)).Name("EmployeesMonth")
		//router.HandleFunc("/employee/{month}/invoice", partial("employee", employee.MainController))

		GET.HandleFunc("/", partial("index", globals, nilTemplateData))

		router.NotFoundHandler = http.HandlerFunc(partial("404", globals, nilTemplateData))
	})
}
