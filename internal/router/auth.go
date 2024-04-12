package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pttrulez/investor-go/internal/lib/api/response"
	"github.com/pttrulez/investor-go/internal/types"
	"golang.org/x/crypto/bcrypt"
)

func AuthRoutes(r chi.Router, repo *types.Repository, tokenAuth *jwtauth.JWTAuth) {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var registerData types.RegisterUser

		// Return error if json couldn't be decoded
		if err := json.NewDecoder(r.Body).Decode(&registerData); err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.Error(err.Error()))
			return
		}

		// Validate request fields
		if err := validator.New().Struct(registerData); err != nil {
			validateErr := err.(validator.ValidationErrors)
			response.WriteValidationErrorsJSON(w, http.StatusBadRequest, validateErr)
			return
		}

		// Check if user with this email already exists
		existingUser, err := repo.User.GetUserByEmail(registerData.Email)
		if existingUser != nil {
			response.WriteErrorJSON(w, http.StatusBadRequest, "Пользователь с таким email уже существует")
			return
		} else if err != nil {
			response.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		encpw, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
		if err != nil {
			response.WriteErrorJSON(w, http.StatusInternalServerError, ":(")
		}

		if registerData.Role == "" {
			registerData.Role = types.Investor
		}

		// Creating new user
		newUser, err := repo.User.CreateUser(types.User{
			Email:          registerData.Email,
			Name:           registerData.Name,
			HashedPassword: string(encpw),
			Role:           registerData.Role,
		})
		if err != nil {
			response.WriteErrorJSON(w, http.StatusInternalServerError, "Failed to create new user")
		}

		response.WriteJSON(w, http.StatusOK, newUser)
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var loginData types.LoginUser

		// Return error if json couldn't be decoded
		if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.Error(err.Error()))
			return
		}

		user, err := repo.User.GetUserByEmail(loginData.Email)
		fmt.Println(loginData.Email)
		if err != nil {
			response.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		} else if user == nil {
			response.WriteErrorJSON(w, http.StatusBadRequest, "Пользователя с таким email не существует")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginData.Password))
		if err != nil {
			response.WriteErrorJSON(w, http.StatusUnauthorized, "Неверные данные")
			return
		}

		claims := jwt.MapClaims{
			"id":    user.Id,
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		}
		jwtauth.SetExpiry(claims, time.Now().Add(time.Hour*6))
		_, tokenString, _ := tokenAuth.Encode(claims)

		response.WriteOKJSON(w, map[string]string{"token": tokenString})
	})
}
