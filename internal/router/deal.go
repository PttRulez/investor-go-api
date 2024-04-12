package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"github.com/pttrulez/investor-go/internal/lib/api/response"
	"github.com/pttrulez/investor-go/internal/lib/helpers"
	"github.com/pttrulez/investor-go/internal/types"
)

func DealRoutes(r chi.Router, repo *types.Repository, tokenAuth *jwtauth.JWTAuth) {

	r.Route("/deal", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		// Create new deal
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var dealData types.CreateDeal
			if err := json.NewDecoder(r.Body).Decode(&dealData); err != nil {
				response.WriteJSON(w, http.StatusBadRequest, response.Error(err.Error()))
				return
			}

			// Validate request fields
			validate := validator.New()
			validate.RegisterValidation("is-exchange", helpers.ValidateExchange)
			validate.RegisterValidation("isodate", helpers.IsISO8601Date)
			validate.RegisterValidation("price", helpers.ValidatePrice)
			validate.RegisterValidation("securityType", helpers.ValidateSecurityType)
			validate.RegisterValidation("dealType", helpers.ValidateDealType)

			if err := validate.Struct(dealData); err != nil {
				validateErr := err.(validator.ValidationErrors)
				response.WriteValidationErrorsJSON(w, http.StatusBadRequest, validateErr)
				return
			}
			fmt.Printf("%+v\n", dealData)
			// Create new Deal
			newDeal, err := repo.Deal.CreateDeal(&dealData)
			if err != nil {
				response.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
			}
			response.WriteCreatedJSON(w, newDeal)
		})
	})
}
