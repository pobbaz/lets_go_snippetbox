package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError() logs the error and stack trace, then sends a 500 Internal Server Error to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// Show correct file and line number
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

// clientError() sends a specific status code (e.g., 400, 404) and message to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound() is just a shortcut for sending a 404 Not Found.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
