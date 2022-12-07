package main

import (
	"net/http"
)

func (app *application) healthzHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Your request couldnot be processed", http.StatusInternalServerError)
		return
	}

}
