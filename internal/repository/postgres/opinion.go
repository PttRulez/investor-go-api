package postgres

import (
	"database/sql"

	"github.com/pttrulez/investor-go/internal/types"
)

type OpinionPostgres struct {
	db *sql.DB
}

func NewOpinionPostgres(db *sql.DB) types.OpinionRepository {
	return &OpinionPostgres{db: db}
}

func (pg *OpinionPostgres) Delete(id int) error {
	queryString := "DELETE FROM opinions where id = $1;"
	_, err := pg.db.Exec(queryString, id)
	if err != nil {
		return err
	}
	return nil
}

func (pg *OpinionPostgres) Insert(o types.Opinion) error {
	queryString := `INSERT INTO opinions (date, exchange, expert_id, security_id, 
    security_type, source_link, target_price, type, user_id) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	_, err := pg.db.Exec(queryString, o.Date, o.Exchange, o.ExpertId, o.SecurityId,
		o.SecurityType, o.SourceLink, o.TargetPrice, o.Type, o.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (pg *OpinionPostgres) Update(o types.Opinion) error {
	queryString := `UPDATE opinions SET date = $1, exchange = $2, expert_id = $3,
    security_id = $4, security_type = $5, source_link = $6, target_price = $7,
    type = $8, user_id = $9 WHERE id = $10;`

	_, err := pg.db.Exec(queryString, o.Date, o.Exchange, o.ExpertId, o.SecurityId,
		o.SecurityType, o.SourceLink, o.TargetPrice, o.Type, o.UserId, o.Id)
	if err != nil {
		return err
	}

	return nil
}

func (pg *OpinionPostgres) GetListByUserId(userId int) ([]*types.Opinion, error) {
	queryString := `SELECT * FROM opinions WHERE user_id = $1;`
	rows, err := pg.db.Query(queryString, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	opinions := []*types.Opinion{}

	for rows.Next() {
		var o types.Opinion
		err = rows.Scan(&o.Id, &o.Date, &o.Exchange, &o.ExpertId, &o.SecurityId,
			&o.SecurityType, &o.SourceLink, &o.TargetPrice, &o.Type, &o.UserId)
		if err != nil {
			return nil, err
		}
		opinions = append(opinions, &o)
	}

	return opinions, nil
}
