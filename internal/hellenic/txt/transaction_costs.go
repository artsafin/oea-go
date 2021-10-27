package txt

type transactionCosts string

const (
	TransactionCostsOurs        = transactionCosts('O')
	TransactionCostsBeneficiary = transactionCosts('B')
	TransactionCostsShared      = transactionCosts('S')
)
