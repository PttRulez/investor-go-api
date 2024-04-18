package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pttrulez/investor-go/internal/types"
	tmoex "github.com/pttrulez/investor-go/internal/types/moex"
)

type IssApi struct {
	baseUrl string
	client  *http.Client
}

func CreateISSApiClient() *IssApi {
	return &IssApi{
		baseUrl: "https://iss.moex.com/iss",
		client:  &http.Client{},
	}
}

func (api *IssApi) GetSecurityInfoByTicker(ticker string) (*tmoex.Share, error) {
	uri := fmt.Sprintf("%s/securities/%s.json", api.baseUrl, ticker)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		fmt.Println("[Api GetSecurityByTicker http.NewRequest]:", err)
		return nil, err
	}

	// фильтруем только то что нам нужно
	params := url.Values{}
	params.Add("iss.meta", "off")
	params.Add("description.columns", "name,value")
	params.Add("boards.columns", "boardid,market,engine,is_primary")
	req.URL.RawQuery = params.Encode()

	resp, err := api.client.Do(req)
	if err != nil {
		fmt.Println("[Api GetSecurityByTicker api.client.Get(moexSecurityUrl)]:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[Api GetSecurityByTicker body ReadAll]:", err)
		return nil, err
	}

	data := &types.MoexApiResponseSecurityInfo{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("[Api GetSecurityByTicker MoexApiResponseSecurityInfo json.Unmarshal]:", err)
		return nil, err
	}

	var (
		name, shortname string
		board           tmoex.Board
		market          tmoex.Market
		engine          tmoex.Engine
		// ok              bool
	)

	for _, item := range data.Description.Data {
		if item[0] == "NAME" {
			name = item[1]
		}
		if item[0] == "SHORTNAME" {
			shortname = item[1]
		}
	}

	var boardData [4]any
	for _, item := range data.Boards.Data {
		if item[3].(float64) == 1 {
			boardData = item
		}
	}
	board = tmoex.Board(boardData[0].(string))
	market = tmoex.Market(boardData[1].(string))
	engine = tmoex.Engine(boardData[2].(string))

	return &tmoex.Share{
		Board:     board,
		Engine:    engine,
		Market:    market,
		Name:      name,
		ShortName: shortname,
		Ticker:    ticker,
	}, nil
}

func (api *IssApi) GetStocksCurrentPrices(market tmoex.Market, tickers []string) (*types.MoexApiResponseCurrentPrices, error) {
	uri := fmt.Sprintf("%s/engines/stock/markets/%s/securities.json", api.baseUrl, market)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("iss.meta", "off")
	params.Add("securities", url.QueryEscape(strings.Join(tickers, ",")))
	params.Add("securities.columns", "SECID,BOARDID,PREVPRICE")
	req.URL.RawQuery = params.Encode()

	fmt.Println("[GetStocksCurrentPrices] RequestURI:", req.URL, req.RequestURI)
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[Api GetStocksCurrentPrices body ReadAll]:", err)
		return nil, err
	}

	data := &types.MoexApiResponseCurrentPrices{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("[Api GetStocksCurrentPrices MoexApiResponseCurrentPrices json.Unmarshal]:", err)
		return nil, err
	}

	return data, nil
}
