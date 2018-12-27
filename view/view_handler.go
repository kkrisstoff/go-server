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

var templates = template.Must(template.ParseFiles(viewsFolder+"index.html", viewsFolder+"view.html"))

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
	// title := r.URL.Path[len(viewsFolder):]
	title := "test"
	fmt.Printf("Title: %s\n", title)
	p, _ := loadPage(title)
	tmpl := "index"
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
