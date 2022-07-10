package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// HomePage
func (app *application) HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "LoveServer")
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 Not Found")
		return

	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.ServerError(w, err)
	}

}

// Show Snippet
func (app *application) ShowSnippet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Server", "LoveServer")
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "This is Note %d", id)

}

// Create Snippet
func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "LoveServer")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Create Snippets Here....")
}
