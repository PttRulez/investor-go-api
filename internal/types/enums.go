package types

type Role string

const (
	Admin    Role = "Admin"
	Investor Role = "Investor"
)

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

type SecurityType string

const (
	Bond     SecurityType = "Bond"
	Currency SecurityType = "Currency"
	Futures  SecurityType = "Futures"
	Index    SecurityType = "Index"
	Pif      SecurityType = "Pif"
	Share    SecurityType = "Share"
)

func (e SecurityType) Validate() bool {
	switch e {
	case Bond:
	case Share:
	default:
		return false
	}
	return true
}
