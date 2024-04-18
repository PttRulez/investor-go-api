package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/pttrulez/investor-go/internal/types"
)

type ServiceContainer struct {
	Deal      *DealService
	Moex      *MoexService
	Position  *PositionService
	Validator *validator.Validate
}

func NewServiceContainer(repo *types.Repository) *ServiceContainer {
	container := &ServiceContainer{}
	container.Deal = NewDealService(repo, container)
	container.Position = NewPositionService(repo)
	container.Moex = NewMoexService(repo)
	container.Validator = NewValidator()

	return container
}
