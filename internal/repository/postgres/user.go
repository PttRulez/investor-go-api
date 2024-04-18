package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pttrulez/investor-go/internal/types"
)

type User = types.User
type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) types.UserRepository {
	return &UserPostgres{db: db}
}

func (pg *UserPostgres) Insert(u User) error {
	querySting := "INSERT INTO users (email, hashed_password, name, role) VALUES ($1, $2, $3, $4);"
	_, err := pg.db.Exec(querySting, u.Email, u.HashedPassword, u.Name, u.Role)
	if err != nil {
		fmt.Println("[UserPostgres.Insert] Failed to execute query:", err)
		return err
	}

	return nil
}

func (pg *UserPostgres) GetByEmail(email string) (*User, error) {
	querySting := `SELECT * FROM users WHERE email = $1 LIMIT 1;`
	row := pg.db.QueryRow(querySting, email)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var u User
	err := row.Scan(&u.Id, &u.Email, &u.HashedPassword, &u.Name, &u.Role)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		break
	default:
		return nil, err
	}

	return &u, nil
}

func (pg *UserPostgres) GetById(id int) (*User, error) {
	querySting := `SELECT * FROM users WHERE id = $1 LIMIT 1;`
	row := pg.db.QueryRow(querySting, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var u User
	err := row.Scan(&u.Id, &u.Email, &u.HashedPassword, &u.Name, &u.Role)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		break
	default:
		return nil, err
	}

	return &u, nil
}
