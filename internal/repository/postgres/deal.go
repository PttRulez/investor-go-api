package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pttrulez/investor-go/internal/types"
)

type DealPostgres struct {
	db *sql.DB
}

func NewDealPostgres(db *sql.DB) types.DealRepository {
	return &DealPostgres{db: db}
}

func (pg *DealPostgres) Delete(id int) error {
	queryString := `DELETE FROM deals WHERE id = $1;`
	_, err := pg.db.Exec(queryString, id)
	if err != nil {
		fmt.Println("[DealPostgres.Delete] - Failed to execute query:", err)
		return err
	}

	return nil
}

func (pg *DealPostgres) GetDealsListForSecurity(exchange types.Exchange, portfolioId int,
	securityType types.SecurityType, securityId int) ([]*types.Deal, error) {
	queryString := `SELECT * FROM deals WHERE exchange = $1 AND security_type = $2 AND security_id = $3 portfolio_id = $4 ORDER BY id DESC;`
	rows, err := pg.db.Query(queryString, exchange, securityType, securityId, portfolioId)
	if err != nil {
		fmt.Println("[DealPostgres.GetDealsBySecurityId] - Failed to execute query:", err)
		return nil, err
	}

	deals := make([]*types.Deal, 0)
	for rows.Next() {
		var deal types.Deal
		err := rows.Scan(&deal.Id, &deal.Amount, &deal.Date, &deal.Exchange, &deal.PortfolioId, &deal.Price, &deal.SecurityId, &deal.SecurityType, &deal.Ticker, &deal.Type)
		if err != nil {
			fmt.Println("[DealPostgres.GetDealsBySecurityId] - Failed to scan:", err)
			return nil, err
		}
		deals = append(deals, &deal)
	}
	return deals, nil
}

func (pg *DealPostgres) Insert(d *types.RepoCreateDeal) error {
	queryString := `INSERT INTO deals (amount, date, exchange, portfolio_id, price, security_id, security_type, ticker, type) VALUES 
    ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;`
	_, err := pg.db.Exec(queryString, d.Amount, d.Date, d.Exchange, d.PortfolioId, d.Price, d.SecurityId, d.SecurityType, d.Ticker, d.Type)
	if err != nil {
		fmt.Println("[DealPostgres.CreateDeal] failed to execute query:", err)
		return err
	}
	return nil
}

func (pg *DealPostgres) Update(d types.Deal) error {
	return nil
}
