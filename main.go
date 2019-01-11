package main

import (
	"fmt"
	"github.com/kkrisstoff/go-server/config"
	"github.com/kkrisstoff/go-server/controllers"
	"github.com/kkrisstoff/go-server/view"
	"log"
	"net/http"
)

func main() {
	conf, _ := config.GetConfig()
	// TODO: handle error
	StartServer(conf)
}

func StartServer(conf config.Config) {
	http.HandleFunc("/api/addItem", controllers.AddItem)
	http.HandleFunc("/api/item", controllers.Item)
	http.HandleFunc("/api/items", controllers.GetItems)

	http.HandleFunc("/", view.ViewHandler)
	fmt.Printf("...starting on %s\n", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Host, conf.Port), nil))
}
