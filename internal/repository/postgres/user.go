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

func (pg *UserPostgres) CreateUser(u User) (*User, error) {
	querySting := "INSERT INTO users (email, hashed_password, name, role) VALUES ($1, $2, $3, $4) RETURNING *;"
	row := pg.db.QueryRow(querySting, u.Email, u.HashedPassword, u.Name, u.Role)
	if row.Err() != nil {
		fmt.Println("Failed to execute query:", row.Err())
		return nil, row.Err()
	}

	var newUser User
	if err := row.Scan(&newUser.Id, &newUser.Email, &newUser.HashedPassword, &newUser.Name, &newUser.Role); err != nil {
		fmt.Println("Failed to scan:", err)
		return nil, err
	}

	return &newUser, nil
}

func (pg *UserPostgres) GetUserByEmail(email string) (*User, error) {
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

func (r *UserPostgres) GetUserById(id int) (*User, error) {
	return &User{}, nil
}
