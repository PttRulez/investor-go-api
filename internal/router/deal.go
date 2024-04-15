package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"github.com/pttrulez/investor-go/internal/lib/api/response"
	"github.com/pttrulez/investor-go/internal/services"
	"github.com/pttrulez/investor-go/internal/types"
)

func DealRoutes(r chi.Router, repo *types.Repository, services *services.ServiceContainer, tokenAuth *jwtauth.JWTAuth) {
	r.Route("/deal", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		// Create new deal
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			// Анмаршалим данные
			var dealData types.RepoCreateDeal
			if err := json.NewDecoder(r.Body).Decode(&dealData); err != nil {
				response.WriteJSON(w, http.StatusBadRequest, response.Error(err.Error()))
				return
			}

			// Валидация пришедших данных
			if err := services.Validator.Struct(dealData); err != nil {
				validateErr := err.(validator.ValidationErrors)
				response.WriteValidationErrorsJSON(w, http.StatusBadRequest, validateErr)
				return
			}

			// Проверяем есть ли в бд такакя бумага. Если нет, то сервис создаст её.
			// Также если тикер в dealData будет неправильный, то будет ошибка
			if dealData.SecurityType == types.Share && dealData.Exchange == types.Moex {
				_, err := services.Moex.Shares.GetByTicker(
					dealData.Ticker,
				)
				if err != nil {
					fmt.Println("[/deal] POST handler: ", err)
					response.WriteErrorJSON(w, http.StatusInternalServerError, "Что-то пошло не так")
					return
				}
			}

			// Создаем сделку
			newDeal, err := repo.Deal.CreateDeal(&dealData)
			if err != nil {
				response.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
			}



			response.WriteCreatedJSON(w, newDeal)
		})
	})
}
