package services

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/pttrulez/investor-go/internal/types"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("dealType", ValidateDealType)
	validate.RegisterValidation("is-exchange", ValidateExchange)
	validate.RegisterValidation("isodate", IsISO8601Date)
	validate.RegisterValidation("opinionType", ValidatePrice)
	validate.RegisterValidation("price", ValidatePrice)
	validate.RegisterValidation("securityType", ValidateSecurityType)
	return validate
}

func ValidateExchange(fl validator.FieldLevel) bool {
	return fl.Field().Interface().(types.Exchange).Validate()
}

func ValidateSecurityType(fl validator.FieldLevel) bool {
	return fl.Field().Interface().(types.SecurityType).Validate()
}

func ValidateDealType(fl validator.FieldLevel) bool {
	return fl.Field().Interface().(types.DealType).Validate()
}

func IsISO8601Date(fl validator.FieldLevel) bool {
	ISO8601DateRegexString := "^(?:[1-9]\\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d(?:\\.\\d{1,9})?(?:Z|[+-][01]\\d:[0-5]\\d)$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(fl.Field().String())
}

func ValidateOpinion(fl validator.FieldLevel) bool {
	return fl.Field().Interface().(types.OpinionType).Validate()
}

func ValidatePrice(fl validator.FieldLevel) bool {
	priceFloat, ok := fl.Field().Interface().(float64)
	if !ok {
		return false
	}
	match, _ := regexp.MatchString(`^\d+$|^\d+\.\d{1,2}$`, strconv.FormatFloat(priceFloat, 'f', -1, 64))
	return match
}
