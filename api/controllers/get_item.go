package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kkrisstoff/go-server/models"
)

func GetItemById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		//TODO: write errors app_error.WriteError(w, err)
	}

	if r.Method == "GET" {
		id := r.Form["id"]

		if len(id) < 1 {
			w.Write([]byte("{\"error\":\"bad request\",\"message\":\"don't have id in request\"}"))
			return
		}

		item, err := models.ItemsStoreMapped.GetItemByID(id[0])
		if err != nil {
			// handle error
			return
		}

		fmt.Println(item)
		idStr := strconv.Itoa(item.ID)
		w.Write([]byte("{\"id\":" + idStr + ", \"message\":" + item.Message + "}"))
	}

}
