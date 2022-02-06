package routes

import (
	"github.com/alvinarthas/myrepo/business/controller"
	"github.com/go-chi/chi"
)

func DummyRoutes() *chi.Mux {

	routes := chi.NewRouter()
	routes.Get("/", controller.GetDummyList)
	routes.Post("/", controller.CreateDummy)

	return routes

}
