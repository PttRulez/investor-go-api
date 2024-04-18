package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pttrulez/investor-go/internal/types"
	tmoex "github.com/pttrulez/investor-go/internal/types/moex"
)

type MoexBondsPostgres struct {
	db *sql.DB
}

func NewMoexBondsPostgres(db *sql.DB) types.MoexBondRepository {
	return &MoexBondsPostgres{db: db}
}

func (pg *MoexBondsPostgres) GetByTicker(ticker string) (*tmoex.Bond, error) {
	queryString := `SELECT * FROM moex_bonds WHERE ticker = $1;`

	row := pg.db.QueryRow(queryString, ticker)

	bond := &tmoex.Bond{}
	err := row.Scan(&bond.Board, &bond.Engine, &bond.Market, &bond.Id, &bond.Name, &bond.ShortName, &bond.Ticker)
	if err != nil {
		fmt.Println("[MoexBondsPostgres GetByTicker] Failed to execute query:", err)
		return nil, err
	}

	return bond, nil
}

func (pg *MoexBondsPostgres) Insert(bond tmoex.Bond) error {
	querySting := `INSERT INTO moex_bonds (board, engine, market, id, name, shortname, ticker)
    VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;`

	_, err := pg.db.Exec(querySting, bond.Board, bond.Engine, bond.Market, bond.Id, bond.Name, bond.ShortName, bond.Ticker)
	if err != nil {
		fmt.Println("[NewMoexBondsPostgres Insert] Failed to execute query:", err)
		return err
	}

	return nil
}
