package view

import (
	"fmt"
	"github.com/kkrisstoff/go-server/csv"
	"github.com/kkrisstoff/go-server/models"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type Page struct {
	Title string
	Body  []byte
}

const templatesFolder = "templates/"
const viewPrefix = "/"
const itmsView = "items"

var layoutPath = path.Join("templates", "layout.html")
var templates = template.Must(template.ParseFiles(
	layoutPath,
	templatesFolder+itmsView+".html",
))

func log(s string) {
	fmt.Printf(">>> %s\n", s)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "static" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	view := r.URL.Path[len(viewPrefix):]
	fp := path.Join("templates", r.URL.Path)

	log(fp)
	log(view)
	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}
	csv.ReadData()
	renderItems(w, fp)
	// renderView("index", w)
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

func AddItem(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	message := r.FormValue("message")
	models.ItemsStoreMapped.AddItem(name, message)
	csv.WriteData(name, message)
	// renderItems(w)
}

type ItemsData struct {
	Title string
	Items []models.Item
}

func renderItems(w http.ResponseWriter, path string) {
	items := models.ItemsStoreMapped.GetItems()
	data := ItemsData{
		Title: "Page Items",
		Items: items,
	}

	templates, err := template.ParseFiles(layoutPath, path)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "layout", &data)
	// err := templates.ExecuteTemplate(w, layoutPath, &data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
