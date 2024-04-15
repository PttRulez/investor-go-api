package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/pttrulez/investor-go/internal/services/moex"
	validatorservice "github.com/pttrulez/investor-go/internal/services/validator"
	"github.com/pttrulez/investor-go/internal/types"
)

type ServiceContainer struct {
	Moex      *moex.MoexService
	Validator *validator.Validate
}

func NewServiceContainer(repo *types.Repository) *ServiceContainer {
	return &ServiceContainer{
		Moex:      moex.NewMoexService(repo),
		Validator: validatorservice.NewValidator(),
	}
}
