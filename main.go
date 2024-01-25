package main

import (
	"fmt"
	"gorest/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	//routes
	routes.UserRoutes(router)
	routes.MessageRoutes(router)

	fmt.Println("Server started at 🚀 http://localhost:3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
