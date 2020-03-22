package dto

type MonthsColumns struct {
	Month,
	WorkDays,
	Year,
	ID,
	PreviousMonthLink,
	PreviousMonth,
	IsCurrent string
}

type InvoicesColumns struct {
	Id                   string
	InvoiceNo            string
	Month                string
	Employee             string
	PreviousInvoice      string
	AmountRub            string
	EurRubExpected       string
	RequestedSubtotalEur string
	RoundingPrevMonEur   string
	Rounding             string
	AmountRequestedEur   string
	Hours                string
	EurRubActual         string
	AmountRubActual      string
	RateErrorRub         string
	CostOfDay            string
	BankTotalFees        string
	OpeningDateIp        string
	UnpaidDay            string
	CorrectionRefs       string
	CorrectionRub        string
	SalaryRub            string
	PatentRub            string
	TaxesRub             string
	PatentRefs           string
	TaxesRefs            string
	UnpaidDayRefs        string
	BaseSalary           string
	PaymentService       string
}

type CorrectionsColumns struct {
	Comment                  string
	AbsoluteCorrectionEur    string
	AbsoluteCorrectionRub    string
	EurRubExpected           string
	AbsCorrectionEurInRub    string
	PerDayType               string
	NumberOfDays             string
	CostOfDay                string
	PerDay                   string
	TotalCorrectionRub       string
	PaymentInvoice           string
	PerDayCoefficient        string
	PerDayCalculationInvoice string
	Display                  string
	Category                 string
}

type TaxesColumns struct {
	Invoice            string
	OpeningDateIp      string
	PeriodStart        string
	PeriodEnd          string
	AmountIpDays       string
	MedicalFund        string
	PensionFundFixed   string
	PensionFundPercent string
	Amount             string
}

type PatentColumns struct {
	Invoice       string
	OpeningPatent string
	PeriodEnd     string
	FullMonths    string
	AnnualCost    string
	PeriodCost    string
	Period        string
}

type IdOnly struct {
	Id string
}

type InvoicesDecl struct {
	IdOnly
	Cols InvoicesColumns
}

type MonthsDecl struct {
	IdOnly
	Cols MonthsColumns
}

type CorrectionsDecl struct {
	IdOnly
	Cols CorrectionsColumns
}

type TaxesDecl struct {
	IdOnly
	Cols TaxesColumns
}

type PatentDecl struct {
	IdOnly
	Cols PatentColumns
}

type CodaFormulasDecl struct {
	CurrentMonth string
}

type CodaViewsDecl struct {
	ThisMonthInvoices    string // table-_k4m2krdyp
	ThisMonthCorrections string // table-29Y7W0xzVA
	ThisMonthTaxes       string // table-X3zgg4jJ8-
	ThisMonthPatents     string // table-n2aIqJxO80
}

type IdStruct struct {
	Invoices     InvoicesDecl
	Corrections  CorrectionsDecl
	Taxes        TaxesDecl
	Patent       PatentDecl
	Months       MonthsDecl
	CodaFormulas CodaFormulasDecl
	//CodaViews    CodaViewsDecl
}

var Ids IdStruct

func init() {
	Ids = IdStruct{
		Invoices: InvoicesDecl{
			IdOnly{"grid-Wdy6Agpxou"},
			InvoicesColumns{
				Id:                   "c-bZ_nLfZufG",
				InvoiceNo:            "c-eJ2e_cRCaM",
				Month:                "c-wR0IONcxGH",
				Employee:             "c-bbHUhqlbfN",
				PreviousInvoice:      "c-FQ7rKmbXr6",
				AmountRub:            "c-ygtGw9Kilw",
				EurRubExpected:       "c-tvtGu9juVL",
				RequestedSubtotalEur: "c-9rnJJZ6gA7",
				RoundingPrevMonEur:   "c-hLrmDsk89g",
				Rounding:             "c-Tri-EGUP_n",
				AmountRequestedEur:   "c-bJpHVxywXD",
				Hours:                "c-KtVV9if8P7",
				EurRubActual:         "c-kLIyv9EvyH",
				AmountRubActual:      "c-AxLSgrt7e3",
				RateErrorRub:         "c-SsHRhKa_uC",
				CostOfDay:            "c-yJnq9stsgi",
				BankTotalFees:        "c-IYFe5lD4td",
				OpeningDateIp:        "c-hI1iZG3xzY",
				UnpaidDay:            "c-23E3m2Yk_O",
				CorrectionRefs:       "c-tpeCMU21_I",
				CorrectionRub:        "c-jNcl4nZe_h",
				PatentRub:            "c-qA_pPM9kuZ",
				TaxesRub:             "c-ug709va8_K",
				PatentRefs:           "c-bkVlencAZt",
				TaxesRefs:            "c-aYkpi97eXt",
				UnpaidDayRefs:        "c-IvUKXU4063",
				BaseSalary:           "c-wqNhZf9EQY",
				PaymentService:       "c-sRGR6jYC7g",
			},
		},
		Corrections: CorrectionsDecl{
			IdOnly{"grid-wBmvgFgaGi"},
			CorrectionsColumns{
				Comment:                  "c--_r48PQnSn",
				AbsoluteCorrectionEur:    "c-pRmEece9pf",
				AbsoluteCorrectionRub:    "c-P2x5IJuMXN",
				EurRubExpected:           "c-8LD34cnmCh",
				AbsCorrectionEurInRub:    "c-_GvG9w3Qs7",
				PerDayType:               "c-3Ivn-M1j7-",
				NumberOfDays:             "c-gDOyigH1cm",
				CostOfDay:                "c-K_Iy0iERKR",
				PerDay:                   "c-Y2E1Vwe2_-",
				TotalCorrectionRub:       "c-0arkfr4qXv",
				PaymentInvoice:           "c-7SU0iOBY9J",
				PerDayCoefficient:        "c-pz6W2IRzFR",
				PerDayCalculationInvoice: "c-bK4qXZUCqs",
				Display:                  "c-FVW_9PPzZ2",
				Category:                 "c-zDY58PF0P6",
			},
		},
		Taxes: TaxesDecl{
			IdOnly{"grid-057HtXfvYH"},
			TaxesColumns{
				Invoice:            "c-sYkd0N-9Ef",
				OpeningDateIp:      "c-Q83cwaB-vf",
				PeriodStart:        "c-UgsMCGn1oX",
				PeriodEnd:          "c-7Bt0dg_odm",
				AmountIpDays:       "c-wraB-EUIPu",
				MedicalFund:        "c-hRpESBxMnx",
				PensionFundFixed:   "c-5Gm9BIf7sa",
				PensionFundPercent: "c-nO-EnVUZSb",
				Amount:             "c-FlQQoKxoau",
			},
		},
		Patent: PatentDecl{
			IdOnly{"grid-_IJllxLQCt"},
			PatentColumns{
				Invoice:       "c-sYkd0N-9Ef",
				OpeningPatent: "c-GfXRAu21-_",
				PeriodEnd:     "c-7Bt0dg_odm",
				FullMonths:    "c-gARgn0s4h-",
				AnnualCost:    "c-ayAtt8TYPI",
				PeriodCost:    "c-FlQQoKxoau",
				Period:        "c-gtV-Qz9osQ",
			},
		},
		Months: MonthsDecl{
			IdOnly: IdOnly{"grid-laH8qsdDyP"},
			Cols: MonthsColumns{
				Month:             "c-u_Pgrgevw7",
				WorkDays:          "c-2eD8ott5yA",
				Year:              "c-NgPwXZshJM",
				ID:                "c-iRCjZ0JBcM",
				PreviousMonthLink: "c-3Cc_lYdvmW",
				PreviousMonth:     "c-vuW159vf-o",
				IsCurrent:         "c-OLxcAJLQoW",
			},
		},
		CodaFormulas: CodaFormulasDecl{
			CurrentMonth: "f-rnJn4-MytN",
		},
		//CodaViews: CodaViewsDecl{
		//	ThisMonthInvoices:    "table-_k4m2krdyp",
		//	ThisMonthCorrections: "table-29Y7W0xzVA",
		//	ThisMonthTaxes:       "table-X3zgg4jJ8-",
		//	ThisMonthPatents:     "table-n2aIqJxO80",
		//},
	}
}
