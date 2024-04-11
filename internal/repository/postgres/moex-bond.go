package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pttrulez/investor-go/internal/services/moex"
	"github.com/pttrulez/investor-go/internal/types"
)

type MoexBondsPostgres struct {
	db *sql.DB
}

func NewMoexBondsPostgres(db *sql.DB) types.MoexBondRepository {
	return &MoexBondsPostgres{db: db}
}

func (pg *MoexBondsPostgres) Create(bond moex.Bond) (*moex.Bond, error) {
	querySting := `INSERT INTO moex_bonds (board, engine, market, id, name, shortname, ticker)
    VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;`

	row := pg.db.QueryRow(querySting, bond.Board, bond.Engine, bond.Market, bond.Id, bond.Name, bond.ShortName, bond.Ticker)
	if row.Err() != nil {
		fmt.Println("Failed to execute query:", row.Err())
		return nil, row.Err()
	}

	var newBond moex.Bond
	err := row.Scan(&newBond.Board, &newBond.Engine, &newBond.Market, &newBond.Id, &newBond.Name,
		&newBond.ShortName, &newBond.Ticker)
	if err != nil {
		fmt.Println("[MoexBondsPostgres.Create] - Failed to scan:", err)
		return nil, err
	}

	return &newBond, nil
}

func (pg *MoexBondsPostgres) GetBulk(ids []int) ([]*moex.Bond, error) {

	return nil, nil
}

func (pg *MoexBondsPostgres) GetByTicker(ticker string) (*moex.Bond, error) {

	return nil, nil
}

func (pg *MoexBondsPostgres) GetById(id int) (*moex.Bond, error) {

	return nil, nil
}