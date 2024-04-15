package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pttrulez/investor-go/internal/config"
	"github.com/pttrulez/investor-go/internal/types"
)

func NewPostgresRepo(cfg config.PostgresConfig) (*types.Repository, error) {
	connStr := fmt.Sprintf(`postgresql://%v:%v@%v:%v/%v?sslmode=%v`,
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("sql.Open err", err)
		return nil, err
	}

	err = createAllTables(db)
	if err != nil {
		log.Fatal("createAllTables err", err)
		return nil, err
	}

	return &types.Repository{
		Deal: NewDealPostgres(db),
		Moex: types.MoexRepository{
			Bonds:  NewMoexBondsPostgres(db),
			Shares: NewMoexSharesPostgres(db),
		},
		Portfolio: NewPortfolioPostgres(db),
		Position:  NewPositionPostgres(db),
		User:      NewUserPostgres(db),
	}, nil
}
func createAllTables(db *sql.DB) error {
	err := createUsersTable(db)
	if err != nil {
		return err
	}
	err = createPortfoliosTable(db)
	if err != nil {
		return err
	}
	err = createExpertsTable(db)
	if err != nil {
		return err
	}
	err = createDealsTable(db)
	if err != nil {
		return err
	}
	err = createOpinionsTable(db)
	if err != nil {
		return err
	}
	err = createOpinionsOnPositionsTable(db)
	if err != nil {
		return err
	}
	err = createPositionsTable(db)
	if err != nil {
		return err
	}

	//  ---------------------- MOEX ----------------------
	err = createMoexBondsTable(db)
	if err != nil {
		return err
	}
	err = createMoexSharesTable(db)
	if err != nil {
		return err
	}
	err = createMoexCurrenciesTable(db)
	if err != nil {
		return err
	}
	err = createTransactionsTable(db)
	if err != nil {
		return err
	}

	return err
}

// ---------------------- FUNCS FOR TABLES CREATION  ----------------------
func createDealsTable(db *sql.DB) error {
	queryString := `create table if not exists deals (
		id serial primary key,
		amount integer not null,
		date date not null,
		exchange varchar(50) not null,
		portfolio_id integer references portfolios(id) not null,
		price numeric(10, 2) not null,
		security_type varchar(50) not null,
		ticker varchar(50) not null,
		type varchar(50) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createDealsTable] err", err)
	}
	return err
}
func createExpertsTable(db *sql.DB) error {
	queryString := `create table if not exists experts (
		id serial primary key,
		avatarUrl varchar(100),
		name varchar(50) not null,
		user_id integer references users(id) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createExpertsTable] err", err)
	}
	return err
}
func createOpinionsTable(db *sql.DB) error {
	queryString := `create table if not exists opinions (
		id serial primary key,
		date date not null,
		exchange varchar(50) not null,
		expert_id integer references experts(id) not null,
		text text not null,
		security_id integer not null,
		security_type varchar(50) not null,
		source_link varchar(120),
		target_price numeric(10, 2),
		type varchar(50) not null,
		user_id integer references users(id) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createOpinionsTable] err", err)
	}
	return err
}
func createOpinionsOnPositionsTable(db *sql.DB) error {
	queryString := `create table if not exists opinions_on_positions (
		id serial primary key,
		opinion_id integer references opinions(id) not null,
		portfolio_id integer references portfolios(id) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createOpinionsOnPositionsTable] err", err)
	}
	return err
}
func createPortfoliosTable(db *sql.DB) error {
	queryString := `create table if not exists portfolios (
		id serial primary key,
		compound boolean not null,
		name varchar(50) not null,
		user_id integer references users(id) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createPortfoliosTable] err: ", err)
	}
	return err
}
func createPositionsTable(db *sql.DB) error {
	queryString := `create table if not exists positions (
		id serial primary key,
		amount integer not null,
		average_price numeric(10, 2) not null,
		comment text,
		exchange varchar(50) not null,
		portfolio_id integer references portfolios(id) not null,
		security_id integer not null,
		security_type varchar(50) not null,
		trade_saldo integer not null,
		target_price numeric(10, 2)
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createPositionsTable] err", err)
	}
	return err
}
func createTransactionsTable(db *sql.DB) error {
	queryString := `create table if not exists transactions (
		id serial primary key,
		amount integer not null,
		date date not null,
		portfolio_id integer references portfolios(id) not null,
		type varchar(50) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createTransactionsTable] err", err)
	}
	return err
}
func createUsersTable(db *sql.DB) error {
	queryString := `create table if not exists users (
		id serial primary key,
		email varchar(50) unique not null,
		hashed_password varchar(100) not null,
		name varchar(50) not null,
		role varchar(10) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createUsersTable] err", err)
	}
	return err
}

func createMoexBondsTable(db *sql.DB) error {
	queryString := `create table if not exists moex_bonds (
		id serial primary key,
		board varchar(50) not null,
		engine varchar(50) not null,
		market varchar(50) not null,
		name varchar(100) not null,
		shortName varchar(50) not null,
		ticker varchar(10) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createMoexBondsTable] err", err)
	}
	return err
}
func createMoexCurrenciesTable(db *sql.DB) error {
	queryString := `create table if not exists moex_currencies (
		id serial primary key,
		board varchar(50) not null,
		engine varchar(50) not null,
		market varchar(50) not null,
		name varchar(120) not null,
		shortName varchar(50) not null,
		ticker varchar(10) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createMoexSharesTable] err", err)
	}
	return err
}
func createMoexSharesTable(db *sql.DB) error {
	queryString := `create table if not exists moex_shares (
		id serial primary key,
		board varchar(50) not null,
		engine varchar(50) not null,
		market varchar(50) not null,
		name varchar(120) not null,
		shortName varchar(50) not null,
		ticker varchar(10) not null
	)`

	_, err := db.Exec(queryString)
	if err != nil {
		log.Fatal("[createMoexSharesTable] err", err)
	}
	return err
}
