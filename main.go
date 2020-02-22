package main

//go:generate tpl/jade -basedir tpl -d tpl -pkg tpl -writer post.jade

import (
	"fmt"
	"os"
	"sync"

	"github.com/artsafin/ofa-go/dto"
	"github.com/artsafin/ofa-go/tpl"
	"github.com/spf13/viper"
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
	if !viper.IsSet("api_token") || !viper.IsSet("base_uri") {
		panic("Config parameters are not set")
	}

	ccl := NewCodaClient(viper.GetString("base_uri"), viper.GetString("api_token"))

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
	os.Mkdir(outDir, os.FileMode(0777))

	file, err := os.Create(fmt.Sprintf("%s/%s.html", outDir, invoice.Filename))
	defer file.Close()
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func() {
		tpl.Post(invoice, expensesByCategory, history, file)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		RenderTemplate(outDir, "templates/invoice_template.xlsx", invoice)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Finished")
}
