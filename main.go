package main

import (
	"fmt"
	"gorest/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type person struct {
	name string
	age  int
}

func main() {
	router := chi.NewRouter()
	routes.UserRoutes(router)
	fmt.Println("Server started at 🚀 http://localhost:3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
