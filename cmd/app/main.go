package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pttrulez/investor-go/internal/config"
	postgres "github.com/pttrulez/investor-go/internal/repository"
	"github.com/pttrulez/investor-go/internal/router"
)

func main() {

	cfg := config.MustLoad()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	repository, err := postgres.NewPostgresRepo(cfg.Pg)
	if err != nil {
		panic("Failed to initialize postgres repository: " + err.Error())
	}
	router.Init(r, repository, cfg)

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
