package types

type Exchange string

const (
	Moex Exchange = "Moex"
)

func (e Exchange) Validate() bool {
	switch e {
	case Moex:
	default:
		return false
	}
	return true
}



type Role string

const (
	Admin    Role = "Admin"
	Investor Role = "Investor"
)

type SecurityType string

const (
	Bond     = "Bond"
	Currency = "Currency"
	Futures  = "Futures"
	Index    = "Index"
	Pif      = "Pif"
	Share    = "Share"
)

func (e SecurityType) Validate() bool {
	switch e {
	case Bond:
	case Currency:
	case Futures:
	case Index:
	case Pif:
	case Share:
	default:
		return false
	}
	return true
}

type TransactionType string

const (
	Cashout = "Cashout"
	Deposit = "Deposit"
)

func (e TransactionType) Validate() bool {
	switch e {
	case Cashout:
	case Deposit:
	default:
		return false
	}
	return true
}
