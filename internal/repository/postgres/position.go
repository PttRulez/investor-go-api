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

func (pg *PositionPostgres) Get(exchange types.Exchange, portfolioId int, securityId int,
	securityType types.SecurityType) (*types.Position, error) {
	queryString := `SELECT * FROM positions WHERE exchange = $1 AND portfolio_id = $2 AND security_id = $3
		AND security_type = $4;`

	var position types.Position

	row := pg.db.QueryRow(queryString, exchange, portfolioId, securityId, securityType)

	err := row.Scan(&position.Id, &position.Amount, &position.AveragePrice, &position.Comment, &position.Exchange,
		&position.PortfolioId, &position.SecurityId, &position.SecurityType, &position.TargetPrice)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		fmt.Println("[PositionPostgres Get] Failed to scan:", err)
		return nil, err
	}

	return &position, nil
}

func (pg *PositionPostgres) GetListByPortfolioId(id int) ([]*types.Position, error) {
	queryString := `SELECT * FROM positions WHERE portfolio_id = $1;`

	rows, err := pg.db.Query(queryString, id)
	if err != nil {
		fmt.Println("[PositionPostgres GetPositionsByPortfolioId] Failed to execute query:", err)
		return nil, err
	}
	defer rows.Close()

	positions := make([]*types.Position, 0)

	for rows.Next() {
		var position types.Position
		err := rows.Scan(&position.Id, &position.Amount, &position.AveragePrice, &position.Comment, &position.Exchange,
			&position.PortfolioId, &position.SecurityId, &position.SecurityType, &position.TargetPrice)
		if err != nil {
			fmt.Println("[PositionPostgres GetPositionsByPortfolioId] Failed to scan:", err)
			return nil, err
		}
		positions = append(positions, &position)
	}

	return positions, nil
}

func (pg *PositionPostgres) Insert(p *types.Position) error {
	queryString := `INSERT INTO positions (amount, average_price, comment, exchange,
    portfolio_id, security_id, security_type,  target_price) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ;`

	_, err := pg.db.Exec(queryString, p.Amount, p.AveragePrice, p.Comment, p.Exchange,
		p.PortfolioId, p.SecurityId, p.SecurityType, p.TargetPrice)
	if err != nil {
		fmt.Println("[PositionPostgres Insert] Failed to execute query:", err)
		return err
	}
	return nil
}

func (pg *PositionPostgres) Update(p *types.Position) error {
	queryString := `UPDATE positions SET amount = $1, average_price = $2, comment = $3, exchange = $4,
    portfolio_id = $5, security_id = $6, security_type = $7, target_price = $8
    WHERE id = $9;`

	_, err := pg.db.Exec(queryString, p.Amount, p.AveragePrice, p.Comment, p.Exchange,
		p.PortfolioId, p.SecurityId, p.SecurityType, p.TargetPrice, p.Id)
	if err != nil {
		fmt.Println("[PositionPostgres Update] Failed to execute query:", err)
		return err
	}
	return nil
}
