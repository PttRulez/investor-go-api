package types

import (
	tmoex "github.com/pttrulez/investor-go/internal/types/moex"
)

type Repository struct {
	Deal      DealRepository
	Expert    ExpertRepository
	Moex      MoexRepository
	Position  PositionRepository
	Portfolio PortfolioRepository
	User      UserRepository
}

type CashoutRepository interface {
	Delete(id int) error
	Insert(c *Cashout) error
}
type DepositRepository interface {
	Delete(id int) error
	Insert(c *Deposit) error
}
type DealRepository interface {
	Delete(id int) error
	GetDealsListForSecurity(exchange Exchange, portfolioId int, securityType SecurityType, securityId int) ([]*Deal, error)
	Insert(d *RepoCreateDeal) error
	Update(d Deal) error
}

type ExpertRepository interface {
	Delete(id int) error
	GetListByUserId(userId int) ([]*Expert, error)
	Insert(e Expert) error
	Update(e Expert) error
}

type MoexRepository struct {
	Bonds  MoexBondRepository
	Shares MoexShareRepository
}

type MoexBondRepository interface {
	Insert(bond tmoex.Bond) error
	GetByTicker(ticker string) (*tmoex.Bond, error)
}

type MoexShareRepository interface {
	Insert(share *tmoex.Share) error
	GetByTicker(ticker string) (*tmoex.Share, error)
}

type OpinionRepository interface {
	Insert(e Opinion) error
	GetListByUserId(userId int) ([]*Opinion, error)
	Update(e Opinion) error
	Delete(id int) error
}

type PortfolioRepository interface {
	Delete(id int) error
	GetById(id int) (*Portfolio, error)
	GetListByUserId(userId int) ([]*Portfolio, error)
	Insert(u Portfolio) error
	Update(u PortfolioUpdate) error
}

type PositionRepository interface {
	GetListByPortfolioId(id int) ([]*Position, error)
	Get(exchange Exchange, portfolioId int, securityId int,
		securityType SecurityType) (*Position, error)
	Insert(p *Position) error
	Update(p *Position) error
}

type UserRepository interface {
	Insert(u User) error
	GetByEmail(email string) (*User, error)
	GetById(id int) (*User, error)
}
