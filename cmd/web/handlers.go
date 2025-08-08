package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home handler function
func home(w http.ResponseWriter, r *http.Request) {
	// request url must be exactly "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// execute the base template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// show specific snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	// get tht id parameter and convert it to integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Snippet %d", id)
}

// create a new snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// add "Allow: POST" header to response
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("this is the snippet create"))
}
