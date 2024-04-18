package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"github.com/pttrulez/investor-go/internal/lib/api/response"
	"github.com/pttrulez/investor-go/internal/services"
	"github.com/pttrulez/investor-go/internal/types"
)

func ExpertsRoutes(r chi.Router, repo *types.Repository, services *services.ServiceContainer, tokenAuth *jwtauth.JWTAuth) {
	r.Route("/expert", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var expert types.Expert
			if err := json.NewDecoder(r.Body).Decode(&expert); err != nil {
				response.WriteJSON(w, http.StatusBadRequest, response.Error(err.Error()))
				return
			}

			// Validate request fields
			if err := services.Validator.Struct(expert); err != nil {
				validateErr := err.(validator.ValidationErrors)
				response.WriteValidationErrorsJSON(w, http.StatusBadRequest, validateErr)
				return
			}

			// Expert must be created by user
			_, claims, _ := jwtauth.FromContext(r.Context())
			expert.UserId = int(claims["id"].(float64))

			// Save new Expert in DB
			err := repo.Expert.Insert(expert)
			if err != nil {
				response.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
			}
			w.WriteHeader(http.StatusCreated)
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			experts, err := repo.Expert.GetListByUserId(getUserIdFrowJwt(r))
			if err != nil {
				response.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
				return
			}
			response.WriteOKJSON(w, experts)
		})
	})
}
