package main

import "net/http"

func (app *application) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to V.1.0!"))
}
func (app *application) status(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     "1.0",
		},
	}
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *application) findContactByID(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
        app.writeJSON(w, http.StatusBadRequest, envelope{"error": "Invalid ID parameter"}, nil)
        return
    }

    contact, err := app.sqlc.FindContactsByName(r.Context(), id) // Assuming you have a FindContactByID query
    if err != nil {
        app.writeJSON(w, http.StatusInternalServerError,envelope{"error": "Failed to find contact"}, nil)
        return
    }

    app.writeJSON(w, http.StatusOK, envelope{"contact": contact}, nil)
}