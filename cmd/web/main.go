package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// define addr command line flag
	addr := flag.String("addr", ":4000", "http service address")
	// parse the command line flag
	flag.Parse()

	// create logger for info message
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// create logger for error message
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// instance of the application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Go HTTP server logs its own errors using the standard logger
	// For consistency, we should make it use our errorLog so we manually create an http.Server struct and tell it to use errorLog
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// start the web server
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	// if there is error log it and stop the program
	errorLog.Fatal(err)

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
