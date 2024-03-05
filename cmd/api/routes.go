package main

import (
	"net/http"
)

func (app *application) Routes() {
	app.router.Get("/", app.rootHandler)
	app.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Page Not Found"))
	})
}
