package moex

import (
	"database/sql"

	moexapi "github.com/pttrulez/investor-go/internal/services/moex-api"
	"github.com/pttrulez/investor-go/internal/types"
	tmoex "github.com/pttrulez/investor-go/internal/types/moex"
)

type ShareService struct {
	moexApi *moexapi.IssApi
	repo    *types.Repository
}

func (s *ShareService) GetByTicker(ticker string) (*tmoex.Share, error) {
	// Проверяем есть ли уже в базе
	share, err := s.repo.Moex.Shares.GetByTicker(ticker)
	if err == sql.ErrNoRows {
		// если нет то создаем и возвращаем
		return s.createByTicker(ticker)
	} else if err != nil {
		return nil, err
	}

	// если уже была в базе, то просто возвращаем
	return share, nil
}

func (s *ShareService) createByTicker(ticker string) (*tmoex.Share, error) {
	moexShare, err := s.moexApi.GetSecurityInfoByTicker(ticker)
	if err != nil {
		return nil, err
	}

	moexShare, err = s.repo.Moex.Shares.Create(moexShare)
	if err != nil {
		return nil, err
	}
	return moexShare, nil
}

func NewShareService(repo *types.Repository) *ShareService {
	return &ShareService{
		moexApi: moexapi.CreateISSApiClient(),
		repo:    repo,
	}
}
