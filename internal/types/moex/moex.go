package tmoex

type MoexApiResponseSecurityInfo struct {
	Description struct {
		Data [][]string `json:"data"`
	}
	Boards struct {
		Data [][4]any `json:"data"`
	}
}

type MoexSecurity struct {
	Board     Board  `json:"board" db:"board"`
	Engine    Engine `json:"engine" db:"engine"`
	Market    Market `json:"market" db:"market"`
	Name      string `json:"name" db:"name"`
	ShortName string `json:"shortname" db:"shortname"`
	Ticker    string `json:"ticker" db:"ticker"`
}

type MoexShareInfo struct {
	MoexSecurity
}

type MoexBondInfo struct {
	MoexSecurity
}

type Share struct {
	Board     Board  `json:"board" db:"board"`
	Engine    Engine `json:"engine" db:"engine"`
	Market    Market `json:"market" db:"market"`
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	ShortName string `json:"shortname" db:"shortname"`
	Ticker    string `json:"ticker" db:"ticker"`
}

type Bond struct {
	Board     Board  `json:"board" db:"board"`
	Engine    Engine `json:"engine" db:"engine"`
	Market    Market `json:"market" db:"market"`
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	ShortName string `json:"shortname" db:"shortname"`
	Ticker    string `json:"ticker" db:"ticker"`
}
