package v1

import (
	"net/http"

	_ "github.com/Klef99/bhs-task/docs"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/Klef99/bhs-task/pkg/jwtgenerator"
	"github.com/Klef99/bhs-task/pkg/logger"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter -.
// Swagger spec:
// @title       Bhs-task
// @description A test assignment for a backend developer at BHS
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
// @schemes     http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func NewRouter(handler chi.Router, l logger.Interface, t usecase.User, a usecase.Asset, jwt jwtgenerator.Interface, enableSwagger bool) {
	// Options
	handler.Use(middleware.Logger)
	handler.Use(middleware.Recoverer)

	// K8s probe
	handler.Get("/healthz", func(resp http.ResponseWriter, req *http.Request) { resp.WriteHeader(http.StatusOK) })

	// Swagger
	if enableSwagger {
		handler.Get("/swagger/*", httpSwagger.WrapHandler)
	}
	// v1 api declaration
	r := chi.NewRouter()
	NewUserRoutes(r, t, l, jwt)
	NewAssetRoutes(r, a, l, jwt)
	handler.Mount("/v1", r)
}
