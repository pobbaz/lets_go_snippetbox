package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// create new serveMux(router)
	mux := http.NewServeMux()

	fileServer := http.FileServer(customFileSystem{http.Dir("./static")})
	// Strip the "/static" prefix so FileServer can find the right file
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// bind the url path to the handler func
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// start the web server
	log.Print("Listening on localhost:8000")
	err := http.ListenAndServe(":8000", mux)
	// if there is a error log it and stop the program
	log.Fatal(err)

}

// wrap the http.FileSystem with customFileSystem
type customFileSystem struct {
	fs http.FileSystem
}

func (cfs customFileSystem) Open(path string) (http.File, error) {
	f, err := cfs.fs.Open(path)
	if err != nil {
		return nil, err
	}
	// stat() -->  retrieves information about a file
	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		// check if directory contain index.html
		index := filepath.Join(path, "index.html")
		if _, err := cfs.fs.Open(index); err != nil {
			return nil, err
		}
		closeErr := f.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, err
	}
	return f, nil
}
