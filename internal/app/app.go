package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Klef99/bhs-task/config"
	v1 "github.com/Klef99/bhs-task/internal/controller/http/v1"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/Klef99/bhs-task/internal/usecase/repo"
	"github.com/Klef99/bhs-task/pkg/hasher"
	"github.com/Klef99/bhs-task/pkg/httpserver"
	"github.com/Klef99/bhs-task/pkg/jwtgenerator"
	"github.com/Klef99/bhs-task/pkg/logger"
	"github.com/Klef99/bhs-task/pkg/postgres"
	"github.com/go-chi/chi/v5"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Jwt generator
	jtg, err := jwtgenerator.New(cfg.Jwt.Secret,
		jwtgenerator.TokenNbf(time.Second*time.Duration(cfg.Jwt.Nbf)),
		jwtgenerator.TokenExp(time.Second*time.Duration(cfg.Jwt.Exp)),
	)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - jwtgenerator.New: %w", err))
	}

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	UserUseCase := usecase.NewUserUseCase(
		repo.NewUserRepository(pg, hasher.NewHasher()),
	)
	AssetUseCase := usecase.NewAssetUseCase(
		repo.NewAssetRepository(pg),
	)
	// HTTP Server
	handler := chi.NewRouter()
	v1.NewRouter(handler, l, UserUseCase, AssetUseCase, jtg, cfg.HTTP.Swagger)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
