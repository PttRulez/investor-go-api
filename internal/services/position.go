package services

import (
	"fmt"
	"math"

	"github.com/pttrulez/investor-go/internal/types"
)

type PositionService struct {
	repo *types.Repository
}

func NewPositionService(repo *types.Repository) *PositionService {
	return &PositionService{repo: repo}
}

func (s *PositionService) UpdatePositionInDB(exchange types.Exchange, portfolioId int, securityType types.SecurityType,
	securityId int) error {
	allDeals, err := s.repo.Deal.GetDealsListForSecurity(exchange, portfolioId, securityType, securityId)
	if err != nil {
		return err
	}

	var position *types.Position
	oldPosition, err := s.repo.Position.Get(exchange, portfolioId, securityId, securityType)
	if err != nil {
		return err
	}
	if oldPosition != nil {
		position = oldPosition
	} else {
		position = &types.Position{
			Exchange:     exchange,
			PortfolioId:  portfolioId,
			SecurityId:   securityId,
			SecurityType: securityType,
		}
	}

	// Calculate and add to position Amount
	var amount int
	var totalAmount int
	for _, deal := range allDeals {
		amount = deal.Amount
		if deal.Type == types.Sell {
			amount = -amount
		}
		totalAmount += amount
	}
	position.Amount = totalAmount

	// Calculate and add AveragePrice to position
	m := make(map[float64]int)
	left := position.Amount
	for _, deal := range allDeals {
		if deal.Type == types.Sell {
			continue
		}
		if left > deal.Amount {
			m[deal.Price] += deal.Amount
			left -= deal.Amount
		} else {
			m[deal.Price] += left
			break
		}
	}
	fmt.Printf("prices m: %+v\n", m)
	var avPrice float64 = 0
	for price, amount := range m {
		avPrice += float64(amount) / float64(position.Amount) * price
	}
	position.AveragePrice = math.Floor(avPrice*100) / 100

	if position.Id == 0 {
		err = s.repo.Position.Insert(position)
		if err != nil {
			return err
		}
	} else {
		err = s.repo.Position.Update(position)
		if err != nil {
			return err
		}
	}

	return nil
}
