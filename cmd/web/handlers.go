package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/kellemNegasi/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// ts, err := template.ParseFiles("./ui/html/home.page.tmpl")

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }
	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.serverError(w, err) // here using the server error helper
	// 	return
	// }

	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	app.serverError(w, err) // using the server error
	// 	return
	// }
}
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, s)
	if err != nil {
		app.serverError(w, err)
	}
	// fmt.Fprintf(w, "%v", s)
}
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allowed", "POST")
		app.clientError(w, http.StatusMethodNotAllowed) // using the client error helper
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly,“slowly!\n\n– Kobayashi Issa"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}

	// w.Write([]byte("Create a specific snippet\n"))
	// redirect the user to the added item
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
