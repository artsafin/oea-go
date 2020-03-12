package main

//go:generate go-bindata resources/

import (
	"fmt"
	"github.com/spf13/viper"
	"oea-go/common"
	"oea-go/office"
	"os"
	"os/user"
	"sync"
	"text/template"
	"time"

	officeDto "oea-go/office/dto"
)

func prepareConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("out_dir", "generated")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if !viper.IsSet("api_token_of") || !viper.IsSet("base_uri") || !viper.IsSet("doc_id_of") {
		panic("Config parameters are not set")
	}
}

func loadOfficeData(baseUri, apiTokenOf, docId string) OfficeTemplateData {
	officeReq := office.Requests{
		Client: common.NewCodaClient(baseUri, apiTokenOf),
		DocId:  docId,
	}

	var invoice *officeDto.Invoice
	var expensesByCategory officeDto.ExpenseGroupMap
	var history *officeDto.History

	wg := sync.WaitGroup{}

	invoiceID := officeReq.WaitForInvoice()

	wg.Add(1)
	go func() {
		fmt.Println("Loading invoice...")
		invoice = officeReq.GetInvoice(invoiceID)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Loading expenses...")
		expenses := officeReq.GetExpenses(invoiceID)
		expensesByCategory = officeDto.GroupExpensesByCategory(expenses)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Loading history...")
		history = officeReq.GetHistory()
		wg.Done()
	}()

	wg.Wait()

	return OfficeTemplateData{
		*invoice, expensesByCategory, *history,
	}
}

func main() {
	prepareConfig()

	officeData := loadOfficeData(viper.GetString("base_uri"), viper.GetString("api_token_of"), viper.GetString("doc_id_of"))

	outDir := viper.GetString("out_dir")
	if err := os.MkdirAll(outDir, os.FileMode(0777)); err != nil {
		panic(fmt.Sprintf("Unable to create target dir: %s", err))
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.html", outDir, officeData.Invoice.Filename))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file")
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("Building post...")

		tpl, err := template.New("post").Parse(string(MustAsset("resources/post.go.html")))
		if err != nil {
			panic(err)
		}
		currUser, userErr := user.Current()
		if userErr != nil {
			panic(userErr)
		}

		err = tpl.Execute(file, TemplateData{
			Timestamp: time.Now(),
			Author:    currUser,
			Office:    officeData,
		})

		if err != nil {
			panic(err)
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Building excel...")
		RenderExcelTemplate(outDir, MustAsset("resources/invoice_template.xlsx"), &officeData.Invoice)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Finished")
}
