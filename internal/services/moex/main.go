package moex

import (
	"github.com/pttrulez/investor-go/internal/types"
)

type MoexService struct {
	Shares ShareService
}

func NewMoexService(repo *types.Repository) *MoexService {
	return &MoexService{
		Shares: *NewShareService(repo),
	}
}
