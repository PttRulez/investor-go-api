package postgres

import (
	"database/sql"

	"github.com/pttrulez/investor-go/internal/types"
)

type DealPostgres struct {
	db *sql.DB
}

func NewDealPostgres(db *sql.DB) types.DealRepository {
	return &DealPostgres{db: db}
}

func (pg *DealPostgres) CreateDeal(d *types.RepoCreateDeal) (*types.Deal, error) {
	queryString := `INSERT INTO deals (amount, date, exchange, portfolio_id, price, security_type, ticker, type) VALUES 
    ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;`
	row := pg.db.QueryRow(queryString, d.Amount, d.Date, d.Exchange, d.PortfolioId, d.Price, d.SecurityType, d.Ticker, d.Type)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var deal types.Deal
	err := row.Scan(&deal.Id, &deal.Amount, &deal.Date, &deal.Exchange, &deal.PortfolioId, &deal.Price, &deal.SecurityType, &deal.Ticker, &deal.Type)
	if err != nil {
		return nil, err
	}
	return &deal, nil
}

func (pg *DealPostgres) GetDealListByPortfolioId(portfolioId int) ([]*types.Deal, error) {

	return nil, nil
}

func (pg *DealPostgres) UpdateDeal(d types.Deal) (*types.Deal, error) {
	return nil, nil
}

func (pg *DealPostgres) DeleteDealById(id int) error {
	return nil
}

func (pg *DealPostgres) GetDealById(id int) (*types.Deal, error) {
	return nil, nil
}
