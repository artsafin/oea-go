package dto

type InvoicesColumns struct {
	No               string
	Status           string
	Number           string
	Date             string
	Filename         string
	HourRate         string
	ApprovalLink     string
	EurFixedRate     string
	EurRateWorst     string
	ExpensesRub      string
	ExpensesEur      string
	ReturnOfRounding string
	Subtotal         string
	HourRateRounding string
	TotalEur         string
	Hours            string
	ActuallySpent    string
	PendingSpend     string
	Balance          string
	PrevInvoice      string
}

type ExpensesColumns struct {
	ID              string
	Invoice         string
	Subject         string
	Category        string
	Comment         string
	AmountRub       string
	AmountEur       string
	Status          string
	ActuallySpent   string
	RejectionReason string
	PendingSpend    string
	Balance         string
	CashFlow        string
	LastCashOutDate string
	Group           string
}

type IdOnly struct {
	Id string
}

type InvoicesDecl struct {
	IdOnly
	Cols InvoicesColumns
}

type ExpensesDecl struct {
	IdOnly
	Cols ExpensesColumns
}

type CodaViewsDecl struct {
	PlannedInvoices string
	PlannedExpenses string
}

type CodaFormulasDecl struct {
	LastInvoice string
}

type IdStruct struct {
	Invoices     InvoicesDecl
	Expenses     ExpensesDecl
	CodaViews    CodaViewsDecl
	CodaFormulas CodaFormulasDecl
}

var Ids IdStruct

func init() {
	Ids = IdStruct{
		Invoices: InvoicesDecl{
			IdOnly{"grid-H_lQoXT4Hn"},
			InvoicesColumns{
				No:               "c-1FLrjxKHe5",
				Status:           "c-rqUgKu_1z_",
				Number:           "c-qw6CtPRr6R",
				Date:             "c-r6ev-Hfy0R",
				Filename:         "c-gVFw0tg27k",
				HourRate:         "c-wu0qKZLU0Z",
				ApprovalLink:     "c-mINTWSuok_",
				EurFixedRate:     "c-39yF51XkOy",
				EurRateWorst:     "c-eqfmnjpn6E",
				ExpensesRub:      "c-LnWC3NrrYo",
				ExpensesEur:      "c-Bj-zR8QBZP",
				ReturnOfRounding: "c-cEdaY7AsWT",
				Subtotal:         "c-EfZ0XNkbSp",
				HourRateRounding: "c-wbNLVBU5oW",
				TotalEur:         "c-gywxqGt4uK",
				Hours:            "c-2kpnqmokJ5",
				ActuallySpent:    "c-XY-LhijBo3",
				PendingSpend:     "c-bccwx_DDwv",
				Balance:          "c-MW_JQd5O0b",
				PrevInvoice:      "c-RWqrc6Q-Zf",
			},
		},
		Expenses: ExpensesDecl{
			IdOnly{"grid-j1yvp-c7Xq"},
			ExpensesColumns{
				ID:              "c-NIDES-BkAW",
				Invoice:         "c-_Vwi1N_WO9",
				Subject:         "c-dWrUpt7GHg",
				Category:        "c-HgZFvbcM6u",
				Comment:         "c-hLtZdaFaOE",
				AmountRub:       "c-NWO2JHrrdw",
				AmountEur:       "c-85jKBvNSVF",
				Status:          "c-TlBKoSAJmV",
				ActuallySpent:   "c-otkrTpI7VA",
				RejectionReason: "c-OwY64F1V4B",
				PendingSpend:    "c-MTRmKDu-3A",
				Balance:         "c-GAIziVHBvd",
				LastCashOutDate: "c-BhuJTIsQEP",
				Group:           "c-18G6IQXBvd",
			},
		},
		CodaViews: CodaViewsDecl{
			PlannedInvoices: "table-0oiNc6Kkw8",
			PlannedExpenses: "table-NvLNctjziJ",
		},
		CodaFormulas: CodaFormulasDecl{
			LastInvoice: "f-6rR6RS0Un_",
		},
	}
}
