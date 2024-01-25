package routes

import (
	"gorest/controllers"

	"github.com/go-chi/chi/v5"
)

func MessageRoutes(router chi.Router) {
	router.Route("/message", func(r chi.Router) {
		r.Get("/", controllers.HandleGetAllMessages)
	})
}
