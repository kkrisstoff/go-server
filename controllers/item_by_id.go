package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kkrisstoff/go-server/models"
)

// Item handelt item by ID TODO: refactor
func Item(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == "GET" {
		idStr := r.URL.Query().Get("id")
		if len(idStr) == 0 {
			http.Error(w, errors.New("ID is not exist").Error(), http.StatusInternalServerError)
		}
		item := models.ItemsStoreMapped.GetItemByID(idStr)
		b, err := json.Marshal(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(b))
	}

	if r.Method == "DELETE" {
		idStr := r.URL.Query().Get("id")
		if len(idStr) == 0 {
			http.Error(w, errors.New("ID is not exist").Error(), http.StatusInternalServerError)
		}
		models.ItemsStoreMapped.DeleteItemByID(idStr)
		w.Write([]byte(fmt.Sprintf("{\"id\": \"%s\"}", idStr)))
	}

	if r.Method == "PUT" {
		var requestItem struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Message string `json:"message"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		item := models.ItemsStoreMapped.UpdateItem(requestItem.ID, requestItem.Name, requestItem.Message)
		b, err := json.Marshal(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(b))
	}
}
