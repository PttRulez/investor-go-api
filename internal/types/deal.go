package types

import "time"

type Deal struct {
  amount int `json:"amount" db:"amount"`
  date   time.Time `json:"date" db:"date"`
}