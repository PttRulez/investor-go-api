package services

import (
	"database/sql"

	"github.com/pttrulez/investor-go/internal/types"
	tmoex "github.com/pttrulez/investor-go/internal/types/moex"
)

type ShareService struct {
	moexApi *IssApi
	repo    *types.Repository
}

func (s *ShareService) GetByTicker(ticker string) (*tmoex.Share, error) {
	// Проверяем есть ли уже в базе
	share, err := s.repo.Moex.Shares.GetByTicker(ticker)
	if err == sql.ErrNoRows {
		// если нет то создаем и возвращаем
		// запрос на информацию по бумаге из апишки московской биржи
		moexShare, err := s.moexApi.GetSecurityInfoByTicker(ticker)
		if err != nil {
			return nil, err
		}

		// сохраняем в бд
		err = s.repo.Moex.Shares.Insert(moexShare)
		if err != nil {
			return nil, err
		}

		// ищем её же в бд
		share, err = s.repo.Moex.Shares.GetByTicker(ticker)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	// если уже была в базе, то просто возвращаем
	return share, nil
}

func NewShareService(repo *types.Repository) *ShareService {
	return &ShareService{
		moexApi: CreateISSApiClient(),
		repo:    repo,
	}
}
