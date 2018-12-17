package main

import (
	"log"
	"net/http"

	"github.com/kkrisstoff/go-server/controllers"
)

func main() {
	//conf := conf.GetConfig()
	//fmt.Println("Server config: ", conf)

	//mux := &Mux{}
	//http.ListenAndServe(":9090", mux)
	http.HandleFunc("/api/addItem", controllers.AddItem)
	http.HandleFunc("/api/getItem", controllers.GetItemById)
	http.HandleFunc("/api/items", controllers.GetItems)
	http.HandleFunc("/api/deleteItem", controllers.DeleteItemById)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

// type Mux struct {
// }
// func (p *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		sayHello(w, r)
// 		return
// 	}
// 	http.NotFound(w, r)
// 	return
// }
