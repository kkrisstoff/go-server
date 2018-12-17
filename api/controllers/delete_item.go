package controllers

import (
	"fmt"
	"net/http"

	"github.com/kkrisstoff/go-server/models"
)

//DeleteItemById delete item by id
func DeleteItemById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		//TODO: write errors app_error.WriteError(w, err)
	}
	if r.Method == "DELETE" {
		id := r.Form["id"]
		models.ItemsStoreMapped.DeleteItemByID(id[0])
	} else {
		fmt.Println("Unexpected method: ", r.Method)
	}
}
