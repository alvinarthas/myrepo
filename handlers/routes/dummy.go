package routes

import (
	"github.com/alvinarthas/myrepo/business/controller"
	"github.com/go-chi/chi"
)

func DummyRoutes() *chi.Mux {

	routes := chi.NewRouter()
	routes.Get("/", controller.GetDummyList)
	routes.Post("/", controller.CreateDummy)
	routes.Get("/{id}", controller.GetDummyByID)
	routes.Put("/{id}", controller.UpdateDummy)
	routes.Delete("/{id}", controller.DeleteDummy)

	return routes

}
