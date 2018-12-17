package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kkrisstoff/go-server/api/models"
)

// GetItemById get items by id
func GetItemById(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// fmt.Println(r.Form["id"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }

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
