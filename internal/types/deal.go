package types

import "time"

type Deal struct {
	Amount      int       `json:"amount" db:"amount" validate:"required,number"`
	Date        time.Time `json:"date" db:"date" validate:"required"`
	Exchange    Exchange  `json:"exchange" db:"exchange" validate:"required,is-exchange"`
	Id          int       `json:"id" db:"id"`
	PortfolioId int       `json:"portfolioId" db:"portfolio_id" validate:"required,number"`
	Price       float64   `json:"price" db:"price"  validate:"required,price"`
	// SecurityId   int          `json:"securityId" db:"security_id" validate:"required"`
	SecurityType SecurityType `json:"securityType" db:"security_type" validate:"required,securityType"`
	Ticker       string       `json:"ticker" db:"ticker" validate:"required"`
	Type         DealType     `json:"type" db:"type" validate:"required,dealType"`
}

type RepoCreateDeal struct {
	Deal
	Id int `json:"-" db:"-"`
}

type DealType string

const (
	Buy  DealType = "Buy"
	Sell DealType = "Sell"
)

func (e DealType) Validate() bool {
	switch e {
	case Buy:
	case Sell:
	default:
		return false
	}
	return true
}
