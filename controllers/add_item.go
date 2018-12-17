package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
    "bytes"

	"github.com/kkrisstoff/go-server/models"
)

type reqItem struct {
	message string
}

type resItem struct {
	id      int
	message string
}

// AddItem add new item
func AddItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var ri reqItem

		if err := decoder.Decode(&ri); err != nil {
			panic(err) // sent in response
		}
		defer r.Body.Close()

		newItem := models.ItemsStoreMapped.AddItem(ri.message)
		fmt.Println(newItem)
		idStr := strconv.Itoa(newItem.ID)

		var b bytes.Buffer

		b.WriteString(`{"id":`)
		b.WriteString(idStr)
		b.WriteString(`, "id":`)
		b.WriteString(newItem.Message)
		b.WriteString("}")

		w.Write(b.Bytes())
	}
	if r.Method == "GET" {
		fmt.Fprintf(w, "Use POST for addidng items") // send data to client side
	}
}
