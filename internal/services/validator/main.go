package validatorservice

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("is-exchange", ValidateExchange)
	validate.RegisterValidation("isodate", IsISO8601Date)
	validate.RegisterValidation("price", ValidatePrice)
	validate.RegisterValidation("securityType", ValidateSecurityType)
	validate.RegisterValidation("dealType", ValidateDealType)
	return validate
}
