package moex

type Share struct {
	Board     Board  `json:"board" db:"board"`
	Engine    Engine `json:"engine" db:"engine"`
	Market    Market `json:"market" db:"market"`
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	ShortName string `json:"shortname" db:"shortname"`
	Ticker    string `json:"ticker" db:"ticker"`
}
