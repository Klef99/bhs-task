// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/Klef99/bhs-task/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewRouter -.
// Swagger spec:
// @title       Bhs-task
// @description A test assignment for a backend developer at BHS
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler chi.Router, l logger.Interface) {
	// Options
	handler.Use(middleware.Logger)
	handler.Use(middleware.Recoverer)

	// Swagger
	// swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	// handler.GET("/swagger/*any", swaggerHandler)
	// K8s probe
	handler.Get("/healthz", func(resp http.ResponseWriter, req *http.Request) { resp.WriteHeader(http.StatusOK) })

	// Routers
	// h := handler.Group("/v1")
	// {
	// 	newTranslationRoutes(h, t, l)
	// }
}
