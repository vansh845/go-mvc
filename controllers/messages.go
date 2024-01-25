package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Sender  string
	Content string
}

func HandleGetAllMessages(w http.ResponseWriter, r *http.Request) {
	jsonData := Message{Sender: "vansh", Content: "hi"}
	s, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(s)
}
