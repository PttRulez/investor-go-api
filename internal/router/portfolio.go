package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"github.com/pttrulez/investor-go/internal/lib/api/response"
	"github.com/pttrulez/investor-go/internal/types"
)

func PortfolioRoutes(r chi.Router, repo *types.Repository, tokenAuth *jwtauth.JWTAuth) {

	r.Route("/portfolio", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		// List of portfolios
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			portfolios, err := repo.Portfolio.GetPortfolioListByUserId(getUserIdFrowJwt(r))
			if err != nil {
				response.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
				return
			}
			response.WriteOKJSON(w, portfolios)
		})

		// Create portfolio
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var portfolioData types.Portfolio
			if err := json.NewDecoder(r.Body).Decode(&portfolioData); err != nil {
				response.WriteJSON(w, http.StatusBadRequest, response.Error(err.Error()))
				return
			}

			// Validate request fields
			if err := validator.New().Struct(portfolioData); err != nil {
				validateErr := err.(validator.ValidationErrors)
				response.WriteValidationErrorsJSON(w, http.StatusBadRequest, validateErr)
				return
			}
			_, claims, _ := jwtauth.FromContext(r.Context())
			portfolioData.UserId = int(claims["id"].(float64))

			// Create new Portfolio
			newPortfolio, err := repo.Portfolio.CreatePortfolio(portfolioData)
			if err != nil {
				response.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
			}
			response.WriteOKJSON(w, newPortfolio)
		})

		r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
			//_, claims, _ := jwtauth.FromContext(r.Context())

			var portfolioData types.PortfolioUpdate
			err := json.NewDecoder(r.Body).Decode(&portfolioData)
			if err != nil {
				fmt.Println("Patch portfolioData err", err)
			}

			// Update Portfolio
			updatedPortfolio, err := repo.Portfolio.UpdatePortfolio(chi.URLParam(r, "id"), portfolioData)
			if err != nil {
				response.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
			}
			response.WriteOKJSON(w, updatedPortfolio)
		})

	})
}
