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

func (pg *MoexSharesPostgres) GetByTicker(ticker string) (*tmoex.Share, error) {
	querySting := `SELECT * FROM moex_shares WHERE ticker = $1;`

	row := pg.db.QueryRow(querySting, ticker)
	if row.Err() != nil {
		fmt.Println("[MoexSharesPostgres GetByTicker] Failed to execute query:", row.Err())
		return nil, row.Err()
	}

	var share tmoex.Share
	err := row.Scan(&share.Id, &share.Board, &share.Engine, &share.Market, &share.Name, &share.ShortName, &share.Ticker)
	if err != nil {
		fmt.Println("[MoexSharesPostgres GetByTicker] Failed to scan:", err)
		return nil, err
	}

	return &share, nil
}

func (pg *MoexSharesPostgres) Insert(share *tmoex.Share) error {
	querySting := `INSERT INTO moex_shares (board, engine, market, name, shortname, ticker)
    VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;`

	_, err := pg.db.Exec(querySting, share.Board, share.Engine, share.Market, share.Name, share.ShortName, share.Ticker)
	if err != nil {
		fmt.Println("[MoexSharePostgres Insert] Failed to execute query:", err)
		return err
	}

	return nil
}
