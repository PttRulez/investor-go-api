package types

type Portfolio struct {
	Compound bool   `json:"compound"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserId   int    `json:"-"`
}

type PortfolioUpdate struct {
	Compound *bool   `json:"compound,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type PortfolioRepository interface {
	CreatePortfolio(u Portfolio) (*Portfolio, error)
	DeletePortfolioById(id int) error
	GetPortfolioById(id int) (*Portfolio, error)
	GetPortfolioListByUserId(userId int) ([]*Portfolio, error)
	UpdatePortfolio(id string, u PortfolioUpdate) (*Portfolio, error)
}
