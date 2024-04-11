package types

type Portfolio struct {
	Compound bool   `json:"compound" db:"compound"`
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	UserId   int    `json:"-" db:"user_id"`
}

type PortfolioUpdate struct {
	Compound *bool   `json:"compound,omitempty" db:"compound"`
	Name     *string `json:"name,omitempty" db:"name"`
}
