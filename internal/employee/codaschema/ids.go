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

func (e *_codaEntity) String() string {
	return e.ID
}

type _tableSchema struct {
	AllEmployees                      _allEmployeesTable                      // All employees
	Invoice                           _invoiceTable                           // Invoice
	Months                            _monthsTable                            // Months
	Corrections                       _correctionsTable                       // Corrections
	WorkingEmployees                  _workingEmployeesTable                  // Working Employees
	WorkingEmployeesNames             _workingEmployeesNamesTable             // Working employees names
	AllSalaries                       _allSalariesTable                       // All salaries
	InvoicesOverview                  _invoicesOverviewTable                  // Invoices overview
	Taxes                             _taxesTable                             // Taxes
	PatentPayments                    _patentPaymentsTable                    // Patent Payments
	CorrectionCategory                _correctionCategoryTable                // Correction Category
	TaxesOverview                     _taxesOverviewTable                     // Taxes overview
	PatentsOverview                   _patentsOverviewTable                   // Patents overview
	LegalEntity                       _legalEntityTable                       // Legal entity
	SalariesReport                    _salariesReportTable                    // Salaries report
	TaxYears                          _taxYearsTable                          // Tax years
	EmployeePatents                   _employeePatentsTable                   // Employee Patents
	FullPayrollReport                 _fullPayrollReportTable                 // Full payroll report
	InvoicableEmployees               _invoicableEmployeesTable               // Invoicable Employees
	EmptyActualRates                  _emptyActualRatesTable                  // Empty actual rates
	CorrectionsByType                 _correctionsByTypeTable                 // Corrections by type
	WorkingSalaries                   _workingSalariesTable                   // Working Salaries
	PayrollReportToAdd                _payrollReportToAddTable                // Payroll report to add
	BankFees                          _bankFeesTable                          // Bank fees
	PHCorrections                     _pHCorrectionsTable                     // PH corrections
	PHMonthlyReport                   _pHMonthlyReportTable                   // PH Monthly Report
	Location                          _locationTable                          // Location
	PaymentNotesCorrectionAdding      _paymentNotesCorrectionAddingTable      // Payment Notes Correction Adding
	CompanyRates                      _companyRatesTable                      // Company rates
	SummaryForCurrentPeriod3          _summaryForCurrentPeriod3Table          // Summary for current period 3
	CurrentPayrollV2                  _currentPayrollV2Table                  // Current payroll v2
	HeadcountPerLocation              _headcountPerLocationTable              // Headcount per location
	SuspiciousRates                   _suspiciousRatesTable                   // Suspicious rates
	WorkDays                          _workDaysTable                          // Work days
	CorrectionsByEmployee             _correctionsByEmployeeTable             // Corrections by employee
	EnterCompanyRatesForSelectedMonth _enterCompanyRatesForSelectedMonthTable // Enter company rates for selected month
	CorrectionTemplates               _correctionTemplatesTable               // Correction Templates
	AllEmployeesNames                 _allEmployeesNamesTable                 // All employees names
	BankDetails                       _bankDetailsTable                       // Bank details
	BeneficiaryBank                   _beneficiaryBankTable                   // Beneficiary Bank
	PayrollSchedule                   _payrollScheduleTable                   // Payroll Schedule
}
type _formulaSchema struct {
	CurrentMonth         _codaEntity // currentMonth
	PayrollReportCurRate _codaEntity // payrollReportCurRate
	InvoiceAddingRate    _codaEntity // invoiceAddingRate
}
type _controlSchema struct {
	SelectOverviewMonth     _codaEntity // selectOverviewMonth
	SelectOverviewEmployee  _codaEntity // selectOverviewEmployee
	WorkingSalariesCheckbox _codaEntity // workingSalariesCheckbox
	InvoiceAddingMonth      _codaEntity // invoiceAddingMonth
	PayrollReportCurPeriod  _codaEntity // payrollReportCurPeriod
	PayrollReportAddReport  _codaEntity // payrollReportAddReport
	FillRatesShowFilled     _codaEntity // fillRatesShowFilled
	ChkSalariesOnlyCurrent  _codaEntity // chkSalariesOnlyCurrent
}

// Table All employees
type _allEmployeesTable struct {
	_codaEntity
	Cols _allEmployeesTableColumns
}
type _allEmployeesTableColumns struct {
	Name                    _codaEntity // Name
	StartDate               _codaEntity // Start date
	ProbationEnd            _codaEntity // Probation end
	AnnualLeaveFrom         _codaEntity // Annual leave from
	SalaryBeforeIP          _codaEntity // Salary before IP
	SalaryAfterIP           _codaEntity // Salary after IP
	NetSalaryAfterProbation _codaEntity // Net salary after probation
	EndDate                 _codaEntity // End date
	HourRate                _codaEntity // Hour rate
	BankCC                  _codaEntity // Bank:CC
	BankService             _codaEntity // Bank:Service
	BankTotalFees           _codaEntity // Bank:Total Fees
	Address                 _codaEntity // Address
	OpeningDateIP           _codaEntity // Opening date IP
	StartMonth              _codaEntity // Start month
	MattermostLogin         _codaEntity // Mattermost login
	RussianFullName         _codaEntity // Russian full name
	Position                _codaEntity // Position
	INN                     _codaEntity // INN
	WorkingNow              _codaEntity // Working now?
	AllSalaries             _codaEntity // All salaries
	CurrentSalaryNet        _codaEntity // Current Salary Net
	EnglishFullName         _codaEntity // English full name
	BankRequisites          _codaEntity // Bank requisites
	BillTo                  _codaEntity // Bill To
	LegalEntity             _codaEntity // Legal entity
	Location                _codaEntity // Location
	PaymentNotes            _codaEntity // Payment notes
	ACLSD                   _codaEntity // ACL SD
	FinanceSD               _codaEntity // Finance SD
	GeneralSD               _codaEntity // General SD
	PersonnelSD             _codaEntity // Personnel SD
	ContractNumber          _codaEntity // Contract Number
	ContractDate            _codaEntity // Contract Date
	LegalForm               _codaEntity // Legal form
	InvoiceAdd              _codaEntity // Invoice Add
	PendingInvoice          _codaEntity // Pending Invoice
	InvoiceFieldsErrors     _codaEntity // Invoice fields errors
	PaymentFieldsErrors     _codaEntity // Payment fields errors
	Bank                    _codaEntity // Bank
	Rounding                _codaEntity // Rounding
	BankDetails             _codaEntity // Bank details
}

// Table Invoice
type _invoiceTable struct {
	_codaEntity
	Cols _invoiceTableColumns
}
type _invoiceTableColumns struct {
	ID                   _codaEntity // ID
	InvoiceHash          _codaEntity // Invoice #
	Month                _codaEntity // Month
	Employee             _codaEntity // Employee
	PreviousInvoice      _codaEntity // Previous invoice
	EURRUBExpected       _codaEntity // EURRUB expected
	RequestedSubtotalEUR _codaEntity // Requested subtotal, EUR
	RoundingPrevMonEUR   _codaEntity // Rounding PrevMon, EUR
	Rounding             _codaEntity // Rounding
	RequestedTotalEUR    _codaEntity // Requested total, EUR
	Hours                _codaEntity // Hours
	EURRUBActual         _codaEntity // EURRUB actual
	AmountRUBActual      _codaEntity // Amount RUB actual
	RateErrorRUB         _codaEntity // Rate Error, RUB
	CostOfDay            _codaEntity // Cost of day
	OpeningDateIP        _codaEntity // Opening date IP
	CorrectionRefs       _codaEntity // Correction Refs
	CorrectionRUB        _codaEntity // Correction, RUB
	PatentRUB            _codaEntity // Patent, RUB
	TaxesRUB             _codaEntity // Taxes, RUB
	PatentRefs           _codaEntity // Patent Refs
	TaxesRefs            _codaEntity // Taxes Refs
	BaseSalaryRUB        _codaEntity // Base Salary, RUB
	BankFees             _codaEntity // Bank fees
	Correction           _codaEntity // Correction
	RateErrorPM          _codaEntity // Rate error PM
	Alerts               _codaEntity // Alerts
	Reconciled           _codaEntity // Reconciled
	ToPay                _codaEntity // To Pay
	CurrentSalaryRef     _codaEntity // Current Salary Ref
	PaymentNotes         _codaEntity // Payment notes
	AddToPayrollReport   _codaEntity // Add to payroll report
	RUBSubtotal          _codaEntity // RUB subtotal
	GeneralSD            _codaEntity // General SD
	WorkDaysRef          _codaEntity // Work days ref
	CompanyRate          _codaEntity // Company rate
	BaseSalaryEUR        _codaEntity // Base Salary, EUR
	Template             _codaEntity // Template
	TemplatesRefs        _codaEntity // Templates Refs
	BankDetails          _codaEntity // Bank details
}

// Table Months
type _monthsTable struct {
	_codaEntity
	Cols _monthsTableColumns
}
type _monthsTableColumns struct {
	Month             _codaEntity // Month
	Year              _codaEntity // Year
	ID                _codaEntity // ID
	PreviousMonthLink _codaEntity // Previous month Link
	PreviousMonth     _codaEntity // Previous month
	Current           _codaEntity // Current?
	Near              _codaEntity // Near?
	Future            _codaEntity // Future?
	Upcoming          _codaEntity // Upcoming?
}

// Table Corrections
type _correctionsTable struct {
	_codaEntity
	Cols _correctionsTableColumns
}
type _correctionsTableColumns struct {
	Comment                  _codaEntity // Comment
	AbsoluteCorrectionEUR    _codaEntity // Absolute Correction,EUR
	AbsoluteCorrectionRUB    _codaEntity // Absolute Correction, RUB
	EURRUBExpected           _codaEntity // EURRUB expected
	AbsCorrectionEURInRUB    _codaEntity // Abs Correction, EUR in RUB
	PerDayType               _codaEntity // Per Day Type
	NumberOfDays             _codaEntity // Number of days
	CostOfDay                _codaEntity // Cost of day
	PerDay                   _codaEntity // Per Day
	TotalCorrectionRUB       _codaEntity // Total Correction, RUB
	PaymentInvoice           _codaEntity // Payment Invoice
	PerDayCoefficient        _codaEntity // PerDay Coefficient
	PerDayCalculationInvoice _codaEntity // PerDay calculation invoice
	Display                  _codaEntity // Display
	Category                 _codaEntity // Category
	Month                    _codaEntity // Month
	PaymentNotes             _codaEntity // Payment notes
	EmployeeName             _codaEntity // Employee name
	PercentCorrectionPercent _codaEntity // Percent Correction, %
	PercentCorrectionRUB     _codaEntity // Percent Correction, RUB
	TotalCorrectionEUR       _codaEntity // Total Correction, EUR
	ModifiedOn               _codaEntity // Modified on
	CreatedOn                _codaEntity // Created on
}

// Table Working Employees
type _workingEmployeesTable struct {
	_codaEntity
	Cols _workingEmployeesTableColumns
}
type _workingEmployeesTableColumns struct {
	Name                _codaEntity // Name
	CurrentSalaryNet    _codaEntity // Current Salary Net
	StartDate           _codaEntity // Start date
	GeneralSD           _codaEntity // General SD
	PersonnelSD         _codaEntity // Personnel SD
	FinanceSD           _codaEntity // Finance SD
	ACLSD               _codaEntity // ACL SD
	ProbationEnd        _codaEntity // Probation end
	AnnualLeaveFrom     _codaEntity // Annual leave from
	Location            _codaEntity // Location
	LegalEntity         _codaEntity // Legal entity
	LegalForm           _codaEntity // Legal form
	Position            _codaEntity // Position
	Bank                _codaEntity // Bank
	EnglishFullName     _codaEntity // English full name
	InvoiceFieldsErrors _codaEntity // Invoice fields errors
	PaymentFieldsErrors _codaEntity // Payment fields errors
	ContractNumber      _codaEntity // Contract Number
	ContractDate        _codaEntity // Contract Date
	BankDetails         _codaEntity // Bank details
	BankRequisites      _codaEntity // Bank requisites
	BankTotalFees       _codaEntity // Bank:Total Fees
	PaymentNotes        _codaEntity // Payment notes
	BankCC              _codaEntity // Bank:CC
	BankService         _codaEntity // Bank:Service
}

// Table Working employees names
type _workingEmployeesNamesTable struct {
	_codaEntity
	Cols _workingEmployeesNamesTableColumns
}
type _workingEmployeesNamesTableColumns struct {
	Name      _codaEntity // Name
	Position  _codaEntity // Position
	StartDate _codaEntity // Start date
	Location  _codaEntity // Location
}

// Table All salaries
type _allSalariesTable struct {
	_codaEntity
	Cols _allSalariesTableColumns
}
type _allSalariesTableColumns struct {
	Employee       _codaEntity // Employee
	PeriodFrom     _codaEntity // Period from
	PeriodTo       _codaEntity // Period to
	Salary         _codaEntity // Salary
	PeriodFromReal _codaEntity // Period from real
	PeriodToReal   _codaEntity // Period to real
	ModifiedBy     _codaEntity // Modified by
	ModifiedOn     _codaEntity // Modified on
	CreatedOn      _codaEntity // Created on
	Currency       _codaEntity // Currency
	IsRUB          _codaEntity // IsRUB
	IsEUR          _codaEntity // IsEUR
}

// Table Invoices overview
type _invoicesOverviewTable struct {
	_codaEntity
	Cols _invoicesOverviewTableColumns
}
type _invoicesOverviewTableColumns struct {
	Employee             _codaEntity // Employee
	Correction           _codaEntity // Correction
	Template             _codaEntity // Template
	Month                _codaEntity // Month
	ToPay                _codaEntity // To Pay
	Reconciled           _codaEntity // Reconciled
	Alerts               _codaEntity // Alerts
	RUBSubtotal          _codaEntity // RUB subtotal
	BaseSalaryRUB        _codaEntity // Base Salary, RUB
	BaseSalaryEUR        _codaEntity // Base Salary, EUR
	BankFees             _codaEntity // Bank fees
	RateErrorPM          _codaEntity // Rate error PM
	CorrectionRUB        _codaEntity // Correction, RUB
	CorrectionRefs       _codaEntity // Correction Refs
	PatentRUB            _codaEntity // Patent, RUB
	TaxesRUB             _codaEntity // Taxes, RUB
	EURRUBExpected       _codaEntity // EURRUB expected
	EURRUBActual         _codaEntity // EURRUB actual
	RequestedSubtotalEUR _codaEntity // Requested subtotal, EUR
	RoundingPrevMonEUR   _codaEntity // Rounding PrevMon, EUR
	Rounding             _codaEntity // Rounding
	RequestedTotalEUR    _codaEntity // Requested total, EUR
	RateErrorRUB         _codaEntity // Rate Error, RUB
	CostOfDay            _codaEntity // Cost of day
	BankDetails          _codaEntity // Bank details
}

// Table Taxes
type _taxesTable struct {
	_codaEntity
	Cols _taxesTableColumns
}
type _taxesTableColumns struct {
	Invoice               _codaEntity // Invoice
	OpeningDateIP         _codaEntity // Opening date IP
	PeriodStart           _codaEntity // Period Start
	PeriodEnd             _codaEntity // Period End
	AmountIPDays          _codaEntity // Amount IP days
	SocialInsuranceToPay  _codaEntity // Social Insurance - to pay
	PensionFundFixedToPay _codaEntity // Pension fund fixed - to pay
	PensionFundPercent    _codaEntity // Pension fund percent
	Amount                _codaEntity // Amount
	SocialInsuranceTotal  _codaEntity // Social Insurance - total
	PensionFundFixedTotal _codaEntity // Pension Fund fixed - total
	Year                  _codaEntity // Year
}

// Table Patent Payments
type _patentPaymentsTable struct {
	_codaEntity
	Cols _patentPaymentsTableColumns
}
type _patentPaymentsTableColumns struct {
	Invoice           _codaEntity // Invoice
	PeriodCost        _codaEntity // Period cost
	Period            _codaEntity // Period
	PeriodCostManual  _codaEntity // Period cost manual
	EmployeePatentRef _codaEntity // Employee patent ref
}

// Table Correction Category
type _correctionCategoryTable struct {
	_codaEntity
	Cols _correctionCategoryTableColumns
}
type _correctionCategoryTableColumns struct {
	Category _codaEntity // Category
	Comment  _codaEntity // Comment
}

// Table Taxes overview
type _taxesOverviewTable struct {
	_codaEntity
	Cols _taxesOverviewTableColumns
}
type _taxesOverviewTableColumns struct {
	Invoice               _codaEntity // Invoice
	Amount                _codaEntity // Amount
	OpeningDateIP         _codaEntity // Opening date IP
	PeriodStart           _codaEntity // Period Start
	PeriodEnd             _codaEntity // Period End
	AmountIPDays          _codaEntity // Amount IP days
	SocialInsuranceToPay  _codaEntity // Social Insurance - to pay
	PensionFundFixedToPay _codaEntity // Pension fund fixed - to pay
	PensionFundPercent    _codaEntity // Pension fund percent
}

// Table Patents overview
type _patentsOverviewTable struct {
	_codaEntity
	Cols _patentsOverviewTableColumns
}
type _patentsOverviewTableColumns struct {
	Invoice    _codaEntity // Invoice
	PeriodCost _codaEntity // Period cost
	Period     _codaEntity // Period
}

// Table Legal entity
type _legalEntityTable struct {
	_codaEntity
	Cols _legalEntityTableColumns
}
type _legalEntityTableColumns struct {
	EntityName    _codaEntity // Entity name
	Requisites    _codaEntity // Requisites
	Taxable       _codaEntity // Taxable
	OfficialName  _codaEntity // Official name
	AccountNumber _codaEntity // Account number
}

// Table Salaries report
type _salariesReportTable struct {
	_codaEntity
	Cols _salariesReportTableColumns
}
type _salariesReportTableColumns struct {
	Name             _codaEntity // Name
	StartDate        _codaEntity // Start date
	ProbationEnd     _codaEntity // Probation end
	CurrentSalaryNet _codaEntity // Current Salary Net
	Position         _codaEntity // Position
}

// Table Tax years
type _taxYearsTable struct {
	_codaEntity
	Cols _taxYearsTableColumns
}
type _taxYearsTableColumns struct {
	Year                   _codaEntity // Year
	PensionFundFixedPart   _codaEntity // Pension fund fixed part
	SocialInsurance        _codaEntity // Social Insurance
	IsCurrent              _codaEntity // Is Current
	PensionFundPercentPart _codaEntity // Pension fund percent part
	CalendarDaysInYear     _codaEntity // Calendar days in year
}

// Table Employee Patents
type _employeePatentsTable struct {
	_codaEntity
	Cols _employeePatentsTableColumns
}
type _employeePatentsTableColumns struct {
	Employee     _codaEntity // Employee
	Year         _codaEntity // Year
	Cost         _codaEntity // Cost
	Scan         _codaEntity // Scan
	Date1        _codaEntity // Date 1
	Date1Payment _codaEntity // Date 1 Payment
	Date2        _codaEntity // Date 2
	Date2Payment _codaEntity // Date 2 Payment
	Display      _codaEntity // Display
}

// Table Full payroll report
type _fullPayrollReportTable struct {
	_codaEntity
	Cols _fullPayrollReportTableColumns
}
type _fullPayrollReportTableColumns struct {
	SalaryNet             _codaEntity // Salary Net
	InvoiceRef            _codaEntity // Invoice Ref
	OvertimesRefs         _codaEntity // Overtimes Refs
	OvertimesRUB          _codaEntity // Overtimes, RUB
	BonusesRefs           _codaEntity // Bonuses Refs
	BonusesRUB            _codaEntity // Bonuses, RUB
	PublicHolidaysRefs    _codaEntity // Public Holidays Refs
	PublicHolidaysRUB     _codaEntity // Public holidays, RUB
	RestCorrectionRefs    _codaEntity // Rest correction Refs
	RestCorrectionsRUB    _codaEntity // Rest corrections, RUB
	AnnualContribsRUB     _codaEntity // Annual Contribs, RUB
	StartDayRefs          _codaEntity // Start day Refs
	StartDayRUB           _codaEntity // Start day, RUB
	SalaryChangesRefs     _codaEntity // Salary changes Refs
	SalaryChangesRUB      _codaEntity // Salary changes, RUB
	LoansRefs             _codaEntity // Loans Refs
	LoansRUB              _codaEntity // Loans, RUB
	RUBServiceFees        _codaEntity // RUB Service fees
	EURRounding           _codaEntity // EUR Rounding
	AnnualUnpaidLeavesRUB _codaEntity // Annual/unpaid leaves, RUB
	MonthlyContribsEUR    _codaEntity // Monthly Contribs, EUR
	QuitCompensationsRefs _codaEntity // Quit compensations Refs
	QuitCompensationsRUB  _codaEntity // Quit compensations, RUB
	PaidAnnualLeavesRefs  _codaEntity // Paid annual leaves Refs
	Location              _codaEntity // Location
}

// Table Invoicable Employees
type _invoicableEmployeesTable struct {
	_codaEntity
	Cols _invoicableEmployeesTableColumns
}
type _invoicableEmployeesTableColumns struct {
	Name                _codaEntity // Name
	LegalForm           _codaEntity // Legal form
	InvoiceAdd          _codaEntity // Invoice Add
	PendingInvoice      _codaEntity // Pending Invoice
	InvoiceFieldsErrors _codaEntity // Invoice fields errors
	PaymentFieldsErrors _codaEntity // Payment fields errors
}

// Table Empty actual rates
type _emptyActualRatesTable struct {
	_codaEntity
	Cols _emptyActualRatesTableColumns
}
type _emptyActualRatesTableColumns struct {
	Month          _codaEntity // Month
	Employee       _codaEntity // Employee
	GeneralSD      _codaEntity // General SD
	EURRUBExpected _codaEntity // EURRUB expected
	EURRUBActual   _codaEntity // EURRUB actual
}

// Table Corrections by type
type _correctionsByTypeTable struct {
	_codaEntity
	Cols _correctionsByTypeTableColumns
}
type _correctionsByTypeTableColumns struct {
	Month                    _codaEntity // Month
	PaymentInvoice           _codaEntity // Payment Invoice
	Category                 _codaEntity // Category
	Comment                  _codaEntity // Comment
	TotalCorrectionRUB       _codaEntity // Total Correction, RUB
	AbsoluteCorrectionRUB    _codaEntity // Absolute Correction, RUB
	AbsoluteCorrectionEUR    _codaEntity // Absolute Correction,EUR
	PercentCorrectionPercent _codaEntity // Percent Correction, %
	PerDayType               _codaEntity // Per Day Type
	PerDayCoefficient        _codaEntity // PerDay Coefficient
	PerDayCalculationInvoice _codaEntity // PerDay calculation invoice
	NumberOfDays             _codaEntity // Number of days
	PerDay                   _codaEntity // Per Day
}

// Table Working Salaries
type _workingSalariesTable struct {
	_codaEntity
	Cols _workingSalariesTableColumns
}
type _workingSalariesTableColumns struct {
	Employee   _codaEntity // Employee
	PeriodFrom _codaEntity // Period from
	PeriodTo   _codaEntity // Period to
	Salary     _codaEntity // Salary
	Currency   _codaEntity // Currency
	IsRUB      _codaEntity // IsRUB
	IsEUR      _codaEntity // IsEUR
	ModifiedBy _codaEntity // Modified by
	CreatedOn  _codaEntity // Created on
	ModifiedOn _codaEntity // Modified on
}

// Table Payroll report to add
type _payrollReportToAddTable struct {
	_codaEntity
	Cols _payrollReportToAddTableColumns
}
type _payrollReportToAddTableColumns struct {
	Month              _codaEntity // Month
	Employee           _codaEntity // Employee
	AddToPayrollReport _codaEntity // Add to payroll report
}

// Table Bank fees
type _bankFeesTable struct {
	_codaEntity
	Cols _bankFeesTableColumns
}
type _bankFeesTableColumns struct {
	Name              _codaEntity // Name
	DefaultCCFee      _codaEntity // Default CC Fee
	DefaultServiceFee _codaEntity // Default Service Fee
}

// Table PH corrections
type _pHCorrectionsTable struct {
	_codaEntity
	Cols _pHCorrectionsTableColumns
}
type _pHCorrectionsTableColumns struct {
	Month              _codaEntity // Month
	Comment            _codaEntity // Comment
	TotalCorrectionRUB _codaEntity // Total Correction, RUB
	EURRUBExpected     _codaEntity // EURRUB expected
	TotalCorrectionEUR _codaEntity // Total Correction, EUR
}

// Table PH Monthly Report
type _pHMonthlyReportTable struct {
	_codaEntity
	Cols _pHMonthlyReportTableColumns
}
type _pHMonthlyReportTableColumns struct {
	Month              _codaEntity // Month
	TotalCorrectionRUB _codaEntity // Total Correction, RUB
	TotalCorrectionEUR _codaEntity // Total Correction, EUR
}

// Table Location
type _locationTable struct {
	_codaEntity
	Cols _locationTableColumns
}
type _locationTableColumns struct {
	Location _codaEntity // Location
}

// Table Payment Notes Correction Adding
type _paymentNotesCorrectionAddingTable struct {
	_codaEntity
	Cols _paymentNotesCorrectionAddingTableColumns
}
type _paymentNotesCorrectionAddingTableColumns struct {
	Month          _codaEntity // Month
	Employee       _codaEntity // Employee
	Correction     _codaEntity // Correction
	CorrectionRefs _codaEntity // Correction Refs
	PaymentNotes   _codaEntity // Payment notes
}

// Table Company rates
type _companyRatesTable struct {
	_codaEntity
	Cols _companyRatesTableColumns
}
type _companyRatesTableColumns struct {
	Month _codaEntity // Month
	Rate  _codaEntity // Rate
}

// Table Summary for current period 3
type _summaryForCurrentPeriod3Table struct {
	_codaEntity
	Cols _summaryForCurrentPeriod3TableColumns
}
type _summaryForCurrentPeriod3TableColumns struct {
	Subject      _codaEntity // Subject
	RUB          _codaEntity // RUB
	EUR          _codaEntity // EUR
	Location     _codaEntity // Location
	LocationRefs _codaEntity // Location Refs
}

// Table Current payroll v2
type _currentPayrollV2Table struct {
	_codaEntity
	Cols _currentPayrollV2TableColumns
}
type _currentPayrollV2TableColumns struct {
	InvoiceRef            _codaEntity // Invoice Ref
	Location              _codaEntity // Location
	SalaryNet             _codaEntity // Salary Net
	RUBServiceFees        _codaEntity // RUB Service fees
	EURRounding           _codaEntity // EUR Rounding
	AnnualUnpaidLeavesRUB _codaEntity // Annual/unpaid leaves, RUB
	AnnualContribsRUB     _codaEntity // Annual Contribs, RUB
	MonthlyContribsEUR    _codaEntity // Monthly Contribs, EUR
	PublicHolidaysRefs    _codaEntity // Public Holidays Refs
	PublicHolidaysRUB     _codaEntity // Public holidays, RUB
	OvertimesRefs         _codaEntity // Overtimes Refs
	OvertimesRUB          _codaEntity // Overtimes, RUB
	BonusesRefs           _codaEntity // Bonuses Refs
	BonusesRUB            _codaEntity // Bonuses, RUB
	LoansRefs             _codaEntity // Loans Refs
	LoansRUB              _codaEntity // Loans, RUB
	SalaryChangesRefs     _codaEntity // Salary changes Refs
	SalaryChangesRUB      _codaEntity // Salary changes, RUB
	StartDayRefs          _codaEntity // Start day Refs
	StartDayRUB           _codaEntity // Start day, RUB
	QuitCompensationsRefs _codaEntity // Quit compensations Refs
	QuitCompensationsRUB  _codaEntity // Quit compensations, RUB
	PaidAnnualLeavesRefs  _codaEntity // Paid annual leaves Refs
	RestCorrectionRefs    _codaEntity // Rest correction Refs
	RestCorrectionsRUB    _codaEntity // Rest corrections, RUB
}

// Table Headcount per location
type _headcountPerLocationTable struct {
	_codaEntity
	Cols _headcountPerLocationTableColumns
}
type _headcountPerLocationTableColumns struct {
	Location  _codaEntity // Location
	Headcount _codaEntity // Headcount
}

// Table Suspicious rates
type _suspiciousRatesTable struct {
	_codaEntity
	Cols _suspiciousRatesTableColumns
}
type _suspiciousRatesTableColumns struct {
	Month          _codaEntity // Month
	Employee       _codaEntity // Employee
	GeneralSD      _codaEntity // General SD
	EURRUBExpected _codaEntity // EURRUB expected
	EURRUBActual   _codaEntity // EURRUB actual
}

// Table Work days
type _workDaysTable struct {
	_codaEntity
	Cols _workDaysTableColumns
}
type _workDaysTableColumns struct {
	Month    _codaEntity // Month
	WorkDays _codaEntity // Work days
	Location _codaEntity // Location
}

// Table Corrections by employee
type _correctionsByEmployeeTable struct {
	_codaEntity
	Cols _correctionsByEmployeeTableColumns
}
type _correctionsByEmployeeTableColumns struct {
	Month                    _codaEntity // Month
	Category                 _codaEntity // Category
	Comment                  _codaEntity // Comment
	TotalCorrectionRUB       _codaEntity // Total Correction, RUB
	AbsoluteCorrectionRUB    _codaEntity // Absolute Correction, RUB
	AbsoluteCorrectionEUR    _codaEntity // Absolute Correction,EUR
	PercentCorrectionPercent _codaEntity // Percent Correction, %
	PerDayType               _codaEntity // Per Day Type
	PerDayCoefficient        _codaEntity // PerDay Coefficient
	PerDayCalculationInvoice _codaEntity // PerDay calculation invoice
	NumberOfDays             _codaEntity // Number of days
	PerDay                   _codaEntity // Per Day
	EmployeeName             _codaEntity // Employee name
}

// Table Enter company rates for selected month
type _enterCompanyRatesForSelectedMonthTable struct {
	_codaEntity
	Cols _enterCompanyRatesForSelectedMonthTableColumns
}
type _enterCompanyRatesForSelectedMonthTableColumns struct {
	Month _codaEntity // Month
	Rate  _codaEntity // Rate
}

// Table Correction Templates
type _correctionTemplatesTable struct {
	_codaEntity
	Cols _correctionTemplatesTableColumns
}
type _correctionTemplatesTableColumns struct {
	Employee          _codaEntity // Employee
	MonthFrom         _codaEntity // Month from
	MonthTo           _codaEntity // Month to
	Category          _codaEntity // Category
	Comment           _codaEntity // Comment
	TemplateAmountRUB _codaEntity // Template amount, RUB
	TemplateAmountEUR _codaEntity // Template amount, EUR
}

// Table All employees names
type _allEmployeesNamesTable struct {
	_codaEntity
	Cols _allEmployeesNamesTableColumns
}
type _allEmployeesNamesTableColumns struct {
	Name      _codaEntity // Name
	Position  _codaEntity // Position
	StartDate _codaEntity // Start date
	Location  _codaEntity // Location
}

// Table Bank details
type _bankDetailsTable struct {
	_codaEntity
	Cols _bankDetailsTableColumns
}
type _bankDetailsTableColumns struct {
	Employee        _codaEntity // Employee
	Account         _codaEntity // Account
	Address1        _codaEntity // Address 1
	Address2        _codaEntity // Address 2
	BeneficiaryBank _codaEntity // Beneficiary Bank
	MonthFrom       _codaEntity // Month from
	MonthTo         _codaEntity // Month to
	ID              _codaEntity // ID
	Text            _codaEntity // Text
}

// Table Beneficiary Bank
type _beneficiaryBankTable struct {
	_codaEntity
	Cols _beneficiaryBankTableColumns
}
type _beneficiaryBankTableColumns struct {
	Name              _codaEntity // Name
	Address1          _codaEntity // Address 1
	Address2          _codaEntity // Address 2
	Address3          _codaEntity // Address 3
	BeneficiarySWIFT  _codaEntity // Beneficiary SWIFT
	IntermediarySWIFT _codaEntity // Intermediary SWIFT
}

// Table Payroll Schedule
type _payrollScheduleTable struct {
	_codaEntity
	Cols _payrollScheduleTableColumns
}
type _payrollScheduleTableColumns struct {
	Month         _codaEntity // Month
	ExecutionDate _codaEntity // Execution date
}

var ID _schema

func init() {
	ID = _schema{
		Table: _tableSchema{
			AllEmployees: _allEmployeesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-TGESBHJkVA",
					Name: "All employees",
				},
				Cols: _allEmployeesTableColumns{
					Name: _codaEntity{
						ID:   "c-tCDt6yt4Ix",
						Name: "Name",
					},
					StartDate: _codaEntity{
						ID:   "c-Zs7oQbj-_J",
						Name: "Start date",
					},
					ProbationEnd: _codaEntity{
						ID:   "c-35cUDxDnAo",
						Name: "Probation end",
					},
					AnnualLeaveFrom: _codaEntity{
						ID:   "c-uzP6UrJs3a",
						Name: "Annual leave from",
					},
					SalaryBeforeIP: _codaEntity{
						ID:   "c-MZLKSGSxFn",
						Name: "Salary before IP",
					},
					SalaryAfterIP: _codaEntity{
						ID:   "c-5asgm6tqZs",
						Name: "Salary after IP",
					},
					NetSalaryAfterProbation: _codaEntity{
						ID:   "c-fvHtrLntTN",
						Name: "Net salary after probation",
					},
					EndDate: _codaEntity{
						ID:   "c-7OoHuXqt8n",
						Name: "End date",
					},
					HourRate: _codaEntity{
						ID:   "c-SItGvyE3ie",
						Name: "Hour rate",
					},
					BankCC: _codaEntity{
						ID:   "c-HbuzOpvaPf",
						Name: "Bank:CC",
					},
					BankService: _codaEntity{
						ID:   "c-214wDdnGtE",
						Name: "Bank:Service",
					},
					BankTotalFees: _codaEntity{
						ID:   "c-156kDMsJzb",
						Name: "Bank:Total Fees",
					},
					Address: _codaEntity{
						ID:   "c-lEdEHoijqv",
						Name: "Address",
					},
					OpeningDateIP: _codaEntity{
						ID:   "c-uU3-6piESs",
						Name: "Opening date IP",
					},
					StartMonth: _codaEntity{
						ID:   "c-pnvQkCAJ7U",
						Name: "Start month",
					},
					MattermostLogin: _codaEntity{
						ID:   "c-BmjWOCSXHd",
						Name: "Mattermost login",
					},
					RussianFullName: _codaEntity{
						ID:   "c-7N8qU2h-Zv",
						Name: "Russian full name",
					},
					Position: _codaEntity{
						ID:   "c-4nDWkuySVp",
						Name: "Position",
					},
					INN: _codaEntity{
						ID:   "c-Gsx7-a_kGU",
						Name: "INN",
					},
					WorkingNow: _codaEntity{
						ID:   "c-QfV5QzjuJP",
						Name: "Working now?",
					},
					AllSalaries: _codaEntity{
						ID:   "c-P0cTAu016r",
						Name: "All salaries",
					},
					CurrentSalaryNet: _codaEntity{
						ID:   "c-Is6fozVp1a",
						Name: "Current Salary Net",
					},
					EnglishFullName: _codaEntity{
						ID:   "c-TAlfzDcFzQ",
						Name: "English full name",
					},
					BankRequisites: _codaEntity{
						ID:   "c-OlCoWd7n4S",
						Name: "Bank requisites",
					},
					BillTo: _codaEntity{
						ID:   "c-XnSDzWgAgQ",
						Name: "Bill To",
					},
					LegalEntity: _codaEntity{
						ID:   "c-smk4a68He5",
						Name: "Legal entity",
					},
					Location: _codaEntity{
						ID:   "c-WcmDQXPChx",
						Name: "Location",
					},
					PaymentNotes: _codaEntity{
						ID:   "c-PkLgxjZRQL",
						Name: "Payment notes",
					},
					ACLSD: _codaEntity{
						ID:   "c-aKn3TiCF_A",
						Name: "ACL SD",
					},
					FinanceSD: _codaEntity{
						ID:   "c-NeQmHu-raB",
						Name: "Finance SD",
					},
					GeneralSD: _codaEntity{
						ID:   "c-EMbqnhOkkr",
						Name: "General SD",
					},
					PersonnelSD: _codaEntity{
						ID:   "c-6L5oS7LZ4a",
						Name: "Personnel SD",
					},
					ContractNumber: _codaEntity{
						ID:   "c-XX2OkCkOSR",
						Name: "Contract Number",
					},
					ContractDate: _codaEntity{
						ID:   "c-fDIRo1JHHX",
						Name: "Contract Date",
					},
					LegalForm: _codaEntity{
						ID:   "c-w8WTgIbQ7H",
						Name: "Legal form",
					},
					InvoiceAdd: _codaEntity{
						ID:   "c-F8kuFeh0JE",
						Name: "Invoice Add",
					},
					PendingInvoice: _codaEntity{
						ID:   "c-Saw2_PL8gv",
						Name: "Pending Invoice",
					},
					InvoiceFieldsErrors: _codaEntity{
						ID:   "c-EDwQy_l7fS",
						Name: "Invoice fields errors",
					},
					PaymentFieldsErrors: _codaEntity{
						ID:   "c-UcN9cnw7y0",
						Name: "Payment fields errors",
					},
					Bank: _codaEntity{
						ID:   "c-Hx-nCcWt2v",
						Name: "Bank",
					},
					Rounding: _codaEntity{
						ID:   "c-xvCHSSPdHb",
						Name: "Rounding",
					},
					BankDetails: _codaEntity{
						ID:   "c-wSSi6yT6d4",
						Name: "Bank details",
					},
				},
			},
			Invoice: _invoiceTable{
				_codaEntity: _codaEntity{
					ID:   "grid-Wdy6Agpxou",
					Name: "Invoice",
				},
				Cols: _invoiceTableColumns{
					ID: _codaEntity{
						ID:   "c-bZ_nLfZufG",
						Name: "ID",
					},
					InvoiceHash: _codaEntity{
						ID:   "c-eJ2e_cRCaM",
						Name: "Invoice #",
					},
					Month: _codaEntity{
						ID:   "c-wR0IONcxGH",
						Name: "Month",
					},
					Employee: _codaEntity{
						ID:   "c-bbHUhqlbfN",
						Name: "Employee",
					},
					PreviousInvoice: _codaEntity{
						ID:   "c-FQ7rKmbXr6",
						Name: "Previous invoice",
					},
					EURRUBExpected: _codaEntity{
						ID:   "c-tvtGu9juVL",
						Name: "EURRUB expected",
					},
					RequestedSubtotalEUR: _codaEntity{
						ID:   "c-9rnJJZ6gA7",
						Name: "Requested subtotal, EUR",
					},
					RoundingPrevMonEUR: _codaEntity{
						ID:   "c-hLrmDsk89g",
						Name: "Rounding PrevMon, EUR",
					},
					Rounding: _codaEntity{
						ID:   "c-Tri-EGUP_n",
						Name: "Rounding",
					},
					RequestedTotalEUR: _codaEntity{
						ID:   "c-bJpHVxywXD",
						Name: "Requested total, EUR",
					},
					Hours: _codaEntity{
						ID:   "c-KtVV9if8P7",
						Name: "Hours",
					},
					EURRUBActual: _codaEntity{
						ID:   "c-kLIyv9EvyH",
						Name: "EURRUB actual",
					},
					AmountRUBActual: _codaEntity{
						ID:   "c-AxLSgrt7e3",
						Name: "Amount RUB actual",
					},
					RateErrorRUB: _codaEntity{
						ID:   "c-SsHRhKa_uC",
						Name: "Rate Error, RUB",
					},
					CostOfDay: _codaEntity{
						ID:   "c-yJnq9stsgi",
						Name: "Cost of day",
					},
					OpeningDateIP: _codaEntity{
						ID:   "c-hI1iZG3xzY",
						Name: "Opening date IP",
					},
					CorrectionRefs: _codaEntity{
						ID:   "c-tpeCMU21_I",
						Name: "Correction Refs",
					},
					CorrectionRUB: _codaEntity{
						ID:   "c-jNcl4nZe_h",
						Name: "Correction, RUB",
					},
					PatentRUB: _codaEntity{
						ID:   "c-qA_pPM9kuZ",
						Name: "Patent, RUB",
					},
					TaxesRUB: _codaEntity{
						ID:   "c-ug709va8_K",
						Name: "Taxes, RUB",
					},
					PatentRefs: _codaEntity{
						ID:   "c-bkVlencAZt",
						Name: "Patent Refs",
					},
					TaxesRefs: _codaEntity{
						ID:   "c-aYkpi97eXt",
						Name: "Taxes Refs",
					},
					BaseSalaryRUB: _codaEntity{
						ID:   "c-wqNhZf9EQY",
						Name: "Base Salary, RUB",
					},
					BankFees: _codaEntity{
						ID:   "c-sRGR6jYC7g",
						Name: "Bank fees",
					},
					Correction: _codaEntity{
						ID:   "c-cjYMwmv-1m",
						Name: "Correction",
					},
					RateErrorPM: _codaEntity{
						ID:   "c-_9tuuG4RIN",
						Name: "Rate error PM",
					},
					Alerts: _codaEntity{
						ID:   "c-jkrvvE_SNm",
						Name: "Alerts",
					},
					Reconciled: _codaEntity{
						ID:   "c--A1RtCy-NA",
						Name: "Reconciled",
					},
					ToPay: _codaEntity{
						ID:   "c-DRPGK3XTmD",
						Name: "To Pay",
					},
					CurrentSalaryRef: _codaEntity{
						ID:   "c-KWyHMAHqoq",
						Name: "Current Salary Ref",
					},
					PaymentNotes: _codaEntity{
						ID:   "c-f4weK3W_gH",
						Name: "Payment notes",
					},
					AddToPayrollReport: _codaEntity{
						ID:   "c-FqZuQ1JQN0",
						Name: "Add to payroll report",
					},
					RUBSubtotal: _codaEntity{
						ID:   "c-WLMdyl4WqI",
						Name: "RUB subtotal",
					},
					GeneralSD: _codaEntity{
						ID:   "c-Gkt52a0fjj",
						Name: "General SD",
					},
					WorkDaysRef: _codaEntity{
						ID:   "c-C5AayqIolt",
						Name: "Work days ref",
					},
					CompanyRate: _codaEntity{
						ID:   "c-lsF3Cc0ng3",
						Name: "Company rate",
					},
					BaseSalaryEUR: _codaEntity{
						ID:   "c-B0nU6ZI_Z4",
						Name: "Base Salary, EUR",
					},
					Template: _codaEntity{
						ID:   "c-z0Irm78xhQ",
						Name: "Template",
					},
					TemplatesRefs: _codaEntity{
						ID:   "c-esVje33Tt-",
						Name: "Templates Refs",
					},
					BankDetails: _codaEntity{
						ID:   "c-2TxWM7vHls",
						Name: "Bank details",
					},
				},
			},
			Months: _monthsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-laH8qsdDyP",
					Name: "Months",
				},
				Cols: _monthsTableColumns{
					Month: _codaEntity{
						ID:   "c-u_Pgrgevw7",
						Name: "Month",
					},
					Year: _codaEntity{
						ID:   "c-NgPwXZshJM",
						Name: "Year",
					},
					ID: _codaEntity{
						ID:   "c-iRCjZ0JBcM",
						Name: "ID",
					},
					PreviousMonthLink: _codaEntity{
						ID:   "c-3Cc_lYdvmW",
						Name: "Previous month Link",
					},
					PreviousMonth: _codaEntity{
						ID:   "c-vuW159vf-o",
						Name: "Previous month",
					},
					Current: _codaEntity{
						ID:   "c-OLxcAJLQoW",
						Name: "Current?",
					},
					Near: _codaEntity{
						ID:   "c-2MiYICcxrc",
						Name: "Near?",
					},
					Future: _codaEntity{
						ID:   "c-16wzVYB6nB",
						Name: "Future?",
					},
					Upcoming: _codaEntity{
						ID:   "c-PQzLA1gRk1",
						Name: "Upcoming?",
					},
				},
			},
			Corrections: _correctionsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-wBmvgFgaGi",
					Name: "Corrections",
				},
				Cols: _correctionsTableColumns{
					Comment: _codaEntity{
						ID:   "c--_r48PQnSn",
						Name: "Comment",
					},
					AbsoluteCorrectionEUR: _codaEntity{
						ID:   "c-pRmEece9pf",
						Name: "Absolute Correction,EUR",
					},
					AbsoluteCorrectionRUB: _codaEntity{
						ID:   "c-P2x5IJuMXN",
						Name: "Absolute Correction, RUB",
					},
					EURRUBExpected: _codaEntity{
						ID:   "c-8LD34cnmCh",
						Name: "EURRUB expected",
					},
					AbsCorrectionEURInRUB: _codaEntity{
						ID:   "c-_GvG9w3Qs7",
						Name: "Abs Correction, EUR in RUB",
					},
					PerDayType: _codaEntity{
						ID:   "c-3Ivn-M1j7-",
						Name: "Per Day Type",
					},
					NumberOfDays: _codaEntity{
						ID:   "c-gDOyigH1cm",
						Name: "Number of days",
					},
					CostOfDay: _codaEntity{
						ID:   "c-K_Iy0iERKR",
						Name: "Cost of day",
					},
					PerDay: _codaEntity{
						ID:   "c-Y2E1Vwe2_-",
						Name: "Per Day",
					},
					TotalCorrectionRUB: _codaEntity{
						ID:   "c-0arkfr4qXv",
						Name: "Total Correction, RUB",
					},
					PaymentInvoice: _codaEntity{
						ID:   "c-7SU0iOBY9J",
						Name: "Payment Invoice",
					},
					PerDayCoefficient: _codaEntity{
						ID:   "c-pz6W2IRzFR",
						Name: "PerDay Coefficient",
					},
					PerDayCalculationInvoice: _codaEntity{
						ID:   "c-bK4qXZUCqs",
						Name: "PerDay calculation invoice",
					},
					Display: _codaEntity{
						ID:   "c-FVW_9PPzZ2",
						Name: "Display",
					},
					Category: _codaEntity{
						ID:   "c-zDY58PF0P6",
						Name: "Category",
					},
					Month: _codaEntity{
						ID:   "c-nK3Ad6VblR",
						Name: "Month",
					},
					PaymentNotes: _codaEntity{
						ID:   "c-wa7LZUCn0b",
						Name: "Payment notes",
					},
					EmployeeName: _codaEntity{
						ID:   "c-OrAvunmGm4",
						Name: "Employee name",
					},
					PercentCorrectionPercent: _codaEntity{
						ID:   "c-Wh_7DYRQYP",
						Name: "Percent Correction, %",
					},
					PercentCorrectionRUB: _codaEntity{
						ID:   "c-OCz4oqs3YR",
						Name: "Percent Correction, RUB",
					},
					TotalCorrectionEUR: _codaEntity{
						ID:   "c-U8Bjm-cpmW",
						Name: "Total Correction, EUR",
					},
					ModifiedOn: _codaEntity{
						ID:   "c-7HSE88Op0l",
						Name: "Modified on",
					},
					CreatedOn: _codaEntity{
						ID:   "c-pKkFDJQXHe",
						Name: "Created on",
					},
				},
			},
			WorkingEmployees: _workingEmployeesTable{
				_codaEntity: _codaEntity{
					ID:   "table-XbyycGwUU0",
					Name: "Working Employees",
				},
				Cols: _workingEmployeesTableColumns{
					Name: _codaEntity{
						ID:   "c-tCDt6yt4Ix",
						Name: "Name",
					},
					CurrentSalaryNet: _codaEntity{
						ID:   "c-Is6fozVp1a",
						Name: "Current Salary Net",
					},
					StartDate: _codaEntity{
						ID:   "c-Zs7oQbj-_J",
						Name: "Start date",
					},
					GeneralSD: _codaEntity{
						ID:   "c-EMbqnhOkkr",
						Name: "General SD",
					},
					PersonnelSD: _codaEntity{
						ID:   "c-6L5oS7LZ4a",
						Name: "Personnel SD",
					},
					FinanceSD: _codaEntity{
						ID:   "c-NeQmHu-raB",
						Name: "Finance SD",
					},
					ACLSD: _codaEntity{
						ID:   "c-aKn3TiCF_A",
						Name: "ACL SD",
					},
					ProbationEnd: _codaEntity{
						ID:   "c-35cUDxDnAo",
						Name: "Probation end",
					},
					AnnualLeaveFrom: _codaEntity{
						ID:   "c-uzP6UrJs3a",
						Name: "Annual leave from",
					},
					Location: _codaEntity{
						ID:   "c-WcmDQXPChx",
						Name: "Location",
					},
					LegalEntity: _codaEntity{
						ID:   "c-smk4a68He5",
						Name: "Legal entity",
					},
					LegalForm: _codaEntity{
						ID:   "c-w8WTgIbQ7H",
						Name: "Legal form",
					},
					Position: _codaEntity{
						ID:   "c-4nDWkuySVp",
						Name: "Position",
					},
					Bank: _codaEntity{
						ID:   "c-Hx-nCcWt2v",
						Name: "Bank",
					},
					EnglishFullName: _codaEntity{
						ID:   "c-TAlfzDcFzQ",
						Name: "English full name",
					},
					InvoiceFieldsErrors: _codaEntity{
						ID:   "c-EDwQy_l7fS",
						Name: "Invoice fields errors",
					},
					PaymentFieldsErrors: _codaEntity{
						ID:   "c-UcN9cnw7y0",
						Name: "Payment fields errors",
					},
					ContractNumber: _codaEntity{
						ID:   "c-XX2OkCkOSR",
						Name: "Contract Number",
					},
					ContractDate: _codaEntity{
						ID:   "c-fDIRo1JHHX",
						Name: "Contract Date",
					},
					BankDetails: _codaEntity{
						ID:   "c-wSSi6yT6d4",
						Name: "Bank details",
					},
					BankRequisites: _codaEntity{
						ID:   "c-OlCoWd7n4S",
						Name: "Bank requisites",
					},
					BankTotalFees: _codaEntity{
						ID:   "c-156kDMsJzb",
						Name: "Bank:Total Fees",
					},
					PaymentNotes: _codaEntity{
						ID:   "c-PkLgxjZRQL",
						Name: "Payment notes",
					},
					BankCC: _codaEntity{
						ID:   "c-HbuzOpvaPf",
						Name: "Bank:CC",
					},
					BankService: _codaEntity{
						ID:   "c-214wDdnGtE",
						Name: "Bank:Service",
					},
				},
			},
			WorkingEmployeesNames: _workingEmployeesNamesTable{
				_codaEntity: _codaEntity{
					ID:   "table-dAwpX0i58t",
					Name: "Working employees names",
				},
				Cols: _workingEmployeesNamesTableColumns{
					Name: _codaEntity{
						ID:   "c-tCDt6yt4Ix",
						Name: "Name",
					},
					Position: _codaEntity{
						ID:   "c-4nDWkuySVp",
						Name: "Position",
					},
					StartDate: _codaEntity{
						ID:   "c-Zs7oQbj-_J",
						Name: "Start date",
					},
					Location: _codaEntity{
						ID:   "c-WcmDQXPChx",
						Name: "Location",
					},
				},
			},
			AllSalaries: _allSalariesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-e85pRRoRmQ",
					Name: "All salaries",
				},
				Cols: _allSalariesTableColumns{
					Employee: _codaEntity{
						ID:   "c-jbULjV0Fqy",
						Name: "Employee",
					},
					PeriodFrom: _codaEntity{
						ID:   "c-TA0eJ7Sp1o",
						Name: "Period from",
					},
					PeriodTo: _codaEntity{
						ID:   "c-g3fwvRIwO7",
						Name: "Period to",
					},
					Salary: _codaEntity{
						ID:   "c-3zcX3st-Yw",
						Name: "Salary",
					},
					PeriodFromReal: _codaEntity{
						ID:   "c--eZiTvXerm",
						Name: "Period from real",
					},
					PeriodToReal: _codaEntity{
						ID:   "c-V6JzCUTaB1",
						Name: "Period to real",
					},
					ModifiedBy: _codaEntity{
						ID:   "c-bs4r9N-1Pv",
						Name: "Modified by",
					},
					ModifiedOn: _codaEntity{
						ID:   "c-bdEdjvzyBx",
						Name: "Modified on",
					},
					CreatedOn: _codaEntity{
						ID:   "c-yynSxjRXAT",
						Name: "Created on",
					},
					Currency: _codaEntity{
						ID:   "c-XI2oSaBW6W",
						Name: "Currency",
					},
					IsRUB: _codaEntity{
						ID:   "c-D83OkDvxKl",
						Name: "IsRUB",
					},
					IsEUR: _codaEntity{
						ID:   "c-DHm4CICiqD",
						Name: "IsEUR",
					},
				},
			},
			InvoicesOverview: _invoicesOverviewTable{
				_codaEntity: _codaEntity{
					ID:   "table-_k4m2krdyp",
					Name: "Invoices overview",
				},
				Cols: _invoicesOverviewTableColumns{
					Employee: _codaEntity{
						ID:   "c-bbHUhqlbfN",
						Name: "Employee",
					},
					Correction: _codaEntity{
						ID:   "c-cjYMwmv-1m",
						Name: "Correction",
					},
					Template: _codaEntity{
						ID:   "c-z0Irm78xhQ",
						Name: "Template",
					},
					Month: _codaEntity{
						ID:   "c-wR0IONcxGH",
						Name: "Month",
					},
					ToPay: _codaEntity{
						ID:   "c-DRPGK3XTmD",
						Name: "To Pay",
					},
					Reconciled: _codaEntity{
						ID:   "c--A1RtCy-NA",
						Name: "Reconciled",
					},
					Alerts: _codaEntity{
						ID:   "c-jkrvvE_SNm",
						Name: "Alerts",
					},
					RUBSubtotal: _codaEntity{
						ID:   "c-WLMdyl4WqI",
						Name: "RUB subtotal",
					},
					BaseSalaryRUB: _codaEntity{
						ID:   "c-wqNhZf9EQY",
						Name: "Base Salary, RUB",
					},
					BaseSalaryEUR: _codaEntity{
						ID:   "c-B0nU6ZI_Z4",
						Name: "Base Salary, EUR",
					},
					BankFees: _codaEntity{
						ID:   "c-sRGR6jYC7g",
						Name: "Bank fees",
					},
					RateErrorPM: _codaEntity{
						ID:   "c-_9tuuG4RIN",
						Name: "Rate error PM",
					},
					CorrectionRUB: _codaEntity{
						ID:   "c-jNcl4nZe_h",
						Name: "Correction, RUB",
					},
					CorrectionRefs: _codaEntity{
						ID:   "c-tpeCMU21_I",
						Name: "Correction Refs",
					},
					PatentRUB: _codaEntity{
						ID:   "c-qA_pPM9kuZ",
						Name: "Patent, RUB",
					},
					TaxesRUB: _codaEntity{
						ID:   "c-ug709va8_K",
						Name: "Taxes, RUB",
					},
					EURRUBExpected: _codaEntity{
						ID:   "c-tvtGu9juVL",
						Name: "EURRUB expected",
					},
					EURRUBActual: _codaEntity{
						ID:   "c-kLIyv9EvyH",
						Name: "EURRUB actual",
					},
					RequestedSubtotalEUR: _codaEntity{
						ID:   "c-9rnJJZ6gA7",
						Name: "Requested subtotal, EUR",
					},
					RoundingPrevMonEUR: _codaEntity{
						ID:   "c-hLrmDsk89g",
						Name: "Rounding PrevMon, EUR",
					},
					Rounding: _codaEntity{
						ID:   "c-Tri-EGUP_n",
						Name: "Rounding",
					},
					RequestedTotalEUR: _codaEntity{
						ID:   "c-bJpHVxywXD",
						Name: "Requested total, EUR",
					},
					RateErrorRUB: _codaEntity{
						ID:   "c-SsHRhKa_uC",
						Name: "Rate Error, RUB",
					},
					CostOfDay: _codaEntity{
						ID:   "c-yJnq9stsgi",
						Name: "Cost of day",
					},
					BankDetails: _codaEntity{
						ID:   "c-2TxWM7vHls",
						Name: "Bank details",
					},
				},
			},
			Taxes: _taxesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-057HtXfvYH",
					Name: "Taxes",
				},
				Cols: _taxesTableColumns{
					Invoice: _codaEntity{
						ID:   "c-sYkd0N-9Ef",
						Name: "Invoice",
					},
					OpeningDateIP: _codaEntity{
						ID:   "c-Q83cwaB-vf",
						Name: "Opening date IP",
					},
					PeriodStart: _codaEntity{
						ID:   "c-UgsMCGn1oX",
						Name: "Period Start",
					},
					PeriodEnd: _codaEntity{
						ID:   "c-7Bt0dg_odm",
						Name: "Period End",
					},
					AmountIPDays: _codaEntity{
						ID:   "c-wraB-EUIPu",
						Name: "Amount IP days",
					},
					SocialInsuranceToPay: _codaEntity{
						ID:   "c-hRpESBxMnx",
						Name: "Social Insurance - to pay",
					},
					PensionFundFixedToPay: _codaEntity{
						ID:   "c-5Gm9BIf7sa",
						Name: "Pension fund fixed - to pay",
					},
					PensionFundPercent: _codaEntity{
						ID:   "c-nO-EnVUZSb",
						Name: "Pension fund percent",
					},
					Amount: _codaEntity{
						ID:   "c-FlQQoKxoau",
						Name: "Amount",
					},
					SocialInsuranceTotal: _codaEntity{
						ID:   "c-6OpLGjaPb6",
						Name: "Social Insurance - total",
					},
					PensionFundFixedTotal: _codaEntity{
						ID:   "c-D3jKQu0yHQ",
						Name: "Pension Fund fixed - total",
					},
					Year: _codaEntity{
						ID:   "c-_hn4kyEKo7",
						Name: "Year",
					},
				},
			},
			PatentPayments: _patentPaymentsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-_IJllxLQCt",
					Name: "Patent Payments",
				},
				Cols: _patentPaymentsTableColumns{
					Invoice: _codaEntity{
						ID:   "c-sYkd0N-9Ef",
						Name: "Invoice",
					},
					PeriodCost: _codaEntity{
						ID:   "c-FlQQoKxoau",
						Name: "Period cost",
					},
					Period: _codaEntity{
						ID:   "c-gtV-Qz9osQ",
						Name: "Period",
					},
					PeriodCostManual: _codaEntity{
						ID:   "c-iR7HKVaAvK",
						Name: "Period cost manual",
					},
					EmployeePatentRef: _codaEntity{
						ID:   "c-oRFcL19X2d",
						Name: "Employee patent ref",
					},
				},
			},
			CorrectionCategory: _correctionCategoryTable{
				_codaEntity: _codaEntity{
					ID:   "grid-ZQydbIU73l",
					Name: "Correction Category",
				},
				Cols: _correctionCategoryTableColumns{
					Category: _codaEntity{
						ID:   "c-2AJbX_XNUx",
						Name: "Category",
					},
					Comment: _codaEntity{
						ID:   "c-y50vUFEVml",
						Name: "Comment",
					},
				},
			},
			TaxesOverview: _taxesOverviewTable{
				_codaEntity: _codaEntity{
					ID:   "table-X3zgg4jJ8-",
					Name: "Taxes overview",
				},
				Cols: _taxesOverviewTableColumns{
					Invoice: _codaEntity{
						ID:   "c-sYkd0N-9Ef",
						Name: "Invoice",
					},
					Amount: _codaEntity{
						ID:   "c-FlQQoKxoau",
						Name: "Amount",
					},
					OpeningDateIP: _codaEntity{
						ID:   "c-Q83cwaB-vf",
						Name: "Opening date IP",
					},
					PeriodStart: _codaEntity{
						ID:   "c-UgsMCGn1oX",
						Name: "Period Start",
					},
					PeriodEnd: _codaEntity{
						ID:   "c-7Bt0dg_odm",
						Name: "Period End",
					},
					AmountIPDays: _codaEntity{
						ID:   "c-wraB-EUIPu",
						Name: "Amount IP days",
					},
					SocialInsuranceToPay: _codaEntity{
						ID:   "c-hRpESBxMnx",
						Name: "Social Insurance - to pay",
					},
					PensionFundFixedToPay: _codaEntity{
						ID:   "c-5Gm9BIf7sa",
						Name: "Pension fund fixed - to pay",
					},
					PensionFundPercent: _codaEntity{
						ID:   "c-nO-EnVUZSb",
						Name: "Pension fund percent",
					},
				},
			},
			PatentsOverview: _patentsOverviewTable{
				_codaEntity: _codaEntity{
					ID:   "table-n2aIqJxO80",
					Name: "Patents overview",
				},
				Cols: _patentsOverviewTableColumns{
					Invoice: _codaEntity{
						ID:   "c-sYkd0N-9Ef",
						Name: "Invoice",
					},
					PeriodCost: _codaEntity{
						ID:   "c-FlQQoKxoau",
						Name: "Period cost",
					},
					Period: _codaEntity{
						ID:   "c-gtV-Qz9osQ",
						Name: "Period",
					},
				},
			},
			LegalEntity: _legalEntityTable{
				_codaEntity: _codaEntity{
					ID:   "grid--dTms1XC6V",
					Name: "Legal entity",
				},
				Cols: _legalEntityTableColumns{
					EntityName: _codaEntity{
						ID:   "c-oDFd1kjV93",
						Name: "Entity name",
					},
					Requisites: _codaEntity{
						ID:   "c-5_sQbo5eww",
						Name: "Requisites",
					},
					Taxable: _codaEntity{
						ID:   "c-zCN8KuAI7k",
						Name: "Taxable",
					},
					OfficialName: _codaEntity{
						ID:   "c-eQsvF9_gTl",
						Name: "Official name",
					},
					AccountNumber: _codaEntity{
						ID:   "c-TJFRqfWIGm",
						Name: "Account number",
					},
				},
			},
			SalariesReport: _salariesReportTable{
				_codaEntity: _codaEntity{
					ID:   "table-R_NoGmR2fn",
					Name: "Salaries report",
				},
				Cols: _salariesReportTableColumns{
					Name: _codaEntity{
						ID:   "c-tCDt6yt4Ix",
						Name: "Name",
					},
					StartDate: _codaEntity{
						ID:   "c-Zs7oQbj-_J",
						Name: "Start date",
					},
					ProbationEnd: _codaEntity{
						ID:   "c-35cUDxDnAo",
						Name: "Probation end",
					},
					CurrentSalaryNet: _codaEntity{
						ID:   "c-Is6fozVp1a",
						Name: "Current Salary Net",
					},
					Position: _codaEntity{
						ID:   "c-4nDWkuySVp",
						Name: "Position",
					},
				},
			},
			TaxYears: _taxYearsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-LkqRfNkkB1",
					Name: "Tax years",
				},
				Cols: _taxYearsTableColumns{
					Year: _codaEntity{
						ID:   "c-EUuXlrh2Gj",
						Name: "Year",
					},
					PensionFundFixedPart: _codaEntity{
						ID:   "c-M2knIRBSqf",
						Name: "Pension fund fixed part",
					},
					SocialInsurance: _codaEntity{
						ID:   "c-eXc5b__5KI",
						Name: "Social Insurance",
					},
					IsCurrent: _codaEntity{
						ID:   "c--uAVz6AZiI",
						Name: "Is Current",
					},
					PensionFundPercentPart: _codaEntity{
						ID:   "c-o6tcH5VWM9",
						Name: "Pension fund percent part",
					},
					CalendarDaysInYear: _codaEntity{
						ID:   "c-xwHL0_a7ec",
						Name: "Calendar days in year",
					},
				},
			},
			EmployeePatents: _employeePatentsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-EdCa9AeMSn",
					Name: "Employee Patents",
				},
				Cols: _employeePatentsTableColumns{
					Employee: _codaEntity{
						ID:   "c-_hIvEH4-wm",
						Name: "Employee",
					},
					Year: _codaEntity{
						ID:   "c-6NwC6fSqJ8",
						Name: "Year",
					},
					Cost: _codaEntity{
						ID:   "c-GJiM8dKB7D",
						Name: "Cost",
					},
					Scan: _codaEntity{
						ID:   "c-bSYBkB-1xD",
						Name: "Scan",
					},
					Date1: _codaEntity{
						ID:   "c-GMiD_fuAfh",
						Name: "Date 1",
					},
					Date1Payment: _codaEntity{
						ID:   "c-bURhTzifbl",
						Name: "Date 1 Payment",
					},
					Date2: _codaEntity{
						ID:   "c-xE6fWWIcul",
						Name: "Date 2",
					},
					Date2Payment: _codaEntity{
						ID:   "c-a7_3Kkjsfx",
						Name: "Date 2 Payment",
					},
					Display: _codaEntity{
						ID:   "c-Elg3MNoRYd",
						Name: "Display",
					},
				},
			},
			FullPayrollReport: _fullPayrollReportTable{
				_codaEntity: _codaEntity{
					ID:   "grid-qJKfU6FFmW",
					Name: "Full payroll report",
				},
				Cols: _fullPayrollReportTableColumns{
					SalaryNet: _codaEntity{
						ID:   "c-XQPSD9vAg8",
						Name: "Salary Net",
					},
					InvoiceRef: _codaEntity{
						ID:   "c-h-RzL2yXWh",
						Name: "Invoice Ref",
					},
					OvertimesRefs: _codaEntity{
						ID:   "c-NYgI7ALlUW",
						Name: "Overtimes Refs",
					},
					OvertimesRUB: _codaEntity{
						ID:   "c-t2MioXrUCp",
						Name: "Overtimes, RUB",
					},
					BonusesRefs: _codaEntity{
						ID:   "c-EEUaWatwxL",
						Name: "Bonuses Refs",
					},
					BonusesRUB: _codaEntity{
						ID:   "c-xt7ILGpLk4",
						Name: "Bonuses, RUB",
					},
					PublicHolidaysRefs: _codaEntity{
						ID:   "c--dFWXSXlSs",
						Name: "Public Holidays Refs",
					},
					PublicHolidaysRUB: _codaEntity{
						ID:   "c-bW1uokxXyQ",
						Name: "Public holidays, RUB",
					},
					RestCorrectionRefs: _codaEntity{
						ID:   "c-Lh_T0xuy8D",
						Name: "Rest correction Refs",
					},
					RestCorrectionsRUB: _codaEntity{
						ID:   "c-fwWdBxz5_G",
						Name: "Rest corrections, RUB",
					},
					AnnualContribsRUB: _codaEntity{
						ID:   "c-RtpNgfxxJH",
						Name: "Annual Contribs, RUB",
					},
					StartDayRefs: _codaEntity{
						ID:   "c-7n5ATW-ITP",
						Name: "Start day Refs",
					},
					StartDayRUB: _codaEntity{
						ID:   "c-mSFM6hiUeA",
						Name: "Start day, RUB",
					},
					SalaryChangesRefs: _codaEntity{
						ID:   "c-y0SWuMDsBH",
						Name: "Salary changes Refs",
					},
					SalaryChangesRUB: _codaEntity{
						ID:   "c-ZY9kZCN7ap",
						Name: "Salary changes, RUB",
					},
					LoansRefs: _codaEntity{
						ID:   "c-X4bSWDyr-G",
						Name: "Loans Refs",
					},
					LoansRUB: _codaEntity{
						ID:   "c-0yn_hfYsmz",
						Name: "Loans, RUB",
					},
					RUBServiceFees: _codaEntity{
						ID:   "c-ixLJinrV3u",
						Name: "RUB Service fees",
					},
					EURRounding: _codaEntity{
						ID:   "c-vugGgAGBr1",
						Name: "EUR Rounding",
					},
					AnnualUnpaidLeavesRUB: _codaEntity{
						ID:   "c-Uo4x06gCTd",
						Name: "Annual/unpaid leaves, RUB",
					},
					MonthlyContribsEUR: _codaEntity{
						ID:   "c-wZKc8T-kxa",
						Name: "Monthly Contribs, EUR",
					},
					QuitCompensationsRefs: _codaEntity{
						ID:   "c-mGpqcoJnCA",
						Name: "Quit compensations Refs",
					},
					QuitCompensationsRUB: _codaEntity{
						ID:   "c-vF3Bj7iDMb",
						Name: "Quit compensations, RUB",
					},
					PaidAnnualLeavesRefs: _codaEntity{
						ID:   "c-PsqZc5nM7T",
						Name: "Paid annual leaves Refs",
					},
					Location: _codaEntity{
						ID:   "c-R61qZWF4OL",
						Name: "Location",
					},
				},
			},
			InvoicableEmployees: _invoicableEmployeesTable{
				_codaEntity: _codaEntity{
					ID:   "table-17g3LYb54N",
					Name: "Invoicable Employees",
				},
				Cols: _invoicableEmployeesTableColumns{
					Name: _codaEntity{
						ID:   "c-tCDt6yt4Ix",
						Name: "Name",
					},
					LegalForm: _codaEntity{
						ID:   "c-w8WTgIbQ7H",
						Name: "Legal form",
					},
					InvoiceAdd: _codaEntity{
						ID:   "c-F8kuFeh0JE",
						Name: "Invoice Add",
					},
					PendingInvoice: _codaEntity{
						ID:   "c-Saw2_PL8gv",
						Name: "Pending Invoice",
					},
					InvoiceFieldsErrors: _codaEntity{
						ID:   "c-EDwQy_l7fS",
						Name: "Invoice fields errors",
					},
					PaymentFieldsErrors: _codaEntity{
						ID:   "c-UcN9cnw7y0",
						Name: "Payment fields errors",
					},
				},
			},
			EmptyActualRates: _emptyActualRatesTable{
				_codaEntity: _codaEntity{
					ID:   "table-7A_oBrrTkt",
					Name: "Empty actual rates",
				},
				Cols: _emptyActualRatesTableColumns{
					Month: _codaEntity{
						ID:   "c-wR0IONcxGH",
						Name: "Month",
					},
					Employee: _codaEntity{
						ID:   "c-bbHUhqlbfN",
						Name: "Employee",
					},
					GeneralSD: _codaEntity{
						ID:   "c-Gkt52a0fjj",
						Name: "General SD",
					},
					EURRUBExpected: _codaEntity{
						ID:   "c-tvtGu9juVL",
						Name: "EURRUB expected",
					},
					EURRUBActual: _codaEntity{
						ID:   "c-kLIyv9EvyH",
						Name: "EURRUB actual",
					},
				},
			},
			CorrectionsByType: _correctionsByTypeTable{
				_codaEntity: _codaEntity{
					ID:   "table-03U7R1uke5",
					Name: "Corrections by type",
				},
				Cols: _correctionsByTypeTableColumns{
					Month: _codaEntity{
						ID:   "c-nK3Ad6VblR",
						Name: "Month",
					},
					PaymentInvoice: _codaEntity{
						ID:   "c-7SU0iOBY9J",
						Name: "Payment Invoice",
					},
					Category: _codaEntity{
						ID:   "c-zDY58PF0P6",
						Name: "Category",
					},
					Comment: _codaEntity{
						ID:   "c--_r48PQnSn",
						Name: "Comment",
					},
					TotalCorrectionRUB: _codaEntity{
						ID:   "c-0arkfr4qXv",
						Name: "Total Correction, RUB",
					},
					AbsoluteCorrectionRUB: _codaEntity{
						ID:   "c-P2x5IJuMXN",
						Name: "Absolute Correction, RUB",
					},
					AbsoluteCorrectionEUR: _codaEntity{
						ID:   "c-pRmEece9pf",
						Name: "Absolute Correction,EUR",
					},
					PercentCorrectionPercent: _codaEntity{
						ID:   "c-Wh_7DYRQYP",
						Name: "Percent Correction, %",
					},
					PerDayType: _codaEntity{
						ID:   "c-3Ivn-M1j7-",
						Name: "Per Day Type",
					},
					PerDayCoefficient: _codaEntity{
						ID:   "c-pz6W2IRzFR",
						Name: "PerDay Coefficient",
					},
					PerDayCalculationInvoice: _codaEntity{
						ID:   "c-bK4qXZUCqs",
						Name: "PerDay calculation invoice",
					},
					NumberOfDays: _codaEntity{
						ID:   "c-gDOyigH1cm",
						Name: "Number of days",
					},
					PerDay: _codaEntity{
						ID:   "c-Y2E1Vwe2_-",
						Name: "Per Day",
					},
				},
			},
			WorkingSalaries: _workingSalariesTable{
				_codaEntity: _codaEntity{
					ID:   "table-m2hJO5ZyIz",
					Name: "Working Salaries",
				},
				Cols: _workingSalariesTableColumns{
					Employee: _codaEntity{
						ID:   "c-jbULjV0Fqy",
						Name: "Employee",
					},
					PeriodFrom: _codaEntity{
						ID:   "c-TA0eJ7Sp1o",
						Name: "Period from",
					},
					PeriodTo: _codaEntity{
						ID:   "c-g3fwvRIwO7",
						Name: "Period to",
					},
					Salary: _codaEntity{
						ID:   "c-3zcX3st-Yw",
						Name: "Salary",
					},
					Currency: _codaEntity{
						ID:   "c-XI2oSaBW6W",
						Name: "Currency",
					},
					IsRUB: _codaEntity{
						ID:   "c-D83OkDvxKl",
						Name: "IsRUB",
					},
					IsEUR: _codaEntity{
						ID:   "c-DHm4CICiqD",
						Name: "IsEUR",
					},
					ModifiedBy: _codaEntity{
						ID:   "c-bs4r9N-1Pv",
						Name: "Modified by",
					},
					CreatedOn: _codaEntity{
						ID:   "c-yynSxjRXAT",
						Name: "Created on",
					},
					ModifiedOn: _codaEntity{
						ID:   "c-bdEdjvzyBx",
						Name: "Modified on",
					},
				},
			},
			PayrollReportToAdd: _payrollReportToAddTable{
				_codaEntity: _codaEntity{
					ID:   "table-H3qyUiWeb2",
					Name: "Payroll report to add",
				},
				Cols: _payrollReportToAddTableColumns{
					Month: _codaEntity{
						ID:   "c-wR0IONcxGH",
						Name: "Month",
					},
					Employee: _codaEntity{
						ID:   "c-bbHUhqlbfN",
						Name: "Employee",
					},
					AddToPayrollReport: _codaEntity{
						ID:   "c-FqZuQ1JQN0",
						Name: "Add to payroll report",
					},
				},
			},
			BankFees: _bankFeesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-FqLroBl9Kk",
					Name: "Bank fees",
				},
				Cols: _bankFeesTableColumns{
					Name: _codaEntity{
						ID:   "c-lVHmajh8kg",
						Name: "Name",
					},
					DefaultCCFee: _codaEntity{
						ID:   "c-VuojIAZubU",
						Name: "Default CC Fee",
					},
					DefaultServiceFee: _codaEntity{
						ID:   "c-RRIALNhlhh",
						Name: "Default Service Fee",
					},
				},
			},
			PHCorrections: _pHCorrectionsTable{
				_codaEntity: _codaEntity{
					ID:   "table-yaoOyDdzWV",
					Name: "PH corrections",
				},
				Cols: _pHCorrectionsTableColumns{
					Month: _codaEntity{
						ID:   "c-nK3Ad6VblR",
						Name: "Month",
					},
					Comment: _codaEntity{
						ID:   "c--_r48PQnSn",
						Name: "Comment",
					},
					TotalCorrectionRUB: _codaEntity{
						ID:   "c-0arkfr4qXv",
						Name: "Total Correction, RUB",
					},
					EURRUBExpected: _codaEntity{
						ID:   "c-8LD34cnmCh",
						Name: "EURRUB expected",
					},
					TotalCorrectionEUR: _codaEntity{
						ID:   "c-U8Bjm-cpmW",
						Name: "Total Correction, EUR",
					},
				},
			},
			PHMonthlyReport: _pHMonthlyReportTable{
				_codaEntity: _codaEntity{
					ID:   "grid-uoaCZGuzih",
					Name: "PH Monthly Report",
				},
				Cols: _pHMonthlyReportTableColumns{
					Month: _codaEntity{
						ID:   "c-ykIozjFfuS",
						Name: "Month",
					},
					TotalCorrectionRUB: _codaEntity{
						ID:   "c-hqPXSYAutx",
						Name: "Total Correction, RUB",
					},
					TotalCorrectionEUR: _codaEntity{
						ID:   "c-lQEjHtOlUm",
						Name: "Total Correction, EUR",
					},
				},
			},
			Location: _locationTable{
				_codaEntity: _codaEntity{
					ID:   "grid-zW24rUkABJ",
					Name: "Location",
				},
				Cols: _locationTableColumns{
					Location: _codaEntity{
						ID:   "c-TCMfV9X19C",
						Name: "Location",
					},
				},
			},
			PaymentNotesCorrectionAdding: _paymentNotesCorrectionAddingTable{
				_codaEntity: _codaEntity{
					ID:   "table-JJwP6k3c7a",
					Name: "Payment Notes Correction Adding",
				},
				Cols: _paymentNotesCorrectionAddingTableColumns{
					Month: _codaEntity{
						ID:   "c-wR0IONcxGH",
						Name: "Month",
					},
					Employee: _codaEntity{
						ID:   "c-bbHUhqlbfN",
						Name: "Employee",
					},
					Correction: _codaEntity{
						ID:   "c-cjYMwmv-1m",
						Name: "Correction",
					},
					CorrectionRefs: _codaEntity{
						ID:   "c-tpeCMU21_I",
						Name: "Correction Refs",
					},
					PaymentNotes: _codaEntity{
						ID:   "c-f4weK3W_gH",
						Name: "Payment notes",
					},
				},
			},
			CompanyRates: _companyRatesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-EKocnbzws-",
					Name: "Company rates",
				},
				Cols: _companyRatesTableColumns{
					Month: _codaEntity{
						ID:   "c-qTxlDXIGQE",
						Name: "Month",
					},
					Rate: _codaEntity{
						ID:   "c-fmUjYXxNxl",
						Name: "Rate",
					},
				},
			},
			SummaryForCurrentPeriod3: _summaryForCurrentPeriod3Table{
				_codaEntity: _codaEntity{
					ID:   "grid-ClvroIdiQE",
					Name: "Summary for current period 3",
				},
				Cols: _summaryForCurrentPeriod3TableColumns{
					Subject: _codaEntity{
						ID:   "c-c1sgQCiF-0",
						Name: "Subject",
					},
					RUB: _codaEntity{
						ID:   "c-if5o63hqNX",
						Name: "RUB",
					},
					EUR: _codaEntity{
						ID:   "c-A7oM0_7cO0",
						Name: "EUR",
					},
					Location: _codaEntity{
						ID:   "c-BpKKTPnwp2",
						Name: "Location",
					},
					LocationRefs: _codaEntity{
						ID:   "c-L5lFB2c7l9",
						Name: "Location Refs",
					},
				},
			},
			CurrentPayrollV2: _currentPayrollV2Table{
				_codaEntity: _codaEntity{
					ID:   "table-eqcaSDFij2",
					Name: "Current payroll v2",
				},
				Cols: _currentPayrollV2TableColumns{
					InvoiceRef: _codaEntity{
						ID:   "c-h-RzL2yXWh",
						Name: "Invoice Ref",
					},
					Location: _codaEntity{
						ID:   "c-R61qZWF4OL",
						Name: "Location",
					},
					SalaryNet: _codaEntity{
						ID:   "c-XQPSD9vAg8",
						Name: "Salary Net",
					},
					RUBServiceFees: _codaEntity{
						ID:   "c-ixLJinrV3u",
						Name: "RUB Service fees",
					},
					EURRounding: _codaEntity{
						ID:   "c-vugGgAGBr1",
						Name: "EUR Rounding",
					},
					AnnualUnpaidLeavesRUB: _codaEntity{
						ID:   "c-Uo4x06gCTd",
						Name: "Annual/unpaid leaves, RUB",
					},
					AnnualContribsRUB: _codaEntity{
						ID:   "c-RtpNgfxxJH",
						Name: "Annual Contribs, RUB",
					},
					MonthlyContribsEUR: _codaEntity{
						ID:   "c-wZKc8T-kxa",
						Name: "Monthly Contribs, EUR",
					},
					PublicHolidaysRefs: _codaEntity{
						ID:   "c--dFWXSXlSs",
						Name: "Public Holidays Refs",
					},
					PublicHolidaysRUB: _codaEntity{
						ID:   "c-bW1uokxXyQ",
						Name: "Public holidays, RUB",
					},
					OvertimesRefs: _codaEntity{
						ID:   "c-NYgI7ALlUW",
						Name: "Overtimes Refs",
					},
					OvertimesRUB: _codaEntity{
						ID:   "c-t2MioXrUCp",
						Name: "Overtimes, RUB",
					},
					BonusesRefs: _codaEntity{
						ID:   "c-EEUaWatwxL",
						Name: "Bonuses Refs",
					},
					BonusesRUB: _codaEntity{
						ID:   "c-xt7ILGpLk4",
						Name: "Bonuses, RUB",
					},
					LoansRefs: _codaEntity{
						ID:   "c-X4bSWDyr-G",
						Name: "Loans Refs",
					},
					LoansRUB: _codaEntity{
						ID:   "c-0yn_hfYsmz",
						Name: "Loans, RUB",
					},
					SalaryChangesRefs: _codaEntity{
						ID:   "c-y0SWuMDsBH",
						Name: "Salary changes Refs",
					},
					SalaryChangesRUB: _codaEntity{
						ID:   "c-ZY9kZCN7ap",
						Name: "Salary changes, RUB",
					},
					StartDayRefs: _codaEntity{
						ID:   "c-7n5ATW-ITP",
						Name: "Start day Refs",
					},
					StartDayRUB: _codaEntity{
						ID:   "c-mSFM6hiUeA",
						Name: "Start day, RUB",
					},
					QuitCompensationsRefs: _codaEntity{
						ID:   "c-mGpqcoJnCA",
						Name: "Quit compensations Refs",
					},
					QuitCompensationsRUB: _codaEntity{
						ID:   "c-vF3Bj7iDMb",
						Name: "Quit compensations, RUB",
					},
					PaidAnnualLeavesRefs: _codaEntity{
						ID:   "c-PsqZc5nM7T",
						Name: "Paid annual leaves Refs",
					},
					RestCorrectionRefs: _codaEntity{
						ID:   "c-Lh_T0xuy8D",
						Name: "Rest correction Refs",
					},
					RestCorrectionsRUB: _codaEntity{
						ID:   "c-fwWdBxz5_G",
						Name: "Rest corrections, RUB",
					},
				},
			},
			HeadcountPerLocation: _headcountPerLocationTable{
				_codaEntity: _codaEntity{
					ID:   "grid-h6g2-D4HkR",
					Name: "Headcount per location",
				},
				Cols: _headcountPerLocationTableColumns{
					Location: _codaEntity{
						ID:   "c-ht4rbBoGAx",
						Name: "Location",
					},
					Headcount: _codaEntity{
						ID:   "c-eLHYZpUwyV",
						Name: "Headcount",
					},
				},
			},
			SuspiciousRates: _suspiciousRatesTable{
				_codaEntity: _codaEntity{
					ID:   "table-aBh0H9XkSP",
					Name: "Suspicious rates",
				},
				Cols: _suspiciousRatesTableColumns{
					Month: _codaEntity{
						ID:   "c-wR0IONcxGH",
						Name: "Month",
					},
					Employee: _codaEntity{
						ID:   "c-bbHUhqlbfN",
						Name: "Employee",
					},
					GeneralSD: _codaEntity{
						ID:   "c-Gkt52a0fjj",
						Name: "General SD",
					},
					EURRUBExpected: _codaEntity{
						ID:   "c-tvtGu9juVL",
						Name: "EURRUB expected",
					},
					EURRUBActual: _codaEntity{
						ID:   "c-kLIyv9EvyH",
						Name: "EURRUB actual",
					},
				},
			},
			WorkDays: _workDaysTable{
				_codaEntity: _codaEntity{
					ID:   "grid-YdcFoXJyHm",
					Name: "Work days",
				},
				Cols: _workDaysTableColumns{
					Month: _codaEntity{
						ID:   "c-DanUSfEigL",
						Name: "Month",
					},
					WorkDays: _codaEntity{
						ID:   "c-lGp5-iaTg_",
						Name: "Work days",
					},
					Location: _codaEntity{
						ID:   "c-9EPvcm9xp-",
						Name: "Location",
					},
				},
			},
			CorrectionsByEmployee: _correctionsByEmployeeTable{
				_codaEntity: _codaEntity{
					ID:   "table-wHw0oNhM9z",
					Name: "Corrections by employee",
				},
				Cols: _correctionsByEmployeeTableColumns{
					Month: _codaEntity{
						ID:   "c-nK3Ad6VblR",
						Name: "Month",
					},
					Category: _codaEntity{
						ID:   "c-zDY58PF0P6",
						Name: "Category",
					},
					Comment: _codaEntity{
						ID:   "c--_r48PQnSn",
						Name: "Comment",
					},
					TotalCorrectionRUB: _codaEntity{
						ID:   "c-0arkfr4qXv",
						Name: "Total Correction, RUB",
					},
					AbsoluteCorrectionRUB: _codaEntity{
						ID:   "c-P2x5IJuMXN",
						Name: "Absolute Correction, RUB",
					},
					AbsoluteCorrectionEUR: _codaEntity{
						ID:   "c-pRmEece9pf",
						Name: "Absolute Correction,EUR",
					},
					PercentCorrectionPercent: _codaEntity{
						ID:   "c-Wh_7DYRQYP",
						Name: "Percent Correction, %",
					},
					PerDayType: _codaEntity{
						ID:   "c-3Ivn-M1j7-",
						Name: "Per Day Type",
					},
					PerDayCoefficient: _codaEntity{
						ID:   "c-pz6W2IRzFR",
						Name: "PerDay Coefficient",
					},
					PerDayCalculationInvoice: _codaEntity{
						ID:   "c-bK4qXZUCqs",
						Name: "PerDay calculation invoice",
					},
					NumberOfDays: _codaEntity{
						ID:   "c-gDOyigH1cm",
						Name: "Number of days",
					},
					PerDay: _codaEntity{
						ID:   "c-Y2E1Vwe2_-",
						Name: "Per Day",
					},
					EmployeeName: _codaEntity{
						ID:   "c-OrAvunmGm4",
						Name: "Employee name",
					},
				},
			},
			EnterCompanyRatesForSelectedMonth: _enterCompanyRatesForSelectedMonthTable{
				_codaEntity: _codaEntity{
					ID:   "table-IL1Xkw9xN7",
					Name: "Enter company rates for selected month",
				},
				Cols: _enterCompanyRatesForSelectedMonthTableColumns{
					Month: _codaEntity{
						ID:   "c-qTxlDXIGQE",
						Name: "Month",
					},
					Rate: _codaEntity{
						ID:   "c-fmUjYXxNxl",
						Name: "Rate",
					},
				},
			},
			CorrectionTemplates: _correctionTemplatesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-UCDG_ZiZul",
					Name: "Correction Templates",
				},
				Cols: _correctionTemplatesTableColumns{
					Employee: _codaEntity{
						ID:   "c-i3AkRY8VR2",
						Name: "Employee",
					},
					MonthFrom: _codaEntity{
						ID:   "c-ecmeLn9tY5",
						Name: "Month from",
					},
					MonthTo: _codaEntity{
						ID:   "c-uT9WCaQDwD",
						Name: "Month to",
					},
					Category: _codaEntity{
						ID:   "c-XAPnEUIYZl",
						Name: "Category",
					},
					Comment: _codaEntity{
						ID:   "c-MA1AOZuc-X",
						Name: "Comment",
					},
					TemplateAmountRUB: _codaEntity{
						ID:   "c-JM08fLTvoQ",
						Name: "Template amount, RUB",
					},
					TemplateAmountEUR: _codaEntity{
						ID:   "c-LRs1cqydXQ",
						Name: "Template amount, EUR",
					},
				},
			},
			AllEmployeesNames: _allEmployeesNamesTable{
				_codaEntity: _codaEntity{
					ID:   "table-krFqRkjiag",
					Name: "All employees names",
				},
				Cols: _allEmployeesNamesTableColumns{
					Name: _codaEntity{
						ID:   "c-tCDt6yt4Ix",
						Name: "Name",
					},
					Position: _codaEntity{
						ID:   "c-4nDWkuySVp",
						Name: "Position",
					},
					StartDate: _codaEntity{
						ID:   "c-Zs7oQbj-_J",
						Name: "Start date",
					},
					Location: _codaEntity{
						ID:   "c-WcmDQXPChx",
						Name: "Location",
					},
				},
			},
			BankDetails: _bankDetailsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-RmSMcSeoRT",
					Name: "Bank details",
				},
				Cols: _bankDetailsTableColumns{
					Employee: _codaEntity{
						ID:   "c-G9rzkLyUA6",
						Name: "Employee",
					},
					Account: _codaEntity{
						ID:   "c-7TwwLMMH2Y",
						Name: "Account",
					},
					Address1: _codaEntity{
						ID:   "c-Z4kWEGdXUY",
						Name: "Address 1",
					},
					Address2: _codaEntity{
						ID:   "c-I051sZQZwd",
						Name: "Address 2",
					},
					BeneficiaryBank: _codaEntity{
						ID:   "c-ychArTE7uo",
						Name: "Beneficiary Bank",
					},
					MonthFrom: _codaEntity{
						ID:   "c-jJTCJEAhIj",
						Name: "Month from",
					},
					MonthTo: _codaEntity{
						ID:   "c-zWrlKJa5JL",
						Name: "Month to",
					},
					ID: _codaEntity{
						ID:   "c-6wrPfoVH6_",
						Name: "ID",
					},
					Text: _codaEntity{
						ID:   "c-vWfYwgU4Up",
						Name: "Text",
					},
				},
			},
			BeneficiaryBank: _beneficiaryBankTable{
				_codaEntity: _codaEntity{
					ID:   "grid-nm04UZgIT5",
					Name: "Beneficiary Bank",
				},
				Cols: _beneficiaryBankTableColumns{
					Name: _codaEntity{
						ID:   "c-n7C1h8-HT0",
						Name: "Name",
					},
					Address1: _codaEntity{
						ID:   "c-IXiurFEJhe",
						Name: "Address 1",
					},
					Address2: _codaEntity{
						ID:   "c-4V09Q5_ecT",
						Name: "Address 2",
					},
					Address3: _codaEntity{
						ID:   "c-eXcEw6G3-d",
						Name: "Address 3",
					},
					BeneficiarySWIFT: _codaEntity{
						ID:   "c-ZJ42UOAJJ2",
						Name: "Beneficiary SWIFT",
					},
					IntermediarySWIFT: _codaEntity{
						ID:   "c-0YB127-zuw",
						Name: "Intermediary SWIFT",
					},
				},
			},
			PayrollSchedule: _payrollScheduleTable{
				_codaEntity: _codaEntity{
					ID:   "grid-eTSlW-4Wsv",
					Name: "Payroll Schedule",
				},
				Cols: _payrollScheduleTableColumns{
					Month: _codaEntity{
						ID:   "c-KE4PnR0O-p",
						Name: "Month",
					},
					ExecutionDate: _codaEntity{
						ID:   "c-AQjpD31QKS",
						Name: "Execution date",
					},
				},
			},
		},
		Formula: _formulaSchema{
			CurrentMonth: _codaEntity{
				ID:   "f-rnJn4-MytN",
				Name: "currentMonth",
			},
			PayrollReportCurRate: _codaEntity{
				ID:   "f-29VxTdfxQ1",
				Name: "payrollReportCurRate",
			},
			InvoiceAddingRate: _codaEntity{
				ID:   "f-Szv1kBPaqg",
				Name: "invoiceAddingRate",
			},
		},
		Control: _controlSchema{
			SelectOverviewMonth: _codaEntity{
				ID:   "ctrl-pgiOKuvlsi",
				Name: "selectOverviewMonth",
			},
			SelectOverviewEmployee: _codaEntity{
				ID:   "ctrl-jOuC2qXPDn",
				Name: "selectOverviewEmployee",
			},
			WorkingSalariesCheckbox: _codaEntity{
				ID:   "ctrl-9aVAfOrSPE",
				Name: "workingSalariesCheckbox",
			},
			InvoiceAddingMonth: _codaEntity{
				ID:   "ctrl-tmRmOIk1PB",
				Name: "invoiceAddingMonth",
			},
			PayrollReportCurPeriod: _codaEntity{
				ID:   "ctrl-PdzCvb2Pk_",
				Name: "payrollReportCurPeriod",
			},
			PayrollReportAddReport: _codaEntity{
				ID:   "ctrl-UIeU83BjQI",
				Name: "payrollReportAddReport",
			},
			FillRatesShowFilled: _codaEntity{
				ID:   "ctrl-KsXj7cpo3G",
				Name: "fillRatesShowFilled",
			},
			ChkSalariesOnlyCurrent: _codaEntity{
				ID:   "ctrl-TCZgekQfuf",
				Name: "chkSalariesOnlyCurrent",
			},
		},
	}
}
