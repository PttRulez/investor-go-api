package types

import (
	tmoex "github.com/pttrulez/investor-go/internal/types/moex"
)

type Repository struct {
	Deal      DealRepository
	Moex      MoexRepository
	Position  PositionRepository
	Portfolio PortfolioRepository
	User      UserRepository
}

type DealRepository interface {
	CreateDeal(d *RepoCreateDeal) (*Deal, error)
	UpdateDeal(d Deal) (*Deal, error)
	DeleteDealById(id int) error
	GetDealById(id int) (*Deal, error)
}

type PortfolioRepository interface {
	CreatePortfolio(u Portfolio) (*Portfolio, error)
	DeletePortfolioById(id int) error
	GetPortfolioById(id int) (*Portfolio, error)
	GetPortfolioListByUserId(userId int) ([]*Portfolio, error)
	UpdatePortfolio(id string, u PortfolioUpdate) (*Portfolio, error)
}

type PositionRepository interface {
	CreatePosition(p Position) error
	UpdatePosition(p Position) error
}

type MoexRepository struct {
	Bonds  MoexBondRepository
	Shares MoexShareRepository
}

type MoexBondRepository interface {
	Create(bond tmoex.Bond) (*tmoex.Bond, error)
	GetBulk(ids []int) ([]*tmoex.Bond, error)
	GetByTicker(ticker string) (*tmoex.Bond, error)
	GetById(id int) (*tmoex.Bond, error)
}

type MoexShareRepository interface {
	Create(share *tmoex.Share) (*tmoex.Share, error)
	GetBulk(ids []int) ([]*tmoex.Share, error)
	GetByTicker(ticker string) (*tmoex.Share, error)
	GetById(id int) (*tmoex.Share, error)
}

type UserRepository interface {
	CreateUser(u User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
}
