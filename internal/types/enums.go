package types

type Exchange string
const (
	Moex Exchange = "Moex"
)

type DealType string
const (
	Buy  DealType = "Buy"
	Sell DealType = "Sell"
)

type OpinionType string

const (
	Flat      OpinionType = "Flat"
	General   OpinionType = "General"
	Growth    OpinionType = "Growth"
	Reduction OpinionType = "Reduction"
)

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

type TransactionType string
const (
	Cashout = "Cashout"
	Deposit = "Deposit"
)
