package types

import "github.com/pttrulez/investor-go/internal/services/moex"

type Repository struct {
	Deal      DealRepository
	MoexBond  MoexBondRepository
	Portfolio PortfolioRepository
	User      UserRepository
}

type DealRepository interface {
	Cre
}

type PortfolioRepository interface {
	CreatePortfolio(u Portfolio) (*Portfolio, error)
	DeletePortfolioById(id int) error
	GetPortfolioById(id int) (*Portfolio, error)
	GetPortfolioListByUserId(userId int) ([]*Portfolio, error)
	UpdatePortfolio(id string, u PortfolioUpdate) (*Portfolio, error)
}

type MoexBondRepository interface {
	Create(bond moex.Bond) (*moex.Bond, error)
	GetBulk(ids []int) ([]*moex.Bond, error)
	GetByTicker(ticker string) (*moex.Bond, error)
	GetById(id int) (*moex.Bond, error)
}

type MoexShareRepository interface {
	Create(share moex.Share) (*moex.Share, error)
	GetBulk(ids []int) ([]*moex.Share, error)
	GetByTicker(ticker string) (*moex.Share, error)
	GetById(id int) (*moex.Share, error)
}

type UserRepository interface {
	CreateUser(u User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
}