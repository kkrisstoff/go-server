package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kkrisstoff/go-server/models"
	"net/http"
	"strconv"
)

type reqItem struct {
	Message string `json:"message"`
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var requestItem reqItem
		err := json.NewDecoder(r.Body).Decode(&requestItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer r.Body.Close()

		newItem := models.ItemsStoreMapped.AddItem(requestItem.Message)
		fmt.Printf("Item %v has been added.\n", newItem)
		idStr := strconv.Itoa(newItem.ID)

		var b bytes.Buffer

		b.WriteString(fmt.Sprintf(`{"id":%s, "message":%s}`, idStr, newItem.Message))
	}

	if r.Method == "GET" {
		fmt.Fprintf(w, "Use POST for adding items") // send data to client side
	}
}
