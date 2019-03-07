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
	startServer(conf)
}

func startServer(conf config.Config) {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/api/addItem", controllers.AddItem)
	http.HandleFunc("/api/item", controllers.Item)
	http.HandleFunc("/api/items", controllers.GetItems)

	// TODO: refactor static folder
	http.HandleFunc("/add", view.AddItem)
	http.HandleFunc("/", view.ViewHandler)

	fmt.Printf("...starting on %s\n", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Host, conf.Port), nil))
}
