package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kkrisstoff/go-server/models"
)

func GetItemByID(w http.ResponseWriter, r *http.Request) {
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

}
