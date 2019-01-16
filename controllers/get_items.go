package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kkrisstoff/go-server/models"
	"net/http"
)

// GetItems get items
func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/getItems")
	fmt.Println("method:", r.Method) //get request method

	allItems := models.ItemsStoreMapped.GetItems()
	fmt.Println("All Items:", allItems)

	b, err := json.Marshal(allItems)
	if err != nil {
		fmt.Println("error:", err)
	}

	if r.Method == "GET" {
		//w.Write([]byte("{\"id\":" + items[1].item + "}"))
		w.Write([]byte(b))
	}
}
