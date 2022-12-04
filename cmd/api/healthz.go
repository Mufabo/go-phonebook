package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthzHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Your request couldnot be processed", http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
