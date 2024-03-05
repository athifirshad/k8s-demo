package main

import "net/http"

func (app *application) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to V.1.0!"))
}
