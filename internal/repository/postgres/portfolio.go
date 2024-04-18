package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pttrulez/investor-go/internal/types"
)

type PortfolioPostgres struct {
	db *sql.DB
}

func NewPortfolioPostgres(db *sql.DB) types.PortfolioRepository {
	return &PortfolioPostgres{db: db}
}

func (pg *PortfolioPostgres) Delete(id int) error {
	queryString := "DELETE FROM portfolios where id = $1;"
	_, err := pg.db.Exec(queryString, fmt.Sprint(id))
	if err != nil {
		return err
	}
	return nil
}

func (pg *PortfolioPostgres) GetById(id int) (*types.Portfolio, error) {
	queryString := `SELECT * FROM portfolios where id = $1;`
	// JOIN deals WHERE portfolios.id = deals.portfolio_id
	row := pg.db.QueryRow(queryString, fmt.Sprint(id))
	if row.Err() != nil {
		return nil, row.Err()
	}

	var p types.Portfolio
	err := row.Scan(&p.Id, &p.Compound, &p.Name, &p.UserId)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (pg *PortfolioPostgres) GetListByUserId(id int) ([]*types.Portfolio, error) {
	queryString := `SELECT * FROM portfolios where user_id = $1;`
	// JOIN deals WHERE portfolios.id = deals.portfolio_id
	rows, err := pg.db.Query(queryString, fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	portfolios := []*types.Portfolio{}
	for rows.Next() {
		p := &types.Portfolio{}
		err := rows.Scan(p.Id, p.Compound, p.Name, p.UserId)
		if err != nil {
			return nil, err
		}
		portfolios = append(portfolios, p)
	}

	return portfolios, nil
}

func (pg *PortfolioPostgres) Insert(u types.Portfolio) error {
	queryString := "INSERT INTO portfolios (compound, name, user_id) VALUES ($1, $2, $3);"
	_, err := pg.db.Exec(queryString, u.Compound, u.Name, u.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (pg *PortfolioPostgres) Update(data types.PortfolioUpdate) error {
	queryString := "UPDATE portfolios SET compound = $1, name = $2 WHERE id = $3;"

	_, err := pg.db.Exec(queryString, data.Compound, data.Name, data.Id)
	if err != nil {
		return err
	}
	return nil
}

// func (pg *PortfolioPostgres) Update(id string, data types.PortfolioUpdate) error {
// 	queryString := "UPDATE portfolios SET "
// 	args := []any{}
// 	argsQ := 1
// 	setValues := []string{}
// 	helpers.StructForEach(data, func(key string, value any) {
// 		if !helpers.IsNilPointer(value) {
// 			setValues = append(setValues, fmt.Sprintf("%s=$%d, ", key, argsQ))
// 			argsQ++
// 			args = append(args, value)
// 		}

// 	})
// 	sets := strings.TrimRight(strings.Join(setValues, ", "), ", ")
// 	queryString += fmt.Sprintf("%s WHERE id=$%d;", sets, argsQ)
// 	args = append(args, id)

// 	_, err := pg.db.Exec(queryString, args...)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
