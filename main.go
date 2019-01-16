package main

import (
	"github.com/kkrisstoff/go-server/controllers"
	"log"
	"net/http"
)

func main() {
	StartServer()
}


func StartServer() {
	http.HandleFunc("/api/addItem", controllers.AddItem)
	http.HandleFunc("/api/getItem", controllers.GetItemByID)
	http.HandleFunc("/api/items", controllers.GetItems)
	http.HandleFunc("/api/deleteItem", controllers.DeleteItemById)
	log.Fatal(http.ListenAndServe(":9091", nil))
}