package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pttrulez/investor-go/internal/types"
	tmoex "github.com/pttrulez/investor-go/internal/types/moex"
)

type MoexSharesPostgres struct {
	db *sql.DB
}

func NewMoexSharesPostgres(db *sql.DB) types.MoexShareRepository {
	return &MoexSharesPostgres{db: db}
}

func (pg *MoexSharesPostgres) Create(share *tmoex.Share) (*tmoex.Share, error) {
	querySting := `INSERT INTO moex_shares (board, engine, market, name, shortname, ticker)
    VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;`

	row := pg.db.QueryRow(querySting, share.Board, share.Engine, share.Market, share.Name, share.ShortName, share.Ticker)
	if row.Err() != nil {
		fmt.Println("Failed to execute query:", row.Err())
		return nil, row.Err()
	}

	var newShare tmoex.Share
	err := row.Scan(&newShare.Id, &newShare.Board, &newShare.Engine, &newShare.Market, &newShare.Name,
		&newShare.ShortName, &newShare.Ticker)
	if err != nil {
		fmt.Println("[MoexSharesPostgres.Create] - Failed to scan:", err)
		return nil, err
	}

	return &newShare, nil
}

func (pg *MoexSharesPostgres) GetBulk(ids []int) ([]*tmoex.Share, error) {

	return nil, nil
}

func (pg *MoexSharesPostgres) GetByTicker(ticker string) (*tmoex.Share, error) {
	querySting := `SELECT * FROM moex_bonds WHERE ticker = $1;`

	row := pg.db.QueryRow(querySting, ticker)
	if row.Err() != nil {
		fmt.Println("[MoexSharesPostgres GetByTicker] Failed to execute query:", row.Err())
		return nil, row.Err()
	}

	var share tmoex.Share
	err := row.Scan(&share.Board, &share.Engine, &share.Market, &share.Id, &share.Name, &share.ShortName, &share.Ticker)
	if err != nil {
		return nil, err
	}

	return &share, nil
}

func (pg *MoexSharesPostgres) GetById(id int) (*tmoex.Share, error) {

	return nil, nil
}
