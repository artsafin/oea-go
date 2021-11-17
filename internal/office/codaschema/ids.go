package codaschema

// Basic types

type _schema struct {
	Table   _tableSchema
	Formula _formulaSchema
	Control _controlSchema
}

type _codaEntity struct {
	ID   string
	Name string
}

func (e _codaEntity) String() string {
	return e.ID
}

type _tableSchema struct {
	Months                         _monthsTable                         // Months
	CashFlow                       _cashFlowTable                       // Cash flow
	InvoicePayment                 _invoicePaymentTable                 // Invoice payment
	Expenses                       _expensesTable                       // Expenses
	Invoices                       _invoicesTable                       // Invoices
	Accounts                       _accountsTable                       // Accounts
	UnplannedLast40d               _unplannedLast40dTable               // Unplanned (last 40d)
	PlannedSpendsAndCurrentInvoice _plannedSpendsAndCurrentInvoiceTable // Planned spends and current invoice
	InventoryTypes                 _inventoryTypesTable                 // Inventory Types
	InvoiceTemplate                _invoiceTemplateTable                // Invoice template
	PlannedAndPreviousInvoices     _plannedAndPreviousInvoicesTable     // Planned and previous invoices
	PlannedExpenses                _plannedExpensesTable                // Planned expenses
	CashFlow2020                   _cashFlow2020Table                   // Cash flow 2020
	CashFlowOfPersonalAcc          _cashFlowOfPersonalAccTable          // Cash flow of personal acc
	CashFlow2021                   _cashFlow2021Table                   // Cash flow 2021
	AuditCategory                  _auditCategoryTable                  // Audit Category
	ExpensesByCategory             _expensesByCategoryTable             // Expenses by category
}
type _formulaSchema struct {
	PlannedInvoice          _codaEntity // plannedInvoice
	LastInvoice             _codaEntity // lastInvoice
	EvaluatedCompanyBalance _codaEntity // evaluatedCompanyBalance
	EvaluatedTotalPending   _codaEntity // evaluatedTotalPending
	TplMonth                _codaEntity // tplMonth
	TplNextMonth            _codaEntity // tplNextMonth
}
type _controlSchema struct {
	BtnAddUnplanned _codaEntity // btnAddUnplanned
	BtnInvoiceAdd   _codaEntity // btnInvoiceAdd
	Button1         _codaEntity // button 1
}

// Table Months
type _monthsTable struct {
	_codaEntity
	Cols _monthsTableColumns
}
type _monthsTableColumns struct {
	Row               _codaEntity // Row
	Synced            _codaEntity // Synced
	SyncAccount       _codaEntity // Sync account
	Name              _codaEntity // Name
	Month             _codaEntity // Month
	PreviousMonth     _codaEntity // Previous Month
	MonthName         _codaEntity // Month Name
	WorkDays          _codaEntity // Work Days
	Year              _codaEntity // Year
	PreviousMonthLink _codaEntity // Previous Month Link
}

// Table Cash flow
type _cashFlowTable struct {
	_codaEntity
	Cols _cashFlowTableColumns
}
type _cashFlowTableColumns struct {
	No              _codaEntity // No
	Date            _codaEntity // Date
	Comment         _codaEntity // Comment
	AmountRUB       _codaEntity // Amount, RUB
	Account         _codaEntity // Account
	CashOutPurpose  _codaEntity // Cash out purpose
	CashIn          _codaEntity // Cash in
	ComissionAmtRUB _codaEntity // Comission amt, RUB
	PaidAmtRUB      _codaEntity // Paid amt, RUB
	Reconciled      _codaEntity // Reconciled
	Count           _codaEntity // Count
	InclInBalance   _codaEntity // Incl in balance
	Author          _codaEntity // Author
	PersonalPaidIn  _codaEntity // Personal paid in
	CreatedOn       _codaEntity // Created on
}

// Table Invoice payment
type _invoicePaymentTable struct {
	_codaEntity
	Cols _invoicePaymentTableColumns
}
type _invoicePaymentTableColumns struct {
	ReceivedDate _codaEntity // Received date
	AmountEUR    _codaEntity // Amount EUR
	PaymentRate  _codaEntity // Payment rate
	PlannedRate  _codaEntity // Planned rate
	RateDiff     _codaEntity // Rate diff
	AmountRUB    _codaEntity // Amount RUB
	RateBalance  _codaEntity // Rate balance
	SentDate     _codaEntity // Sent date
	RateError    _codaEntity // Rate error
	Invoice      _codaEntity // Invoice
	CashFlow     _codaEntity // Cash Flow
}

// Table Expenses
type _expensesTable struct {
	_codaEntity
	Cols _expensesTableColumns
}
type _expensesTableColumns struct {
	ID              _codaEntity // ID
	Invoice         _codaEntity // Invoice
	Subject         _codaEntity // Subject
	Category        _codaEntity // Category
	Comment         _codaEntity // Comment
	AmountRUB       _codaEntity // Amount, RUB
	AmountEUR       _codaEntity // Amount, EUR
	Status          _codaEntity // Status
	CashOutRefs     _codaEntity // Cash out refs
	ActuallySpent   _codaEntity // Actually spent
	RejectionReason _codaEntity // Rejection reason
	PendingSpend    _codaEntity // Pending spend
	Balance         _codaEntity // Balance
	CashFlow        _codaEntity // Cash flow
	LastCashOutDate _codaEntity // Last cash out date
	Sort            _codaEntity // Sort
	Group           _codaEntity // Group
	InventoryRef    _codaEntity // Inventory Ref
	ModifiedBy      _codaEntity // Modified by
	CreatedBy       _codaEntity // Created by
	CreatedOn       _codaEntity // Created on
	ModifiedOn      _codaEntity // Modified on
	AuditCategory   _codaEntity // Audit Category
}

// Table Invoices
type _invoicesTable struct {
	_codaEntity
	Cols _invoicesTableColumns
}
type _invoicesTableColumns struct {
	No                _codaEntity // No
	Status            _codaEntity // Status
	Number            _codaEntity // Number
	Date              _codaEntity // Date
	Filename          _codaEntity // Filename
	HourRate          _codaEntity // Hour rate
	ApprovalLink      _codaEntity // Approval Link
	EURFixedRate      _codaEntity // EUR fixed rate
	EURRateWorst      _codaEntity // EUR rate worst
	ExpensesRUB       _codaEntity // Expenses, RUB
	ExpensesEUR       _codaEntity // Expenses, EUR
	ReturnOfRounding  _codaEntity // Return of rounding
	Subtotal          _codaEntity // Subtotal
	HourRateRounding  _codaEntity // Hour Rate Rounding
	TotalEUR          _codaEntity // Total, EUR
	Hours             _codaEntity // Hours
	InvoicePayment    _codaEntity // Invoice payment
	PSCCFee           _codaEntity // PS: CC fee
	PSRateRisk        _codaEntity // PS: Rate risk
	PrevInvoice       _codaEntity // Prev invoice
	PlannedExpenses   _codaEntity // Planned expenses
	InsertTemplate    _codaEntity // Insert template
	ActuallySpent     _codaEntity // Actually spent
	PendingSpend      _codaEntity // Pending spend
	Balance           _codaEntity // Balance
	InvoicePaymentAdd _codaEntity // Invoice payment add
	IsRecent          _codaEntity // Is recent
}

// Table Accounts
type _accountsTable struct {
	_codaEntity
	Cols _accountsTableColumns
}
type _accountsTableColumns struct {
	Account       _codaEntity // Account
	Comment       _codaEntity // Comment
	InclInBalance _codaEntity // Incl in balance
}

// Table Unplanned (last 40d)
type _unplannedLast40dTable struct {
	_codaEntity
	Cols _unplannedLast40dTableColumns
}
type _unplannedLast40dTableColumns struct {
	Subject         _codaEntity // Subject
	LastCashOutDate _codaEntity // Last cash out date
	Comment         _codaEntity // Comment
	Category        _codaEntity // Category
	Status          _codaEntity // Status
	ActuallySpent   _codaEntity // Actually spent
	Balance         _codaEntity // Balance
	CashFlow        _codaEntity // Cash flow
	CashOutRefs     _codaEntity // Cash out refs
}

// Table Planned spends and current invoice
type _plannedSpendsAndCurrentInvoiceTable struct {
	_codaEntity
	Cols _plannedSpendsAndCurrentInvoiceTableColumns
}
type _plannedSpendsAndCurrentInvoiceTableColumns struct {
	Subject       _codaEntity // Subject
	Status        _codaEntity // Status
	AmountRUB     _codaEntity // Amount, RUB
	ActuallySpent _codaEntity // Actually spent
	PendingSpend  _codaEntity // Pending spend
	Balance       _codaEntity // Balance
	CashFlow      _codaEntity // Cash flow
	Group         _codaEntity // Group
}

// Table Inventory Types
type _inventoryTypesTable struct {
	_codaEntity
	Cols _inventoryTypesTableColumns
}
type _inventoryTypesTableColumns struct {
	Row         _codaEntity // Row
	Synced      _codaEntity // Synced
	SyncAccount _codaEntity // Sync account
	Code        _codaEntity // Code
	Type        _codaEntity // Type
	Name        _codaEntity // Name
	Link        _codaEntity // Link
	Comments    _codaEntity // Comments
	Price       _codaEntity // Price
}

// Table Invoice template
type _invoiceTemplateTable struct {
	_codaEntity
	Cols _invoiceTemplateTableColumns
}
type _invoiceTemplateTableColumns struct {
	Subject    _codaEntity // Subject
	Comment    _codaEntity // Comment
	AmountRUB  _codaEntity // Amount, RUB
	Category   _codaEntity // Category
	Expense    _codaEntity // Expense
	SubjectTpl _codaEntity // Subject Tpl
	Active     _codaEntity // Active
}

// Table Planned and previous invoices
type _plannedAndPreviousInvoicesTable struct {
	_codaEntity
	Cols _plannedAndPreviousInvoicesTableColumns
}
type _plannedAndPreviousInvoicesTableColumns struct {
	No                _codaEntity // No
	InsertTemplate    _codaEntity // Insert template
	Status            _codaEntity // Status
	Number            _codaEntity // Number
	Date              _codaEntity // Date
	HourRate          _codaEntity // Hour rate
	ApprovalLink      _codaEntity // Approval Link
	EURFixedRate      _codaEntity // EUR fixed rate
	EURRateWorst      _codaEntity // EUR rate worst
	PlannedExpenses   _codaEntity // Planned expenses
	ExpensesRUB       _codaEntity // Expenses, RUB
	ExpensesEUR       _codaEntity // Expenses, EUR
	ReturnOfRounding  _codaEntity // Return of rounding
	Subtotal          _codaEntity // Subtotal
	HourRateRounding  _codaEntity // Hour Rate Rounding
	TotalEUR          _codaEntity // Total, EUR
	Hours             _codaEntity // Hours
	InvoicePayment    _codaEntity // Invoice payment
	InvoicePaymentAdd _codaEntity // Invoice payment add
	PSRateRisk        _codaEntity // PS: Rate risk
	PSCCFee           _codaEntity // PS: CC fee
	ActuallySpent     _codaEntity // Actually spent
	PendingSpend      _codaEntity // Pending spend
	Balance           _codaEntity // Balance
	PrevInvoice       _codaEntity // Prev invoice
}

// Table Planned expenses
type _plannedExpensesTable struct {
	_codaEntity
	Cols _plannedExpensesTableColumns
}
type _plannedExpensesTableColumns struct {
	Invoice       _codaEntity // Invoice
	Subject       _codaEntity // Subject
	Comment       _codaEntity // Comment
	AuditCategory _codaEntity // Audit Category
	Category      _codaEntity // Category
	Status        _codaEntity // Status
	AmountRUB     _codaEntity // Amount, RUB
	AmountEUR     _codaEntity // Amount, EUR
	ActuallySpent _codaEntity // Actually spent
	PendingSpend  _codaEntity // Pending spend
	Balance       _codaEntity // Balance
	CashFlow      _codaEntity // Cash flow
}

// Table Cash flow 2020
type _cashFlow2020Table struct {
	_codaEntity
	Cols _cashFlow2020TableColumns
}
type _cashFlow2020TableColumns struct {
	Author          _codaEntity // Author
	Reconciled      _codaEntity // Reconciled
	Account         _codaEntity // Account
	Date            _codaEntity // Date
	AmountRUB       _codaEntity // Amount, RUB
	PaidAmtRUB      _codaEntity // Paid amt, RUB
	CashOutPurpose  _codaEntity // Cash out purpose
	Count           _codaEntity // Count
	ComissionAmtRUB _codaEntity // Comission amt, RUB
	Comment         _codaEntity // Comment
	CashIn          _codaEntity // Cash in
}

// Table Cash flow of personal acc
type _cashFlowOfPersonalAccTable struct {
	_codaEntity
	Cols _cashFlowOfPersonalAccTableColumns
}
type _cashFlowOfPersonalAccTableColumns struct {
	Account        _codaEntity // Account
	Date           _codaEntity // Date
	AmountRUB      _codaEntity // Amount, RUB
	Comment        _codaEntity // Comment
	CashOutPurpose _codaEntity // Cash out purpose
	PersonalPaidIn _codaEntity // Personal paid in
}

// Table Cash flow 2021
type _cashFlow2021Table struct {
	_codaEntity
	Cols _cashFlow2021TableColumns
}
type _cashFlow2021TableColumns struct {
	Author          _codaEntity // Author
	Reconciled      _codaEntity // Reconciled
	Account         _codaEntity // Account
	Date            _codaEntity // Date
	AmountRUB       _codaEntity // Amount, RUB
	PaidAmtRUB      _codaEntity // Paid amt, RUB
	CashOutPurpose  _codaEntity // Cash out purpose
	Count           _codaEntity // Count
	ComissionAmtRUB _codaEntity // Comission amt, RUB
	Comment         _codaEntity // Comment
	CashIn          _codaEntity // Cash in
}

// Table Audit Category
type _auditCategoryTable struct {
	_codaEntity
	Cols _auditCategoryTableColumns
}
type _auditCategoryTableColumns struct {
	AuditCategory _codaEntity // Audit Category
}

// Table Expenses by category
type _expensesByCategoryTable struct {
	_codaEntity
	Cols _expensesByCategoryTableColumns
}
type _expensesByCategoryTableColumns struct {
	Group         _codaEntity // Group
	AmountRUB     _codaEntity // Amount, RUB
	AmountEUR     _codaEntity // Amount, EUR
	AuditCategory _codaEntity // Audit Category
}

var ID _schema

func init() {
	ID = _schema{
		Table: _tableSchema{
			Months: _monthsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-sync-1054-Table-dynamic-aca1dcec432de5424fa83d4c6d93c83946906407bcf71db42c915765fbc14401",
					Name: "Months",
				},
				Cols: _monthsTableColumns{
					Row: _codaEntity{
						ID:   "value",
						Name: "Row",
					},
					Synced: _codaEntity{
						ID:   "synced",
						Name: "Synced",
					},
					SyncAccount: _codaEntity{
						ID:   "connection",
						Name: "Sync account",
					},
					Name: _codaEntity{
						ID:   "c-XqOHdzXcLR",
						Name: "Name",
					},
					Month: _codaEntity{
						ID:   "c-fEt-r8-748",
						Name: "Month",
					},
					PreviousMonth: _codaEntity{
						ID:   "c-RUsq1_2_81",
						Name: "Previous Month",
					},
					MonthName: _codaEntity{
						ID:   "c-UHSQHFZLkE",
						Name: "Month Name",
					},
					WorkDays: _codaEntity{
						ID:   "c-RlcX07v2ny",
						Name: "Work Days",
					},
					Year: _codaEntity{
						ID:   "c-mnHhV-4LX7",
						Name: "Year",
					},
					PreviousMonthLink: _codaEntity{
						ID:   "c-hikG-Ajf5k",
						Name: "Previous Month Link",
					},
				},
			},
			CashFlow: _cashFlowTable{
				_codaEntity: _codaEntity{
					ID:   "grid-7yE1NK5upn",
					Name: "Cash flow",
				},
				Cols: _cashFlowTableColumns{
					No: _codaEntity{
						ID:   "c-9TdkReuS8G",
						Name: "No",
					},
					Date: _codaEntity{
						ID:   "c-HuVxndys-5",
						Name: "Date",
					},
					Comment: _codaEntity{
						ID:   "c-7Sr6bn-OVs",
						Name: "Comment",
					},
					AmountRUB: _codaEntity{
						ID:   "c-Ka5rkoQEvD",
						Name: "Amount, RUB",
					},
					Account: _codaEntity{
						ID:   "c-HvpjCI4gRX",
						Name: "Account",
					},
					CashOutPurpose: _codaEntity{
						ID:   "c-ak93zMH1dQ",
						Name: "Cash out purpose",
					},
					CashIn: _codaEntity{
						ID:   "c-FfATqG9jTV",
						Name: "Cash in",
					},
					ComissionAmtRUB: _codaEntity{
						ID:   "c-6pxTqY1M65",
						Name: "Comission amt, RUB",
					},
					PaidAmtRUB: _codaEntity{
						ID:   "c-NcMKzn_fpI",
						Name: "Paid amt, RUB",
					},
					Reconciled: _codaEntity{
						ID:   "c-P_CSVSFk0L",
						Name: "Reconciled",
					},
					Count: _codaEntity{
						ID:   "c-taBOV-t5cz",
						Name: "Count",
					},
					InclInBalance: _codaEntity{
						ID:   "c-PV5Apg8Ajw",
						Name: "Incl in balance",
					},
					Author: _codaEntity{
						ID:   "c-2Xkn5hlxUM",
						Name: "Author",
					},
					PersonalPaidIn: _codaEntity{
						ID:   "c-OPYJnVV_u7",
						Name: "Personal paid in",
					},
					CreatedOn: _codaEntity{
						ID:   "c-F-Ffna2NpT",
						Name: "Created on",
					},
				},
			},
			InvoicePayment: _invoicePaymentTable{
				_codaEntity: _codaEntity{
					ID:   "grid-R9oRxgWMUH",
					Name: "Invoice payment",
				},
				Cols: _invoicePaymentTableColumns{
					ReceivedDate: _codaEntity{
						ID:   "c-HevgF6a6li",
						Name: "Received date",
					},
					AmountEUR: _codaEntity{
						ID:   "c-IE-iOWanac",
						Name: "Amount EUR",
					},
					PaymentRate: _codaEntity{
						ID:   "c-zKkSve_s9b",
						Name: "Payment rate",
					},
					PlannedRate: _codaEntity{
						ID:   "c-VXDlksbd7o",
						Name: "Planned rate",
					},
					RateDiff: _codaEntity{
						ID:   "c-qH9PR8hAuY",
						Name: "Rate diff",
					},
					AmountRUB: _codaEntity{
						ID:   "c-Wx9eWgoHWl",
						Name: "Amount RUB",
					},
					RateBalance: _codaEntity{
						ID:   "c-9FRsy0d1l9",
						Name: "Rate balance",
					},
					SentDate: _codaEntity{
						ID:   "c-M1NogFDhKz",
						Name: "Sent date",
					},
					RateError: _codaEntity{
						ID:   "c-tSU9bZrnUG",
						Name: "Rate error",
					},
					Invoice: _codaEntity{
						ID:   "c-GD0FEsDsIu",
						Name: "Invoice",
					},
					CashFlow: _codaEntity{
						ID:   "c-doy2oZs_0X",
						Name: "Cash Flow",
					},
				},
			},
			Expenses: _expensesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-j1yvp-c7Xq",
					Name: "Expenses",
				},
				Cols: _expensesTableColumns{
					ID: _codaEntity{
						ID:   "c-NIDES-BkAW",
						Name: "ID",
					},
					Invoice: _codaEntity{
						ID:   "c-_Vwi1N_WO9",
						Name: "Invoice",
					},
					Subject: _codaEntity{
						ID:   "c-dWrUpt7GHg",
						Name: "Subject",
					},
					Category: _codaEntity{
						ID:   "c-HgZFvbcM6u",
						Name: "Category",
					},
					Comment: _codaEntity{
						ID:   "c-hLtZdaFaOE",
						Name: "Comment",
					},
					AmountRUB: _codaEntity{
						ID:   "c-NWO2JHrrdw",
						Name: "Amount, RUB",
					},
					AmountEUR: _codaEntity{
						ID:   "c-85jKBvNSVF",
						Name: "Amount, EUR",
					},
					Status: _codaEntity{
						ID:   "c-TlBKoSAJmV",
						Name: "Status",
					},
					CashOutRefs: _codaEntity{
						ID:   "c-o0llMz254Z",
						Name: "Cash out refs",
					},
					ActuallySpent: _codaEntity{
						ID:   "c-otkrTpI7VA",
						Name: "Actually spent",
					},
					RejectionReason: _codaEntity{
						ID:   "c-OwY64F1V4B",
						Name: "Rejection reason",
					},
					PendingSpend: _codaEntity{
						ID:   "c-MTRmKDu-3A",
						Name: "Pending spend",
					},
					Balance: _codaEntity{
						ID:   "c-GAIziVHBvd",
						Name: "Balance",
					},
					CashFlow: _codaEntity{
						ID:   "c-sncvkVSbcJ",
						Name: "Cash flow",
					},
					LastCashOutDate: _codaEntity{
						ID:   "c-BhuJTIsQEP",
						Name: "Last cash out date",
					},
					Sort: _codaEntity{
						ID:   "c-lzNMnSE_dh",
						Name: "Sort",
					},
					Group: _codaEntity{
						ID:   "c-18G6IQXBvd",
						Name: "Group",
					},
					InventoryRef: _codaEntity{
						ID:   "c-4-a6cFeLHl",
						Name: "Inventory Ref",
					},
					ModifiedBy: _codaEntity{
						ID:   "c-BFcg4AF5oM",
						Name: "Modified by",
					},
					CreatedBy: _codaEntity{
						ID:   "c-ZhqiSyi4mD",
						Name: "Created by",
					},
					CreatedOn: _codaEntity{
						ID:   "c-9PVZmwivnx",
						Name: "Created on",
					},
					ModifiedOn: _codaEntity{
						ID:   "c-8TigMDiu_5",
						Name: "Modified on",
					},
					AuditCategory: _codaEntity{
						ID:   "c-45roqK-a18",
						Name: "Audit Category",
					},
				},
			},
			Invoices: _invoicesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-H_lQoXT4Hn",
					Name: "Invoices",
				},
				Cols: _invoicesTableColumns{
					No: _codaEntity{
						ID:   "c-1FLrjxKHe5",
						Name: "No",
					},
					Status: _codaEntity{
						ID:   "c-rqUgKu_1z_",
						Name: "Status",
					},
					Number: _codaEntity{
						ID:   "c-qw6CtPRr6R",
						Name: "Number",
					},
					Date: _codaEntity{
						ID:   "c-r6ev-Hfy0R",
						Name: "Date",
					},
					Filename: _codaEntity{
						ID:   "c-gVFw0tg27k",
						Name: "Filename",
					},
					HourRate: _codaEntity{
						ID:   "c-wu0qKZLU0Z",
						Name: "Hour rate",
					},
					ApprovalLink: _codaEntity{
						ID:   "c-mINTWSuok_",
						Name: "Approval Link",
					},
					EURFixedRate: _codaEntity{
						ID:   "c-39yF51XkOy",
						Name: "EUR fixed rate",
					},
					EURRateWorst: _codaEntity{
						ID:   "c-eqfmnjpn6E",
						Name: "EUR rate worst",
					},
					ExpensesRUB: _codaEntity{
						ID:   "c-LnWC3NrrYo",
						Name: "Expenses, RUB",
					},
					ExpensesEUR: _codaEntity{
						ID:   "c-Bj-zR8QBZP",
						Name: "Expenses, EUR",
					},
					ReturnOfRounding: _codaEntity{
						ID:   "c-cEdaY7AsWT",
						Name: "Return of rounding",
					},
					Subtotal: _codaEntity{
						ID:   "c-EfZ0XNkbSp",
						Name: "Subtotal",
					},
					HourRateRounding: _codaEntity{
						ID:   "c-wbNLVBU5oW",
						Name: "Hour Rate Rounding",
					},
					TotalEUR: _codaEntity{
						ID:   "c-gywxqGt4uK",
						Name: "Total, EUR",
					},
					Hours: _codaEntity{
						ID:   "c-2kpnqmokJ5",
						Name: "Hours",
					},
					InvoicePayment: _codaEntity{
						ID:   "c-_D1oWYheaQ",
						Name: "Invoice payment",
					},
					PSCCFee: _codaEntity{
						ID:   "c-isMKk3UN7G",
						Name: "PS: CC fee",
					},
					PSRateRisk: _codaEntity{
						ID:   "c-vrecptlM-5",
						Name: "PS: Rate risk",
					},
					PrevInvoice: _codaEntity{
						ID:   "c-RWqrc6Q-Zf",
						Name: "Prev invoice",
					},
					PlannedExpenses: _codaEntity{
						ID:   "c-cgiJ5IeX4n",
						Name: "Planned expenses",
					},
					InsertTemplate: _codaEntity{
						ID:   "c-sTMJM9Pq-d",
						Name: "Insert template",
					},
					ActuallySpent: _codaEntity{
						ID:   "c-XY-LhijBo3",
						Name: "Actually spent",
					},
					PendingSpend: _codaEntity{
						ID:   "c-bccwx_DDwv",
						Name: "Pending spend",
					},
					Balance: _codaEntity{
						ID:   "c-MW_JQd5O0b",
						Name: "Balance",
					},
					InvoicePaymentAdd: _codaEntity{
						ID:   "c-MQfrZ0gvsC",
						Name: "Invoice payment add",
					},
					IsRecent: _codaEntity{
						ID:   "c-yBkBMvhtOT",
						Name: "Is recent",
					},
				},
			},
			Accounts: _accountsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-pwTQC7XRKL",
					Name: "Accounts",
				},
				Cols: _accountsTableColumns{
					Account: _codaEntity{
						ID:   "c-FiYh8LlPM6",
						Name: "Account",
					},
					Comment: _codaEntity{
						ID:   "c-tokAPP6nTJ",
						Name: "Comment",
					},
					InclInBalance: _codaEntity{
						ID:   "c-G8_c9D7Jsy",
						Name: "Incl in balance",
					},
				},
			},
			UnplannedLast40d: _unplannedLast40dTable{
				_codaEntity: _codaEntity{
					ID:   "table-2iAFNgZh_z",
					Name: "Unplanned (last 40d)",
				},
				Cols: _unplannedLast40dTableColumns{
					Subject: _codaEntity{
						ID:   "c-dWrUpt7GHg",
						Name: "Subject",
					},
					LastCashOutDate: _codaEntity{
						ID:   "c-BhuJTIsQEP",
						Name: "Last cash out date",
					},
					Comment: _codaEntity{
						ID:   "c-hLtZdaFaOE",
						Name: "Comment",
					},
					Category: _codaEntity{
						ID:   "c-HgZFvbcM6u",
						Name: "Category",
					},
					Status: _codaEntity{
						ID:   "c-TlBKoSAJmV",
						Name: "Status",
					},
					ActuallySpent: _codaEntity{
						ID:   "c-otkrTpI7VA",
						Name: "Actually spent",
					},
					Balance: _codaEntity{
						ID:   "c-GAIziVHBvd",
						Name: "Balance",
					},
					CashFlow: _codaEntity{
						ID:   "c-sncvkVSbcJ",
						Name: "Cash flow",
					},
					CashOutRefs: _codaEntity{
						ID:   "c-o0llMz254Z",
						Name: "Cash out refs",
					},
				},
			},
			PlannedSpendsAndCurrentInvoice: _plannedSpendsAndCurrentInvoiceTable{
				_codaEntity: _codaEntity{
					ID:   "table-5bvJSmcSgd",
					Name: "Planned spends and current invoice",
				},
				Cols: _plannedSpendsAndCurrentInvoiceTableColumns{
					Subject: _codaEntity{
						ID:   "c-dWrUpt7GHg",
						Name: "Subject",
					},
					Status: _codaEntity{
						ID:   "c-TlBKoSAJmV",
						Name: "Status",
					},
					AmountRUB: _codaEntity{
						ID:   "c-NWO2JHrrdw",
						Name: "Amount, RUB",
					},
					ActuallySpent: _codaEntity{
						ID:   "c-otkrTpI7VA",
						Name: "Actually spent",
					},
					PendingSpend: _codaEntity{
						ID:   "c-MTRmKDu-3A",
						Name: "Pending spend",
					},
					Balance: _codaEntity{
						ID:   "c-GAIziVHBvd",
						Name: "Balance",
					},
					CashFlow: _codaEntity{
						ID:   "c-sncvkVSbcJ",
						Name: "Cash flow",
					},
					Group: _codaEntity{
						ID:   "c-18G6IQXBvd",
						Name: "Group",
					},
				},
			},
			InventoryTypes: _inventoryTypesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-sync-1054-Table-dynamic-ef49a72921eadd1e229c45fac0199e3555d7772b87293e602f033a92d9caea1d",
					Name: "Inventory Types",
				},
				Cols: _inventoryTypesTableColumns{
					Row: _codaEntity{
						ID:   "value",
						Name: "Row",
					},
					Synced: _codaEntity{
						ID:   "synced",
						Name: "Synced",
					},
					SyncAccount: _codaEntity{
						ID:   "connection",
						Name: "Sync account",
					},
					Code: _codaEntity{
						ID:   "c-IedG4qMAaB",
						Name: "Code",
					},
					Type: _codaEntity{
						ID:   "c-fPa38HqyKw",
						Name: "Type",
					},
					Name: _codaEntity{
						ID:   "c-qVRPgjXTwg",
						Name: "Name",
					},
					Link: _codaEntity{
						ID:   "c-O-47pRG6yt",
						Name: "Link",
					},
					Comments: _codaEntity{
						ID:   "c-vTVmek8RJw",
						Name: "Comments",
					},
					Price: _codaEntity{
						ID:   "c-xzD6XFmymP",
						Name: "Price",
					},
				},
			},
			InvoiceTemplate: _invoiceTemplateTable{
				_codaEntity: _codaEntity{
					ID:   "grid-zsR0KGVVNX",
					Name: "Invoice template",
				},
				Cols: _invoiceTemplateTableColumns{
					Subject: _codaEntity{
						ID:   "c-z9FQ_YZ4WD",
						Name: "Subject",
					},
					Comment: _codaEntity{
						ID:   "c-1CAyQoTxTK",
						Name: "Comment",
					},
					AmountRUB: _codaEntity{
						ID:   "c-xsUIIY0puP",
						Name: "Amount, RUB",
					},
					Category: _codaEntity{
						ID:   "c-jGM_Cns-Wa",
						Name: "Category",
					},
					Expense: _codaEntity{
						ID:   "c-s0XIiK7UdC",
						Name: "Expense",
					},
					SubjectTpl: _codaEntity{
						ID:   "c-a-y2sDK_aA",
						Name: "Subject Tpl",
					},
					Active: _codaEntity{
						ID:   "c--5hsp0sV5S",
						Name: "Active",
					},
				},
			},
			PlannedAndPreviousInvoices: _plannedAndPreviousInvoicesTable{
				_codaEntity: _codaEntity{
					ID:   "table-0oiNc6Kkw8",
					Name: "Planned and previous invoices",
				},
				Cols: _plannedAndPreviousInvoicesTableColumns{
					No: _codaEntity{
						ID:   "c-1FLrjxKHe5",
						Name: "No",
					},
					InsertTemplate: _codaEntity{
						ID:   "c-sTMJM9Pq-d",
						Name: "Insert template",
					},
					Status: _codaEntity{
						ID:   "c-rqUgKu_1z_",
						Name: "Status",
					},
					Number: _codaEntity{
						ID:   "c-qw6CtPRr6R",
						Name: "Number",
					},
					Date: _codaEntity{
						ID:   "c-r6ev-Hfy0R",
						Name: "Date",
					},
					HourRate: _codaEntity{
						ID:   "c-wu0qKZLU0Z",
						Name: "Hour rate",
					},
					ApprovalLink: _codaEntity{
						ID:   "c-mINTWSuok_",
						Name: "Approval Link",
					},
					EURFixedRate: _codaEntity{
						ID:   "c-39yF51XkOy",
						Name: "EUR fixed rate",
					},
					EURRateWorst: _codaEntity{
						ID:   "c-eqfmnjpn6E",
						Name: "EUR rate worst",
					},
					PlannedExpenses: _codaEntity{
						ID:   "c-cgiJ5IeX4n",
						Name: "Planned expenses",
					},
					ExpensesRUB: _codaEntity{
						ID:   "c-LnWC3NrrYo",
						Name: "Expenses, RUB",
					},
					ExpensesEUR: _codaEntity{
						ID:   "c-Bj-zR8QBZP",
						Name: "Expenses, EUR",
					},
					ReturnOfRounding: _codaEntity{
						ID:   "c-cEdaY7AsWT",
						Name: "Return of rounding",
					},
					Subtotal: _codaEntity{
						ID:   "c-EfZ0XNkbSp",
						Name: "Subtotal",
					},
					HourRateRounding: _codaEntity{
						ID:   "c-wbNLVBU5oW",
						Name: "Hour Rate Rounding",
					},
					TotalEUR: _codaEntity{
						ID:   "c-gywxqGt4uK",
						Name: "Total, EUR",
					},
					Hours: _codaEntity{
						ID:   "c-2kpnqmokJ5",
						Name: "Hours",
					},
					InvoicePayment: _codaEntity{
						ID:   "c-_D1oWYheaQ",
						Name: "Invoice payment",
					},
					InvoicePaymentAdd: _codaEntity{
						ID:   "c-MQfrZ0gvsC",
						Name: "Invoice payment add",
					},
					PSRateRisk: _codaEntity{
						ID:   "c-vrecptlM-5",
						Name: "PS: Rate risk",
					},
					PSCCFee: _codaEntity{
						ID:   "c-isMKk3UN7G",
						Name: "PS: CC fee",
					},
					ActuallySpent: _codaEntity{
						ID:   "c-XY-LhijBo3",
						Name: "Actually spent",
					},
					PendingSpend: _codaEntity{
						ID:   "c-bccwx_DDwv",
						Name: "Pending spend",
					},
					Balance: _codaEntity{
						ID:   "c-MW_JQd5O0b",
						Name: "Balance",
					},
					PrevInvoice: _codaEntity{
						ID:   "c-RWqrc6Q-Zf",
						Name: "Prev invoice",
					},
				},
			},
			PlannedExpenses: _plannedExpensesTable{
				_codaEntity: _codaEntity{
					ID:   "table-NvLNctjziJ",
					Name: "Planned expenses",
				},
				Cols: _plannedExpensesTableColumns{
					Invoice: _codaEntity{
						ID:   "c-_Vwi1N_WO9",
						Name: "Invoice",
					},
					Subject: _codaEntity{
						ID:   "c-dWrUpt7GHg",
						Name: "Subject",
					},
					Comment: _codaEntity{
						ID:   "c-hLtZdaFaOE",
						Name: "Comment",
					},
					AuditCategory: _codaEntity{
						ID:   "c-45roqK-a18",
						Name: "Audit Category",
					},
					Category: _codaEntity{
						ID:   "c-HgZFvbcM6u",
						Name: "Category",
					},
					Status: _codaEntity{
						ID:   "c-TlBKoSAJmV",
						Name: "Status",
					},
					AmountRUB: _codaEntity{
						ID:   "c-NWO2JHrrdw",
						Name: "Amount, RUB",
					},
					AmountEUR: _codaEntity{
						ID:   "c-85jKBvNSVF",
						Name: "Amount, EUR",
					},
					ActuallySpent: _codaEntity{
						ID:   "c-otkrTpI7VA",
						Name: "Actually spent",
					},
					PendingSpend: _codaEntity{
						ID:   "c-MTRmKDu-3A",
						Name: "Pending spend",
					},
					Balance: _codaEntity{
						ID:   "c-GAIziVHBvd",
						Name: "Balance",
					},
					CashFlow: _codaEntity{
						ID:   "c-sncvkVSbcJ",
						Name: "Cash flow",
					},
				},
			},
			CashFlow2020: _cashFlow2020Table{
				_codaEntity: _codaEntity{
					ID:   "table-0iqeriGwB9",
					Name: "Cash flow 2020",
				},
				Cols: _cashFlow2020TableColumns{
					Author: _codaEntity{
						ID:   "c-2Xkn5hlxUM",
						Name: "Author",
					},
					Reconciled: _codaEntity{
						ID:   "c-P_CSVSFk0L",
						Name: "Reconciled",
					},
					Account: _codaEntity{
						ID:   "c-HvpjCI4gRX",
						Name: "Account",
					},
					Date: _codaEntity{
						ID:   "c-HuVxndys-5",
						Name: "Date",
					},
					AmountRUB: _codaEntity{
						ID:   "c-Ka5rkoQEvD",
						Name: "Amount, RUB",
					},
					PaidAmtRUB: _codaEntity{
						ID:   "c-NcMKzn_fpI",
						Name: "Paid amt, RUB",
					},
					CashOutPurpose: _codaEntity{
						ID:   "c-ak93zMH1dQ",
						Name: "Cash out purpose",
					},
					Count: _codaEntity{
						ID:   "c-taBOV-t5cz",
						Name: "Count",
					},
					ComissionAmtRUB: _codaEntity{
						ID:   "c-6pxTqY1M65",
						Name: "Comission amt, RUB",
					},
					Comment: _codaEntity{
						ID:   "c-7Sr6bn-OVs",
						Name: "Comment",
					},
					CashIn: _codaEntity{
						ID:   "c-FfATqG9jTV",
						Name: "Cash in",
					},
				},
			},
			CashFlowOfPersonalAcc: _cashFlowOfPersonalAccTable{
				_codaEntity: _codaEntity{
					ID:   "table-dMo-URJL7Z",
					Name: "Cash flow of personal acc",
				},
				Cols: _cashFlowOfPersonalAccTableColumns{
					Account: _codaEntity{
						ID:   "c-HvpjCI4gRX",
						Name: "Account",
					},
					Date: _codaEntity{
						ID:   "c-HuVxndys-5",
						Name: "Date",
					},
					AmountRUB: _codaEntity{
						ID:   "c-Ka5rkoQEvD",
						Name: "Amount, RUB",
					},
					Comment: _codaEntity{
						ID:   "c-7Sr6bn-OVs",
						Name: "Comment",
					},
					CashOutPurpose: _codaEntity{
						ID:   "c-ak93zMH1dQ",
						Name: "Cash out purpose",
					},
					PersonalPaidIn: _codaEntity{
						ID:   "c-OPYJnVV_u7",
						Name: "Personal paid in",
					},
				},
			},
			CashFlow2021: _cashFlow2021Table{
				_codaEntity: _codaEntity{
					ID:   "table-79if8pVCSx",
					Name: "Cash flow 2021",
				},
				Cols: _cashFlow2021TableColumns{
					Author: _codaEntity{
						ID:   "c-2Xkn5hlxUM",
						Name: "Author",
					},
					Reconciled: _codaEntity{
						ID:   "c-P_CSVSFk0L",
						Name: "Reconciled",
					},
					Account: _codaEntity{
						ID:   "c-HvpjCI4gRX",
						Name: "Account",
					},
					Date: _codaEntity{
						ID:   "c-HuVxndys-5",
						Name: "Date",
					},
					AmountRUB: _codaEntity{
						ID:   "c-Ka5rkoQEvD",
						Name: "Amount, RUB",
					},
					PaidAmtRUB: _codaEntity{
						ID:   "c-NcMKzn_fpI",
						Name: "Paid amt, RUB",
					},
					CashOutPurpose: _codaEntity{
						ID:   "c-ak93zMH1dQ",
						Name: "Cash out purpose",
					},
					Count: _codaEntity{
						ID:   "c-taBOV-t5cz",
						Name: "Count",
					},
					ComissionAmtRUB: _codaEntity{
						ID:   "c-6pxTqY1M65",
						Name: "Comission amt, RUB",
					},
					Comment: _codaEntity{
						ID:   "c-7Sr6bn-OVs",
						Name: "Comment",
					},
					CashIn: _codaEntity{
						ID:   "c-FfATqG9jTV",
						Name: "Cash in",
					},
				},
			},
			AuditCategory: _auditCategoryTable{
				_codaEntity: _codaEntity{
					ID:   "grid-cXSuejWcpk",
					Name: "Audit Category",
				},
				Cols: _auditCategoryTableColumns{
					AuditCategory: _codaEntity{
						ID:   "c-IGinJPS6nr",
						Name: "Audit Category",
					},
				},
			},
			ExpensesByCategory: _expensesByCategoryTable{
				_codaEntity: _codaEntity{
					ID:   "table--3kR_z-HTW",
					Name: "Expenses by category",
				},
				Cols: _expensesByCategoryTableColumns{
					Group: _codaEntity{
						ID:   "c-18G6IQXBvd",
						Name: "Group",
					},
					AmountRUB: _codaEntity{
						ID:   "c-NWO2JHrrdw",
						Name: "Amount, RUB",
					},
					AmountEUR: _codaEntity{
						ID:   "c-85jKBvNSVF",
						Name: "Amount, EUR",
					},
					AuditCategory: _codaEntity{
						ID:   "c-45roqK-a18",
						Name: "Audit Category",
					},
				},
			},
		},
		Formula: _formulaSchema{
			PlannedInvoice: _codaEntity{
				ID:   "f-6rR6RS0Un_",
				Name: "plannedInvoice",
			},
			LastInvoice: _codaEntity{
				ID:   "f-hZIdsBjMw3",
				Name: "lastInvoice",
			},
			EvaluatedCompanyBalance: _codaEntity{
				ID:   "f-9zMPywtJhn",
				Name: "evaluatedCompanyBalance",
			},
			EvaluatedTotalPending: _codaEntity{
				ID:   "f-qRpVm7GQ8F",
				Name: "evaluatedTotalPending",
			},
			TplMonth: _codaEntity{
				ID:   "f-wEKm-kjx1t",
				Name: "tplMonth",
			},
			TplNextMonth: _codaEntity{
				ID:   "f-wR0fPfSRvf",
				Name: "tplNextMonth",
			},
		},
		Control: _controlSchema{
			BtnAddUnplanned: _codaEntity{
				ID:   "ctrl-Sp_u7Cjsqc",
				Name: "btnAddUnplanned",
			},
			BtnInvoiceAdd: _codaEntity{
				ID:   "ctrl-izPLCxLtdl",
				Name: "btnInvoiceAdd",
			},
			Button1: _codaEntity{
				ID:   "ctrl-NTVAm6CmKb",
				Name: "button 1",
			},
		},
	}
}
