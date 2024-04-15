package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pttrulez/investor-go/internal/config"
	postgres "github.com/pttrulez/investor-go/internal/repository/postgres"
	"github.com/pttrulez/investor-go/internal/router"
	"github.com/pttrulez/investor-go/internal/services"
)

func main() {
	cfg := config.MustLoad()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	repository, err := postgres.NewPostgresRepo(cfg.Pg)
	services := services.NewServiceContainer(repository)

	if err != nil {
		panic("Failed to initialize postgres repository: " + err.Error())
	}
	router.Init(r, repository, services, cfg)

	logger := slog.Default()
	logger.Info(fmt.Sprintf("Listening on port %v", cfg.ApiPort))

	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%v", cfg.ApiPort),
		Handler: r,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
		logger.Error(err.Error())
	}
}
