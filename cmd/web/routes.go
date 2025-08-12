package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// create new serveMux(router)
	mux := http.NewServeMux()
	// serve static file
	fileServer := http.FileServer(customFileSystem{http.Dir("./static")})
	// Strip the "/static" prefix so FileServer can find the right file
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	// bind the url path to the handler func
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
