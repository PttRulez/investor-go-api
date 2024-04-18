package types

type Portfolio struct {
	Compound bool   `json:"compound" db:"compound" validate:"required,bool"`
	Id       int    `json:"id" db:"id"  validate:"required,int"`
	Name     string `json:"name" db:"name"  validate:"required,string"`
	UserId   int    `json:"-" db:"user_id"`
}

type PortfolioUpdate struct {
	Compound *bool   `json:"compound,omitempty" db:"compound" validate:"bool"`
	Name     *string `json:"name,omitempty" db:"name" validate:"string"`
	Id       int     `json:"id" db:"id"  validate:"required,int"`
}
