package view

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

const viewsFolder = "static/"
const viewPrefix = "/"
const itmsView = "items"

var templates = template.Must(template.ParseFiles(viewsFolder+"index.html", viewsFolder+itmsView+".html"))

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := viewsFolder + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	view := r.URL.Path[len(viewPrefix):]
	if view == itmsView {
		renderView(itmsView, w)
		return
	}
	renderView("index", w)
}

func renderView(view string, w http.ResponseWriter) {
	fmt.Printf("Render View %s\n", view)
	page := "test"
	p, _ := loadPage(page)
	err := templates.ExecuteTemplate(w, view+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
