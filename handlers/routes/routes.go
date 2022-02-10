package routes

import (
	"net/http"

	"github.com/alvinarthas/myrepo/utils/health"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// GetRouter returns a Mux object that implements the Router interface.
func GetRouter() *chi.Mux {

	router := chi.NewRouter()

	setupMiddleware(router)

	router.Get("/health", health.HealthStatus)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world, welcome to my repo"))
	})

	router.Mount("/dummies", DummyRoutes())

	return router

}

// Setup global middlewares
func setupMiddleware(router *chi.Mux) {

	// Add a own middleware.
	router.Use(
		setCorsOptions(),
		middleware.Recoverer,
		middleware.Logger,
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
