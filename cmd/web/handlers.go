package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// ts, err := template.ParseFiles("./ui/html/home.page.tmpl")

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err) // here using the server error helper
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // using the server error
		return
	}
}
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	// w.Write([]byte("Display a specific snippet\n"))
	fmt.Fprintf(w, "Display a specific snippet with ID %d \n\n", id)
}
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allowed", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a specific snippet\n"))
}
