package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pttrulez/investor-go/internal/types"
)

type PositionPostgres struct {
	db *sql.DB
}

func NewPositionPostgres(db *sql.DB) types.PositionRepository {
	return &PositionPostgres{db: db}
}

func (pg *PositionPostgres) CreatePosition(p types.Position) error {
	queryString := `INSERT INTO positions (amount, average_price, comment, exchange,
    portfolio_id, security_id, security_type, trade_saldo, target_price) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ;`

	row := pg.db.QueryRow(queryString, p.Amount, p.AveragePrice, p.Comment, p.Exchange,
		p.PortfolioId, p.SecurityId, p.SecurityType, p.TradeSaldo, p.TargetPrice)
	if row.Err() != nil {
		fmt.Println("[PositionPostgres CreatePosition] Failed to execute query:", row.Err())
		return row.Err()
	}
	return nil
}

func (pg *PositionPostgres) UpdatePosition(p types.Position) error {
	queryString := `UPDATE positions SET amount = $1, average_price = $2, comment = $3, exchange = $4,
    portfolio_id = $5, security_id = $6, security_type = $7, trade_saldo = $8, target_price = $9
    WHERE id = $10;`

	row := pg.db.QueryRow(queryString, p.Amount, p.AveragePrice, p.Comment, p.Exchange,
		p.PortfolioId, p.SecurityId, p.SecurityType, p.TradeSaldo, p.TargetPrice, p.Id)
	if row.Err() != nil {
		fmt.Println("[PositionPostgres UpdatePosition] Failed to execute query:", row.Err())
		return row.Err()
	}
	return nil
}

func (pg *PositionPostgres) GetPositionsByPortfolioId(id int) ([]types.Position, error) {
	queryString := `SELECT * FROM positions WHERE portfolio_id = $1;`

	rows, err := pg.db.Query(queryString, id)
	if err != nil {
		fmt.Println("[PositionPostgres GetPositionsByPortfolioId] Failed to execute query:", err)
		return nil, err
	}
	defer rows.Close()

	positions := make([]types.Position, 0)

	for rows.Next() {
		var position types.Position
		err := rows.Scan(&position.Id, &position.Amount, &position.AveragePrice, &position.Comment, &position.Exchange,
			&position.PortfolioId, &position.SecurityId, &position.SecurityType, &position.TradeSaldo, &position.TargetPrice)
		if err != nil {
			fmt.Println("[PositionPostgres GetPositionsByPortfolioId] Failed to scan:", err)
			return nil, err
		}
		positions = append(positions, position)
	}

	return positions, nil
}
