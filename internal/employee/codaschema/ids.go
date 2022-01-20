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
	AllEmployees                      _allEmployeesTable                      // All employees
	Invoice                           _invoiceTable                           // Invoice
	Months                            _monthsTable                            // Months
	Entries                           _entriesTable                           // Entries
	WorkingEmployees                  _workingEmployeesTable                  // Working Employees
	WorkingEmployeesNames             _workingEmployeesNamesTable             // Working employees names
	InvoicesOverview                  _invoicesOverviewTable                  // Invoices overview
	PatentCompensation                _patentCompensationTable                // Patent compensation
	EntryType                         _entryTypeTable                         // Entry Type
	PatentsOverview                   _patentsOverviewTable                   // Patents overview
	LegalEntity                       _legalEntityTable                       // Legal entity
	SalariesReport                    _salariesReportTable                    // Salaries report
	TaxYears                          _taxYearsTable                          // Tax years
	EmployeePatents                   _employeePatentsTable                   // Employee Patents
	FullPayrollReport                 _fullPayrollReportTable                 // Full payroll report
	EmptyActualRates                  _emptyActualRatesTable                  // Empty actual rates
	CorrectionsByType                 _correctionsByTypeTable                 // Corrections by type
	PayrollReportToAdd                _payrollReportToAddTable                // Payroll report to add
	BankTariffs                       _bankTariffsTable                       // Bank tariffs
	PHCorrections                     _pHCorrectionsTable                     // PH corrections
	PHMonthlyReport                   _pHMonthlyReportTable                   // PH Monthly Report
	Location                          _locationTable                          // Location
	CompanyRates                      _companyRatesTable                      // Company rates
	SummaryForCurrentPeriod3          _summaryForCurrentPeriod3Table          // Summary for current period 3
	CurrentPayrollV2                  _currentPayrollV2Table                  // Current payroll v2
	HeadcountPerLocation              _headcountPerLocationTable              // Headcount per location
	SuspiciousRates                   _suspiciousRatesTable                   // Suspicious rates
	WorkDays                          _workDaysTable                          // Work days
	CorrectionsByEmployee             _correctionsByEmployeeTable             // Corrections by employee
	EnterCompanyRatesForSelectedMonth _enterCompanyRatesForSelectedMonthTable // Enter company rates for selected month
	TemplateEntries                   _templateEntriesTable                   // Template Entries
	AllEmployeesNames                 _allEmployeesNamesTable                 // All employees names
	BankDetails                       _bankDetailsTable                       // Bank details
	BeneficiaryBank                   _beneficiaryBankTable                   // Beneficiary Bank
	PayrollSchedule                   _payrollScheduleTable                   // Payroll Schedule
	PensionFundFixed                  _pensionFundFixedTable                  // Pension Fund fixed
	SocialInsurance                   _socialInsuranceTable                   // Social Insurance
	PensionFundPercent                _pensionFundPercentTable                // Pension fund percent
	PerDayCalculations                _perDayCalculationsTable                // Per Day Calculations
	PerDayPolicies                    _perDayPoliciesTable                    // Per Day Policies
	Salaries                          _salariesTable                          // Salaries
	PayableEmployees                  _payableEmployeesTable                  // Payable employees
	QuickManualEntry                  _quickManualEntryTable                  // Quick Manual Entry
}
type _formulaSchema struct {
	CurrentMonth                  _codaEntity // currentMonth
	PayrollReportCurRate          _codaEntity // payrollReportCurRate
	InvoiceAddingRate             _codaEntity // invoiceAddingRate
	NewInvoiceValidation          _codaEntity // newInvoiceValidation
	NewInvoiceValidationIcon      _codaEntity // newInvoiceValidationIcon
	NewInvoiceRate                _codaEntity // newInvoiceRate
	NewInvoicePreviousRatesFilled _codaEntity // newInvoicePreviousRatesFilled
}
type _controlSchema struct {
	SelectOverviewMonth    _codaEntity // selectOverviewMonth
	SelectOverviewEmployee _codaEntity // selectOverviewEmployee
	InvoiceAddingMonth     _codaEntity // invoiceAddingMonth
	PayrollReportCurPeriod _codaEntity // payrollReportCurPeriod
	PayrollReportAddReport _codaEntity // payrollReportAddReport
	FillRatesShowFilled    _codaEntity // fillRatesShowFilled
	NewInvoiceBtn          _codaEntity // newInvoiceBtn
	NewInvoiceMonth        _codaEntity // newInvoiceMonth
}

// Table All employees
type _allEmployeesTable struct {
	_codaEntity
	Cols _allEmployeesTableColumns
}
type _allEmployeesTableColumns struct {
	Name                 _codaEntity // Name
	StartDate            _codaEntity // Start date
	EndDate              _codaEntity // End date
	ContractHourRate     _codaEntity // Contract hour rate
	OpeningDateIP        _codaEntity // Opening date IP
	WorkingNow           _codaEntity // Working now?
	EnglishFullName      _codaEntity // English full name
	LegalEntity          _codaEntity // Legal entity
	Location             _codaEntity // Location
	FinanceSD            _codaEntity // Finance SD
	GeneralSD            _codaEntity // General SD
	ContractNumber       _codaEntity // Contract Number
	ContractDate         _codaEntity // Contract Date
	LegalForm            _codaEntity // Legal form
	BankTariff           _codaEntity // Bank tariff
	Rounding             _codaEntity // Rounding
	BankDetails          _codaEntity // Bank details
	SpecialPaymentPolicy _codaEntity // Special payment policy
}

// Table Invoice
type _invoiceTable struct {
	_codaEntity
	Cols _invoiceTableColumns
}
type _invoiceTableColumns struct {
	ID                  _codaEntity // ID
	InvoiceHash         _codaEntity // Invoice #
	Month               _codaEntity // Month
	Employee            _codaEntity // Employee
	PreviousInvoice     _codaEntity // Previous invoice
	EURRUBExpected      _codaEntity // EURRUB expected
	EURSubtotal         _codaEntity // EUR Subtotal
	EURRounding         _codaEntity // EUR Rounding
	EUREntries          _codaEntity // EUR Entries
	Hours               _codaEntity // Hours
	EURRUBActual        _codaEntity // EURRUB actual
	RUBActual           _codaEntity // RUB Actual
	RUBRateError        _codaEntity // RUB Rate Error
	InvoiceEntries      _codaEntity // Invoice Entries
	RUBEntries          _codaEntity // RUB Entries
	PaymentChecksPassed _codaEntity // Payment Checks Passed
	TemplatesRefs       _codaEntity // Templates Refs
	RecipientDetails    _codaEntity // Recipient details
	EURTotal            _codaEntity // EUR Total
	ApprovedManually    _codaEntity // Approved manually
	SenderDetails       _codaEntity // Sender details
	RUBTotal            _codaEntity // RUB Total
	HourRate            _codaEntity // Hour Rate
	WorkDays            _codaEntity // Work days
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

// Table Entries
type _entriesTable struct {
	_codaEntity
	Cols _entriesTableColumns
}
type _entriesTableColumns struct {
	Comment    _codaEntity // Comment
	EURAmount  _codaEntity // EUR Amount
	RUBAmount  _codaEntity // RUB Amount
	Invoice    _codaEntity // Invoice
	Display    _codaEntity // Display
	Type       _codaEntity // Type
	ModifiedOn _codaEntity // Modified on
	CreatedOn  _codaEntity // Created on
	Origin     _codaEntity // Origin
	ModifiedBy _codaEntity // Modified by
	Sort       _codaEntity // Sort
}

// Table Working Employees
type _workingEmployeesTable struct {
	_codaEntity
	Cols _workingEmployeesTableColumns
}
type _workingEmployeesTableColumns struct {
	Name                 _codaEntity // Name
	StartDate            _codaEntity // Start date
	GeneralSD            _codaEntity // General SD
	FinanceSD            _codaEntity // Finance SD
	Location             _codaEntity // Location
	LegalEntity          _codaEntity // Legal entity
	LegalForm            _codaEntity // Legal form
	BankTariff           _codaEntity // Bank tariff
	SpecialPaymentPolicy _codaEntity // Special payment policy
	EnglishFullName      _codaEntity // English full name
	ContractNumber       _codaEntity // Contract Number
	ContractDate         _codaEntity // Contract Date
	BankDetails          _codaEntity // Bank details
}

// Table Working employees names
type _workingEmployeesNamesTable struct {
	_codaEntity
	Cols _workingEmployeesNamesTableColumns
}
type _workingEmployeesNamesTableColumns struct {
	Name      _codaEntity // Name
	StartDate _codaEntity // Start date
	Location  _codaEntity // Location
}

// Table Invoices overview
type _invoicesOverviewTable struct {
	_codaEntity
	Cols _invoicesOverviewTableColumns
}
type _invoicesOverviewTableColumns struct {
	Employee            _codaEntity // Employee
	Month               _codaEntity // Month
	PaymentChecksPassed _codaEntity // Payment Checks Passed
	RUBEntries          _codaEntity // RUB Entries
	InvoiceEntries      _codaEntity // Invoice Entries
	EURRUBExpected      _codaEntity // EURRUB expected
	EURRUBActual        _codaEntity // EURRUB actual
	EURSubtotal         _codaEntity // EUR Subtotal
	EURRounding         _codaEntity // EUR Rounding
	EUREntries          _codaEntity // EUR Entries
	RUBRateError        _codaEntity // RUB Rate Error
	RecipientDetails    _codaEntity // Recipient details
}

// Table Patent compensation
type _patentCompensationTable struct {
	_codaEntity
	Cols _patentCompensationTableColumns
}
type _patentCompensationTableColumns struct {
	PaymentInvoice     _codaEntity // Payment Invoice
	PeriodCost         _codaEntity // Period cost
	Period             _codaEntity // Period
	PeriodCostOverride _codaEntity // Period cost override
	EmployeePatentRef  _codaEntity // Employee patent ref
	Apply              _codaEntity // Apply
}

// Table Entry Type
type _entryTypeTable struct {
	_codaEntity
	Cols _entryTypeTableColumns
}
type _entryTypeTableColumns struct {
	Type      _codaEntity // Type
	Comment   _codaEntity // Comment
	Archetype _codaEntity // Archetype
	Sort      _codaEntity // Sort
}

// Table Patents overview
type _patentsOverviewTable struct {
	_codaEntity
	Cols _patentsOverviewTableColumns
}
type _patentsOverviewTableColumns struct {
	PaymentInvoice _codaEntity // Payment Invoice
	PeriodCost     _codaEntity // Period cost
	Period         _codaEntity // Period
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
	Name      _codaEntity // Name
	StartDate _codaEntity // Start date
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

// Table Empty actual rates
type _emptyActualRatesTable struct {
	_codaEntity
	Cols _emptyActualRatesTableColumns
}
type _emptyActualRatesTableColumns struct {
	Month          _codaEntity // Month
	Employee       _codaEntity // Employee
	EURRUBExpected _codaEntity // EURRUB expected
	EURRUBActual   _codaEntity // EURRUB actual
}

// Table Corrections by type
type _correctionsByTypeTable struct {
	_codaEntity
	Cols _correctionsByTypeTableColumns
}
type _correctionsByTypeTableColumns struct {
	Invoice   _codaEntity // Invoice
	Type      _codaEntity // Type
	Comment   _codaEntity // Comment
	RUBAmount _codaEntity // RUB Amount
	EURAmount _codaEntity // EUR Amount
}

// Table Payroll report to add
type _payrollReportToAddTable struct {
	_codaEntity
	Cols _payrollReportToAddTableColumns
}
type _payrollReportToAddTableColumns struct {
	Month    _codaEntity // Month
	Employee _codaEntity // Employee
}

// Table Bank tariffs
type _bankTariffsTable struct {
	_codaEntity
	Cols _bankTariffsTableColumns
}
type _bankTariffsTableColumns struct {
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
	Comment _codaEntity // Comment
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

// Table Company rates
type _companyRatesTable struct {
	_codaEntity
	Cols _companyRatesTableColumns
}
type _companyRatesTableColumns struct {
	Month  _codaEntity // Month
	EURRUB _codaEntity // EURRUB
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
	Type      _codaEntity // Type
	Comment   _codaEntity // Comment
	RUBAmount _codaEntity // RUB Amount
	EURAmount _codaEntity // EUR Amount
}

// Table Enter company rates for selected month
type _enterCompanyRatesForSelectedMonthTable struct {
	_codaEntity
	Cols _enterCompanyRatesForSelectedMonthTableColumns
}
type _enterCompanyRatesForSelectedMonthTableColumns struct {
	Month  _codaEntity // Month
	EURRUB _codaEntity // EURRUB
}

// Table Template Entries
type _templateEntriesTable struct {
	_codaEntity
	Cols _templateEntriesTableColumns
}
type _templateEntriesTableColumns struct {
	Employee   _codaEntity // Employee
	MonthFrom  _codaEntity // Month from
	MonthTo    _codaEntity // Month to
	Type       _codaEntity // Type
	Comment    _codaEntity // Comment
	RUBAmount  _codaEntity // RUB Amount
	EURAmount  _codaEntity // EUR Amount
	ModifiedBy _codaEntity // Modified by
	CreatedOn  _codaEntity // Created on
	ModifiedOn _codaEntity // Modified on
	Display    _codaEntity // Display
}

// Table All employees names
type _allEmployeesNamesTable struct {
	_codaEntity
	Cols _allEmployeesNamesTableColumns
}
type _allEmployeesNamesTableColumns struct {
	Name      _codaEntity // Name
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
	BankRequisites  _codaEntity // Bank requisites
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

// Table Pension Fund fixed
type _pensionFundFixedTable struct {
	_codaEntity
	Cols _pensionFundFixedTableColumns
}
type _pensionFundFixedTableColumns struct {
	Invoice               _codaEntity // Invoice
	OpeningDateIP         _codaEntity // Opening date IP
	PeriodStart           _codaEntity // Period Start
	PeriodEnd             _codaEntity // Period End
	AmountIPDays          _codaEntity // Amount IP days
	PensionFundFixedToPay _codaEntity // Pension fund fixed - to pay
	PensionFundFixedTotal _codaEntity // Pension Fund fixed - total
	Year                  _codaEntity // Year
}

// Table Social Insurance
type _socialInsuranceTable struct {
	_codaEntity
	Cols _socialInsuranceTableColumns
}
type _socialInsuranceTableColumns struct {
	Invoice              _codaEntity // Invoice
	OpeningDateIP        _codaEntity // Opening date IP
	PeriodStart          _codaEntity // Period Start
	PeriodEnd            _codaEntity // Period End
	AmountIPDays         _codaEntity // Amount IP days
	SocialInsuranceToPay _codaEntity // Social Insurance - to pay
	SocialInsuranceTotal _codaEntity // Social Insurance - total
	Year                 _codaEntity // Year
}

// Table Pension fund percent
type _pensionFundPercentTable struct {
	_codaEntity
	Cols _pensionFundPercentTableColumns
}
type _pensionFundPercentTableColumns struct {
	Invoice            _codaEntity // Invoice
	PensionFundPercent _codaEntity // Pension fund percent
	Year               _codaEntity // Year
}

// Table Per Day Calculations
type _perDayCalculationsTable struct {
	_codaEntity
	Cols _perDayCalculationsTableColumns
}
type _perDayCalculationsTableColumns struct {
	Type               _codaEntity // Type
	NumberOfDays       _codaEntity // Number of days
	CostOfDay          _codaEntity // Cost of day
	Total              _codaEntity // Total
	PaymentInvoice     _codaEntity // Payment invoice
	CalculationPeriod  _codaEntity // Calculation period
	CalculationInvoice _codaEntity // Calculation invoice
	Salary             _codaEntity // Salary
	Apply              _codaEntity // Apply
}

// Table Per Day Policies
type _perDayPoliciesTable struct {
	_codaEntity
	Cols _perDayPoliciesTableColumns
}
type _perDayPoliciesTableColumns struct {
	Name        _codaEntity // Name
	Coefficient _codaEntity // Coefficient
	EntryType   _codaEntity // Entry Type
}

// Table Salaries
type _salariesTable struct {
	_codaEntity
	Cols _salariesTableColumns
}
type _salariesTableColumns struct {
	Employee  _codaEntity // Employee
	MonthFrom _codaEntity // Month from
	MonthTo   _codaEntity // Month to
	Comment   _codaEntity // Comment
	RUBAmount _codaEntity // RUB Amount
	EURAmount _codaEntity // EUR Amount
}

// Table Payable employees
type _payableEmployeesTable struct {
	_codaEntity
	Cols _payableEmployeesTableColumns
}
type _payableEmployeesTableColumns struct {
	Employee           _codaEntity // Employee
	BonusQuarter       _codaEntity // Bonus Quarter
	TargetInvoice      _codaEntity // Target invoice
	AddAny             _codaEntity // Add any
	FlexBenefit        _codaEntity // Flex benefit
	SelfEmplTax        _codaEntity // Self-Empl Tax
	IE6PercentTax      _codaEntity // IE 6% Tax
	ManualEntries      _codaEntity // Manual Entries
	ExcludedEntryTypes _codaEntity // Excluded Entry Types
}

// Table Quick Manual Entry
type _quickManualEntryTable struct {
	_codaEntity
	Cols _quickManualEntryTableColumns
}
type _quickManualEntryTableColumns struct {
	TargetInvoice _codaEntity // Target invoice
	ManualEntries _codaEntity // Manual Entries
	AddAny        _codaEntity // Add any
	BonusQuarter  _codaEntity // Bonus Quarter
	FlexBenefit   _codaEntity // Flex benefit
	SelfEmplTax   _codaEntity // Self-Empl Tax
	IE6PercentTax _codaEntity // IE 6% Tax
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
					EndDate: _codaEntity{
						ID:   "c-7OoHuXqt8n",
						Name: "End date",
					},
					ContractHourRate: _codaEntity{
						ID:   "c-SItGvyE3ie",
						Name: "Contract hour rate",
					},
					OpeningDateIP: _codaEntity{
						ID:   "c-uU3-6piESs",
						Name: "Opening date IP",
					},
					WorkingNow: _codaEntity{
						ID:   "c-QfV5QzjuJP",
						Name: "Working now?",
					},
					EnglishFullName: _codaEntity{
						ID:   "c-TAlfzDcFzQ",
						Name: "English full name",
					},
					LegalEntity: _codaEntity{
						ID:   "c-smk4a68He5",
						Name: "Legal entity",
					},
					Location: _codaEntity{
						ID:   "c-WcmDQXPChx",
						Name: "Location",
					},
					FinanceSD: _codaEntity{
						ID:   "c-NeQmHu-raB",
						Name: "Finance SD",
					},
					GeneralSD: _codaEntity{
						ID:   "c-EMbqnhOkkr",
						Name: "General SD",
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
					BankTariff: _codaEntity{
						ID:   "c-Hx-nCcWt2v",
						Name: "Bank tariff",
					},
					Rounding: _codaEntity{
						ID:   "c-xvCHSSPdHb",
						Name: "Rounding",
					},
					BankDetails: _codaEntity{
						ID:   "c-wSSi6yT6d4",
						Name: "Bank details",
					},
					SpecialPaymentPolicy: _codaEntity{
						ID:   "c-FOuDAl0Fmk",
						Name: "Special payment policy",
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
					EURSubtotal: _codaEntity{
						ID:   "c-9rnJJZ6gA7",
						Name: "EUR Subtotal",
					},
					EURRounding: _codaEntity{
						ID:   "c-Tri-EGUP_n",
						Name: "EUR Rounding",
					},
					EUREntries: _codaEntity{
						ID:   "c-bJpHVxywXD",
						Name: "EUR Entries",
					},
					Hours: _codaEntity{
						ID:   "c-KtVV9if8P7",
						Name: "Hours",
					},
					EURRUBActual: _codaEntity{
						ID:   "c-kLIyv9EvyH",
						Name: "EURRUB actual",
					},
					RUBActual: _codaEntity{
						ID:   "c-AxLSgrt7e3",
						Name: "RUB Actual",
					},
					RUBRateError: _codaEntity{
						ID:   "c-SsHRhKa_uC",
						Name: "RUB Rate Error",
					},
					InvoiceEntries: _codaEntity{
						ID:   "c-tpeCMU21_I",
						Name: "Invoice Entries",
					},
					RUBEntries: _codaEntity{
						ID:   "c-jNcl4nZe_h",
						Name: "RUB Entries",
					},
					PaymentChecksPassed: _codaEntity{
						ID:   "c-DRPGK3XTmD",
						Name: "Payment Checks Passed",
					},
					TemplatesRefs: _codaEntity{
						ID:   "c-esVje33Tt-",
						Name: "Templates Refs",
					},
					RecipientDetails: _codaEntity{
						ID:   "c-2TxWM7vHls",
						Name: "Recipient details",
					},
					EURTotal: _codaEntity{
						ID:   "c-Y5fNzlzJXF",
						Name: "EUR Total",
					},
					ApprovedManually: _codaEntity{
						ID:   "c-9Zy6wPmcmR",
						Name: "Approved manually",
					},
					SenderDetails: _codaEntity{
						ID:   "c-GT8GvOJE1C",
						Name: "Sender details",
					},
					RUBTotal: _codaEntity{
						ID:   "c-WWb4V9ETtt",
						Name: "RUB Total",
					},
					HourRate: _codaEntity{
						ID:   "c-BXmgrJUyUd",
						Name: "Hour Rate",
					},
					WorkDays: _codaEntity{
						ID:   "c-ovkSZVew-L",
						Name: "Work days",
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
			Entries: _entriesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-wBmvgFgaGi",
					Name: "Entries",
				},
				Cols: _entriesTableColumns{
					Comment: _codaEntity{
						ID:   "c--_r48PQnSn",
						Name: "Comment",
					},
					EURAmount: _codaEntity{
						ID:   "c-pRmEece9pf",
						Name: "EUR Amount",
					},
					RUBAmount: _codaEntity{
						ID:   "c-P2x5IJuMXN",
						Name: "RUB Amount",
					},
					Invoice: _codaEntity{
						ID:   "c-7SU0iOBY9J",
						Name: "Invoice",
					},
					Display: _codaEntity{
						ID:   "c-FVW_9PPzZ2",
						Name: "Display",
					},
					Type: _codaEntity{
						ID:   "c-zDY58PF0P6",
						Name: "Type",
					},
					ModifiedOn: _codaEntity{
						ID:   "c-7HSE88Op0l",
						Name: "Modified on",
					},
					CreatedOn: _codaEntity{
						ID:   "c-pKkFDJQXHe",
						Name: "Created on",
					},
					Origin: _codaEntity{
						ID:   "c-aJp16HyPKx",
						Name: "Origin",
					},
					ModifiedBy: _codaEntity{
						ID:   "c-VwCxsKPCnl",
						Name: "Modified by",
					},
					Sort: _codaEntity{
						ID:   "c-7rE4nrnOXc",
						Name: "Sort",
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
					StartDate: _codaEntity{
						ID:   "c-Zs7oQbj-_J",
						Name: "Start date",
					},
					GeneralSD: _codaEntity{
						ID:   "c-EMbqnhOkkr",
						Name: "General SD",
					},
					FinanceSD: _codaEntity{
						ID:   "c-NeQmHu-raB",
						Name: "Finance SD",
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
					BankTariff: _codaEntity{
						ID:   "c-Hx-nCcWt2v",
						Name: "Bank tariff",
					},
					SpecialPaymentPolicy: _codaEntity{
						ID:   "c-FOuDAl0Fmk",
						Name: "Special payment policy",
					},
					EnglishFullName: _codaEntity{
						ID:   "c-TAlfzDcFzQ",
						Name: "English full name",
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
					Month: _codaEntity{
						ID:   "c-wR0IONcxGH",
						Name: "Month",
					},
					PaymentChecksPassed: _codaEntity{
						ID:   "c-DRPGK3XTmD",
						Name: "Payment Checks Passed",
					},
					RUBEntries: _codaEntity{
						ID:   "c-jNcl4nZe_h",
						Name: "RUB Entries",
					},
					InvoiceEntries: _codaEntity{
						ID:   "c-tpeCMU21_I",
						Name: "Invoice Entries",
					},
					EURRUBExpected: _codaEntity{
						ID:   "c-tvtGu9juVL",
						Name: "EURRUB expected",
					},
					EURRUBActual: _codaEntity{
						ID:   "c-kLIyv9EvyH",
						Name: "EURRUB actual",
					},
					EURSubtotal: _codaEntity{
						ID:   "c-9rnJJZ6gA7",
						Name: "EUR Subtotal",
					},
					EURRounding: _codaEntity{
						ID:   "c-Tri-EGUP_n",
						Name: "EUR Rounding",
					},
					EUREntries: _codaEntity{
						ID:   "c-bJpHVxywXD",
						Name: "EUR Entries",
					},
					RUBRateError: _codaEntity{
						ID:   "c-SsHRhKa_uC",
						Name: "RUB Rate Error",
					},
					RecipientDetails: _codaEntity{
						ID:   "c-2TxWM7vHls",
						Name: "Recipient details",
					},
				},
			},
			PatentCompensation: _patentCompensationTable{
				_codaEntity: _codaEntity{
					ID:   "grid-_IJllxLQCt",
					Name: "Patent compensation",
				},
				Cols: _patentCompensationTableColumns{
					PaymentInvoice: _codaEntity{
						ID:   "c-sYkd0N-9Ef",
						Name: "Payment Invoice",
					},
					PeriodCost: _codaEntity{
						ID:   "c-FlQQoKxoau",
						Name: "Period cost",
					},
					Period: _codaEntity{
						ID:   "c-gtV-Qz9osQ",
						Name: "Period",
					},
					PeriodCostOverride: _codaEntity{
						ID:   "c-iR7HKVaAvK",
						Name: "Period cost override",
					},
					EmployeePatentRef: _codaEntity{
						ID:   "c-oRFcL19X2d",
						Name: "Employee patent ref",
					},
					Apply: _codaEntity{
						ID:   "c-SVnSxLqAWW",
						Name: "Apply",
					},
				},
			},
			EntryType: _entryTypeTable{
				_codaEntity: _codaEntity{
					ID:   "grid-ZQydbIU73l",
					Name: "Entry Type",
				},
				Cols: _entryTypeTableColumns{
					Type: _codaEntity{
						ID:   "c-2AJbX_XNUx",
						Name: "Type",
					},
					Comment: _codaEntity{
						ID:   "c-y50vUFEVml",
						Name: "Comment",
					},
					Archetype: _codaEntity{
						ID:   "c-2C7Vno33Ui",
						Name: "Archetype",
					},
					Sort: _codaEntity{
						ID:   "c-YN77WP5gru",
						Name: "Sort",
					},
				},
			},
			PatentsOverview: _patentsOverviewTable{
				_codaEntity: _codaEntity{
					ID:   "table-n2aIqJxO80",
					Name: "Patents overview",
				},
				Cols: _patentsOverviewTableColumns{
					PaymentInvoice: _codaEntity{
						ID:   "c-sYkd0N-9Ef",
						Name: "Payment Invoice",
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
					Invoice: _codaEntity{
						ID:   "c-7SU0iOBY9J",
						Name: "Invoice",
					},
					Type: _codaEntity{
						ID:   "c-zDY58PF0P6",
						Name: "Type",
					},
					Comment: _codaEntity{
						ID:   "c--_r48PQnSn",
						Name: "Comment",
					},
					RUBAmount: _codaEntity{
						ID:   "c-P2x5IJuMXN",
						Name: "RUB Amount",
					},
					EURAmount: _codaEntity{
						ID:   "c-pRmEece9pf",
						Name: "EUR Amount",
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
				},
			},
			BankTariffs: _bankTariffsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-FqLroBl9Kk",
					Name: "Bank tariffs",
				},
				Cols: _bankTariffsTableColumns{
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
					Comment: _codaEntity{
						ID:   "c--_r48PQnSn",
						Name: "Comment",
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
					EURRUB: _codaEntity{
						ID:   "c-fmUjYXxNxl",
						Name: "EURRUB",
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
					Type: _codaEntity{
						ID:   "c-zDY58PF0P6",
						Name: "Type",
					},
					Comment: _codaEntity{
						ID:   "c--_r48PQnSn",
						Name: "Comment",
					},
					RUBAmount: _codaEntity{
						ID:   "c-P2x5IJuMXN",
						Name: "RUB Amount",
					},
					EURAmount: _codaEntity{
						ID:   "c-pRmEece9pf",
						Name: "EUR Amount",
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
					EURRUB: _codaEntity{
						ID:   "c-fmUjYXxNxl",
						Name: "EURRUB",
					},
				},
			},
			TemplateEntries: _templateEntriesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-UCDG_ZiZul",
					Name: "Template Entries",
				},
				Cols: _templateEntriesTableColumns{
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
					Type: _codaEntity{
						ID:   "c-XAPnEUIYZl",
						Name: "Type",
					},
					Comment: _codaEntity{
						ID:   "c-MA1AOZuc-X",
						Name: "Comment",
					},
					RUBAmount: _codaEntity{
						ID:   "c-JM08fLTvoQ",
						Name: "RUB Amount",
					},
					EURAmount: _codaEntity{
						ID:   "c-LRs1cqydXQ",
						Name: "EUR Amount",
					},
					ModifiedBy: _codaEntity{
						ID:   "c-LEF8qqAn8z",
						Name: "Modified by",
					},
					CreatedOn: _codaEntity{
						ID:   "c-09NCJ1zLwk",
						Name: "Created on",
					},
					ModifiedOn: _codaEntity{
						ID:   "c-PorXY2SBXh",
						Name: "Modified on",
					},
					Display: _codaEntity{
						ID:   "c-IF-PVJ6rtI",
						Name: "Display",
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
					BankRequisites: _codaEntity{
						ID:   "c-kvJol7ACwf",
						Name: "Bank requisites",
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
			PensionFundFixed: _pensionFundFixedTable{
				_codaEntity: _codaEntity{
					ID:   "grid-ikOK1_C1fc",
					Name: "Pension Fund fixed",
				},
				Cols: _pensionFundFixedTableColumns{
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
					PensionFundFixedToPay: _codaEntity{
						ID:   "c-5Gm9BIf7sa",
						Name: "Pension fund fixed - to pay",
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
			SocialInsurance: _socialInsuranceTable{
				_codaEntity: _codaEntity{
					ID:   "grid-xke_mOFuzO",
					Name: "Social Insurance",
				},
				Cols: _socialInsuranceTableColumns{
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
					SocialInsuranceTotal: _codaEntity{
						ID:   "c-6OpLGjaPb6",
						Name: "Social Insurance - total",
					},
					Year: _codaEntity{
						ID:   "c-_hn4kyEKo7",
						Name: "Year",
					},
				},
			},
			PensionFundPercent: _pensionFundPercentTable{
				_codaEntity: _codaEntity{
					ID:   "grid-5B0Z-WEqCn",
					Name: "Pension fund percent",
				},
				Cols: _pensionFundPercentTableColumns{
					Invoice: _codaEntity{
						ID:   "c-sYkd0N-9Ef",
						Name: "Invoice",
					},
					PensionFundPercent: _codaEntity{
						ID:   "c-nO-EnVUZSb",
						Name: "Pension fund percent",
					},
					Year: _codaEntity{
						ID:   "c-_hn4kyEKo7",
						Name: "Year",
					},
				},
			},
			PerDayCalculations: _perDayCalculationsTable{
				_codaEntity: _codaEntity{
					ID:   "grid-ik-9DIjBqz",
					Name: "Per Day Calculations",
				},
				Cols: _perDayCalculationsTableColumns{
					Type: _codaEntity{
						ID:   "c-3Ivn-M1j7-",
						Name: "Type",
					},
					NumberOfDays: _codaEntity{
						ID:   "c-gDOyigH1cm",
						Name: "Number of days",
					},
					CostOfDay: _codaEntity{
						ID:   "c-K_Iy0iERKR",
						Name: "Cost of day",
					},
					Total: _codaEntity{
						ID:   "c-Y2E1Vwe2_-",
						Name: "Total",
					},
					PaymentInvoice: _codaEntity{
						ID:   "c-7SU0iOBY9J",
						Name: "Payment invoice",
					},
					CalculationPeriod: _codaEntity{
						ID:   "c-bK4qXZUCqs",
						Name: "Calculation period",
					},
					CalculationInvoice: _codaEntity{
						ID:   "c-QXP63jNl9T",
						Name: "Calculation invoice",
					},
					Salary: _codaEntity{
						ID:   "c-jgeE5L-QmG",
						Name: "Salary",
					},
					Apply: _codaEntity{
						ID:   "c-5wgfO7-p7n",
						Name: "Apply",
					},
				},
			},
			PerDayPolicies: _perDayPoliciesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-MwSd1ImG9X",
					Name: "Per Day Policies",
				},
				Cols: _perDayPoliciesTableColumns{
					Name: _codaEntity{
						ID:   "c-SESlR7_60Q",
						Name: "Name",
					},
					Coefficient: _codaEntity{
						ID:   "c-Cg4D8wBxQx",
						Name: "Coefficient",
					},
					EntryType: _codaEntity{
						ID:   "c-MyD8pD9imA",
						Name: "Entry Type",
					},
				},
			},
			Salaries: _salariesTable{
				_codaEntity: _codaEntity{
					ID:   "table-zWwQuzxVi7",
					Name: "Salaries",
				},
				Cols: _salariesTableColumns{
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
					Comment: _codaEntity{
						ID:   "c-MA1AOZuc-X",
						Name: "Comment",
					},
					RUBAmount: _codaEntity{
						ID:   "c-JM08fLTvoQ",
						Name: "RUB Amount",
					},
					EURAmount: _codaEntity{
						ID:   "c-LRs1cqydXQ",
						Name: "EUR Amount",
					},
				},
			},
			PayableEmployees: _payableEmployeesTable{
				_codaEntity: _codaEntity{
					ID:   "grid-ULBMNv2LxC",
					Name: "Payable employees",
				},
				Cols: _payableEmployeesTableColumns{
					Employee: _codaEntity{
						ID:   "c-E3ueCNW5Td",
						Name: "Employee",
					},
					BonusQuarter: _codaEntity{
						ID:   "c-b7PNViCexb",
						Name: "Bonus Quarter",
					},
					TargetInvoice: _codaEntity{
						ID:   "c-lNa8MQX1Ze",
						Name: "Target invoice",
					},
					AddAny: _codaEntity{
						ID:   "c-5kK0bjRk3J",
						Name: "Add any",
					},
					FlexBenefit: _codaEntity{
						ID:   "c-5k0Q3x_vWp",
						Name: "Flex benefit",
					},
					SelfEmplTax: _codaEntity{
						ID:   "c-0jQDMsAjed",
						Name: "Self-Empl Tax",
					},
					IE6PercentTax: _codaEntity{
						ID:   "c-L-78eDsX4J",
						Name: "IE 6% Tax",
					},
					ManualEntries: _codaEntity{
						ID:   "c-gwL_v6jBRy",
						Name: "Manual Entries",
					},
					ExcludedEntryTypes: _codaEntity{
						ID:   "c-aMaPUsZYA1",
						Name: "Excluded Entry Types",
					},
				},
			},
			QuickManualEntry: _quickManualEntryTable{
				_codaEntity: _codaEntity{
					ID:   "table-jI1G_u-gq-",
					Name: "Quick Manual Entry",
				},
				Cols: _quickManualEntryTableColumns{
					TargetInvoice: _codaEntity{
						ID:   "c-lNa8MQX1Ze",
						Name: "Target invoice",
					},
					ManualEntries: _codaEntity{
						ID:   "c-gwL_v6jBRy",
						Name: "Manual Entries",
					},
					AddAny: _codaEntity{
						ID:   "c-5kK0bjRk3J",
						Name: "Add any",
					},
					BonusQuarter: _codaEntity{
						ID:   "c-b7PNViCexb",
						Name: "Bonus Quarter",
					},
					FlexBenefit: _codaEntity{
						ID:   "c-5k0Q3x_vWp",
						Name: "Flex benefit",
					},
					SelfEmplTax: _codaEntity{
						ID:   "c-0jQDMsAjed",
						Name: "Self-Empl Tax",
					},
					IE6PercentTax: _codaEntity{
						ID:   "c-L-78eDsX4J",
						Name: "IE 6% Tax",
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
			NewInvoiceValidation: _codaEntity{
				ID:   "f-mvuXtisHj6",
				Name: "newInvoiceValidation",
			},
			NewInvoiceValidationIcon: _codaEntity{
				ID:   "f-DEQrOzH6JR",
				Name: "newInvoiceValidationIcon",
			},
			NewInvoiceRate: _codaEntity{
				ID:   "f-y2hpDful7D",
				Name: "newInvoiceRate",
			},
			NewInvoicePreviousRatesFilled: _codaEntity{
				ID:   "f-ROGDhTzctH",
				Name: "newInvoicePreviousRatesFilled",
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
			NewInvoiceBtn: _codaEntity{
				ID:   "ctrl-_VdNjxp4nS",
				Name: "newInvoiceBtn",
			},
			NewInvoiceMonth: _codaEntity{
				ID:   "ctrl-AMUYcRRj1w",
				Name: "newInvoiceMonth",
			},
		},
	}
}
