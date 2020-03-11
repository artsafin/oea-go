package main

//go:generate go-bindata resources/

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
	"ofa-go/dto"
	"text/template"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("out_dir", "generated")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if !viper.IsSet("api_token") || !viper.IsSet("base_uri") || !viper.IsSet("doc_id") {
		panic("Config parameters are not set")
	}

	ccl := NewCodaClient(viper.GetString("base_uri"), viper.GetString("api_token"), viper.GetString("doc_id"))

	var invoice *dto.Invoice
	var expensesByCategory dto.ExpenseGroupMap
	var history *dto.History

	wg := sync.WaitGroup{}

	invoiceID := ccl.WaitForInvoice()

	wg.Add(1)
	go func() {
		invoice = ccl.getInvoice(invoiceID)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		expenses := ccl.getExpenses(invoiceID)
		expensesByCategory = dto.GroupExpensesByCategory(expenses)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		history = ccl.getHistory()
		wg.Done()
	}()

	wg.Wait()

	outDir := viper.GetString("out_dir")
	if err := os.MkdirAll(outDir, os.FileMode(0777)); err != nil {
		panic(fmt.Sprintf("Unable to create target dir: %s", err))
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.html", outDir, invoice.Filename))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file")
		}
	}()

	wg.Add(1)
	go func() {
		tpl, err := template.New("post").Parse(string(MustAsset("resources/post.go.html")))
		if err != nil {
			panic(err)
		}
		err = tpl.Execute(file, struct {
			Invoice       dto.Invoice
			ExpenseGroups dto.ExpenseGroupMap
			History       dto.History
		}{*invoice, expensesByCategory, *history})

		if err != nil {
			panic(err)
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		RenderExcelTemplate(outDir, MustAsset("resources/invoice_template.xlsx"), invoice)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Finished")
}
