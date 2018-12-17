package controllers

import (
	"fmt"
	"net/http"

	"github.com/kkrisstoff/go-server/models"
)

//DeleteItemById delete item by id
func DeleteItemById(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "DELETE" {
		id := r.Form["id"]
		models.ItemsStoreMapped.DeleteItemByID(id[0])
	} else {
		fmt.Println("Unexpected method: ", r.Method)
	}
}
