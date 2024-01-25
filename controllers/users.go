package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Person struct {
	Name string
	Age  int
	Job  string
	Id   string
}

var people []Person = []Person{
	{"John Doe", 30, "Software Engineer", "j30"},
	{"Jane Smith", 25, "Data Scientist", "j25"},
	{"Bob Johnson", 35, "Product Manager", "b35"},
}

func HandleGetUserbyId(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userid")
	fmt.Println(userId)
	for i := 0; i < len(people); i++ {
		if people[i].Id == userId {
			user, err := json.Marshal(people[i])
			if err != nil {
				http.Error(w, err.Error(), 400)
			}
			w.Write(user)
			return
		}
	}
	w.Write([]byte("Nothing found"))
}

func HandleGetAllUserRoute(w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.Marshal(people)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	w.Write(jsonData)
}
