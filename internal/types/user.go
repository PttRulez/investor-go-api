package types

type User struct {
	Email          string `json:"email" db:"email"`
	HashedPassword string `json:"-" db:"hashed_password"`
	Id             int  `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Role           Role   `json:"role" db:"role"`
}

type Role string

const (
	ADMIN    Role = "ADMIN"
	INVESTOR Role = "INVESTOR"
)

type UserRepository interface {
	CreateUser(u User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
}
