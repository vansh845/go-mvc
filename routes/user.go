package routes

import (
	"gorest/controllers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(router chi.Router) {

	router.Route("/users", func(r chi.Router) {
		r.Get("/", controllers.HandleGetAllUserRoute)
		r.Get("/{userid}", controllers.HandleGetUserbyId)
	})

}
