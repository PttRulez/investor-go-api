package services

import (
	"fmt"

	"github.com/pttrulez/investor-go/internal/types"
)

type DealService struct {
	repo     *types.Repository
	services *ServiceContainer
}

func NewDealService(repo *types.Repository, services *ServiceContainer) *DealService {
	return &DealService{
		repo:     repo,
		services: services,
	}
}

func (s *DealService) CreateDeal(dealData types.RepoCreateDeal) error {
	// Проверяем есть ли в бд такакя бумага. Если нет, то сервис создаст её.
	// TODO: Также если тикер в dealData будет неправильный, то будет ошибка
	if dealData.SecurityType == types.Share && dealData.Exchange == types.Moex {
		share, err := s.services.Moex.Shares.GetByTicker(dealData.Ticker)
		if err != nil {
			return err
		}
		fmt.Printf("Найдена акция: %+v\n", share)
		dealData.SecurityId = share.Id
	}

	// Создаем сделку
	err := s.repo.Deal.Insert(&dealData)
	if err != nil {
		return err
	}

	// Обновляем позицию
	err = s.services.Position.UpdatePositionInDB(dealData.Exchange, dealData.PortfolioId, dealData.SecurityType, dealData.SecurityId)

	return err
}
