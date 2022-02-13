package routes

import (
	"net/http"

	_ "github.com/alvinarthas/myrepo/docs"
	"github.com/alvinarthas/myrepo/utils/health"
	"github.com/alvinarthas/myrepo/utils/middleware"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// GetRouter returns a Mux object that implements the Router interface.

func GetRouter() *chi.Mux {

	router := chi.NewRouter()

	setupMiddleware(router)

	router.Get("/health", health.HealthStatus)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:80/swagger/doc.json"), //The url pointing to API definition
	))

	apiV1 := router.Group(nil)
	apiV1.Use(middleware.AppAuthorization)
	apiV1.Route("/api/v1", func(router chi.Router) {
		router.Mount("/dummies", DummyRoutes())
	})

	return router

}

// Setup global middlewares
func setupMiddleware(router *chi.Mux) {

	// Add a own middleware.
	router.Use(
		setCorsOptions(),
		chiMiddleware.Recoverer,
		chiMiddleware.Logger,
	)

}

func setCorsOptions() func(next http.Handler) http.Handler {

	// Basic CORS (cross-origin-resource-sharing)
	// for detail about CORS, see: https://github.com/rs/cors
	var corsOptions = cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Api-Key", "trace_id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // results of a preflight request can be cached in 5 minutes.
	})

	return corsOptions

}
