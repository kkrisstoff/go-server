package controllers

import (
	"encoding/json"
	"github.com/kkrisstoff/go-server/models"
	"net/http"
)

// GetItems get all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == "GET" {
		allItems := models.ItemsStoreMapped.GetItems()

		b, err := json.Marshal(allItems)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(b))
		return
	}
}
