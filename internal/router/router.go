package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/pttrulez/investor-go/internal/config"
	"github.com/pttrulez/investor-go/internal/types"
)

func Init(r chi.Router, repo *types.Repository, cfg *config.Config) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	r.Route("/", func(r chi.Router) {
		AuthRoutes(r, repo, tokenAuth)
		PortfolioRoutes(r, repo, tokenAuth)
		DealRoutes(r, repo, tokenAuth)
	})
}
