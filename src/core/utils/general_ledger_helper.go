package utils

const (
	/* deposit transaction */
	depositService  = "deposit"
	depositTafsilID = 1

	/* order transaction */
	orderService           = "order"
	cancelOrderTafsilID    = 3
	matchOrderTafsilID     = 4
	changeCurrencyTafsilID = 5

	/* withdraw transaction */
	withdrawService             = "withdraw"
	withdrawRequestTafsilID     = 6
	cancelWithdrawTafsilID      = 7
	confirmWithdrawTafsilID     = 8
	bankPortWageTafsilID        = 10
	withdrawFeeCustomerTafsilID = 12
	withdrawFeeBusinessTafsilID = 13

	/* debit and credit types*/
	null           = 0
	cash           = 1
	equity         = 2
	cashReceivable = 3
	cashPayable    = 4
	expense        = 5
	revenue        = 6

	// TODO
	//  تفصیل کمیسیون --> 9
	//  INTERNAL FEE --> 11
	//  سود حاصل از کارمزد شبکه--> 14
)

// deposits ----------------------------------------------------------->

/* get deposit service name */
func GetDepositServiceName() string {
	return depositService
}

/* get deposit tafsil ID */
func GetDepositTafsilID() uint8 {
	return depositTafsilID
}

// orders --------------------------------------------------------------->

/* get order service name */
func GetOrderServiceName() string {
	return orderService
}

/* get cancel order tafsil ID */
func GetCancelOrderTafsilID() uint8 {
	return cancelOrderTafsilID
}

/* get match order tafsil ID */
func GetMatchOrderTafsilID() uint8 {
	return matchOrderTafsilID
}

/* get change currency tafsil ID  */
func GetChangeCurrencyTafsilID() uint8 {
	return changeCurrencyTafsilID
}

// withdraws --------------------------------------------------------------->

/* get withdraw service name */
func GetWithdrawServiceName() string {
	return withdrawService
}

/* get withdraw request tafsil ID */
func GetWithdrawRequestTafsilID() uint8 {
	return withdrawRequestTafsilID
}

/* get cancel withdraw tafsil ID */
func GetCancelWithdrawTafsilID() uint8 {
	return cancelWithdrawTafsilID
}

/* get confirm withdraw tafsil ID */
func GetConfirmWithdrawTafsilID() uint8 {
	return confirmWithdrawTafsilID
}

/* get withdraw bank port wage tafsil ID */
func GetBankPortWageTafsilID() uint8 {
	return bankPortWageTafsilID
}

/* get withdraw fee customer tafsil ID */
func GetWithdrawFeeCustomerTafsilID() uint8 {
	return withdrawFeeCustomerTafsilID
}

/* get withdraw fee business tafsil ID */
func GetWithdrawFeeBusinessTafsilID() uint8 {
	return withdrawFeeBusinessTafsilID
}

// credits and debits types -------------------------------------------------->

/* get null ID */
func GetNull() uint8 {
	return null
}

/* get cash ID */
func GetCash() uint8 {
	return cash
}

/* get equity ID */
func GetEquity() uint8 {
	return equity
}

/* get cashReceivable ID */
func GetCashReceivable() uint8 {
	return cashReceivable
}

/* get cashPayable ID */
func GetCashPayable() uint8 {
	return cashPayable
}

/* get expense ID */
func GetExpense() uint8 {
	return expense
}

/* get revenue ID */
func GetRevenue() uint8 {
	return revenue
}
