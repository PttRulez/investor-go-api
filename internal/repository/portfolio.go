package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/pttrulez/investor-go/internal/lib/helpers"
	"github.com/pttrulez/investor-go/internal/types"
)

type Portfolio = types.Portfolio
type PortfolioRepository = types.PortfolioRepository
type PortfolioPostgres struct {
	db *sql.DB
}

func NewPortfolioPostgres(db *sql.DB) types.PortfolioRepository {
	return &PortfolioPostgres{db: db}
}

func (pg *PortfolioPostgres) CreatePortfolio(u Portfolio) (*Portfolio, error) {
	queryString := "INSERT INTO portfolios (compound, name, user_id) VALUES ($1, $2, $3) RETURNING *;"
	row := pg.db.QueryRow(queryString, u.Compound, u.Name, u.UserId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var p Portfolio
	err := row.Scan(&p.Id, &p.Compound, &p.Name, &p.UserId)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (pg *PortfolioPostgres) DeletePortfolioById(id int) error {
	queryString := "DELETE FROM portfolios where id = $1;"
	_, err := pg.db.Query(queryString, fmt.Sprint(id))
	if err != nil {
		return err
	}
	return nil
}

func (pg *PortfolioPostgres) GetPortfolioById(id int) (*Portfolio, error) {
	queryString := `SELECT * FROM portfolios where id = $1;`
	// JOIN deals WHERE portfolios.id = deals.portfolio_id
	row := pg.db.QueryRow(queryString, fmt.Sprint(id))
	if row.Err() != nil {
		return nil, row.Err()
	}

	var p Portfolio
	err := row.Scan(&p.Id, &p.Compound, &p.Name, &p.UserId)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (pg *PortfolioPostgres) GetPortfolioListByUserId(id int) ([]*Portfolio, error) {
	queryString := `SELECT * FROM portfolios where user_id = $1;`
	// JOIN deals WHERE portfolios.id = deals.portfolio_id
	rows, err := pg.db.Query(queryString, fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	portfolios := []*Portfolio{}
	for rows.Next() {
		p := &Portfolio{}
		err := rows.Scan(p.Id, p.Compound, p.Name, p.UserId)
		if err != nil {
			return nil, err
		}
		portfolios = append(portfolios, p)
	}

	return portfolios, nil
}

func (pg *PortfolioPostgres) UpdatePortfolio(id string, data types.PortfolioUpdate) (*Portfolio, error) {
	queryString := "UPDATE portfolios SET "
	args := []any{}
	argsQ := 1
	setValues := []string{}
	helpers.StructForEach(data, func(key string, value any) {
		if !helpers.IsNilPointer(value) {
			setValues = append(setValues, fmt.Sprintf("%s=$%d, ", key, argsQ))
			argsQ++
			args = append(args, value)
		}

	})
	sets := strings.TrimRight(strings.Join(setValues, ", "), ", ")
	queryString += fmt.Sprintf("%s WHERE id=$%d RETURNING *;", sets, argsQ)
	args = append(args, id)

	fmt.Println("queryString:", queryString)
	fmt.Println()
	row := pg.db.QueryRow(queryString, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var p Portfolio
	err := row.Scan(&p.Id, &p.Compound, &p.Name, &p.UserId)
	if err != nil {
		log.Fatal("[repo UpdatePortfolio] row Scan err:", data)
		return nil, err
	}
	return &p, nil
}
