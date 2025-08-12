package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// request url must be exactly "/"
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.html",
	}

	// parse the template file
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Print(err.Error())
		app.serverError(w, err)
		return
	}

	// execute the base template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// show specific snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// get tht id parameter and convert it to integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Snippet %d", id)
}

// create a new snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// add "Allow: POST" header to response
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("this is the snippet create"))
}
