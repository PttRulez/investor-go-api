package types

type User struct {
	Email          string `json:"email" db:"email"`
	HashedPassword string `json:"-" db:"hashed_password"`
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Role           Role   `json:"role" db:"role"`
}
