package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kkrisstoff/go-server/models"
	"net/http"
)

type reqItem struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// TODO: reuse this type
type message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == "POST" {
		var requestItem reqItem
		err := json.NewDecoder(r.Body).Decode(&requestItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newItem := models.ItemsStoreMapped.AddItem(requestItem.Name, requestItem.Message)
		body, err := json.Marshal(message{
			newItem.ID,
			newItem.Name,
			newItem.Message,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("Item %v has been added.\n", body)
		w.Write([]byte(body))
	}

	if r.Method == "GET" {
		fmt.Fprintf(w, "Use POST for adding new item")
	}
}
