package main

import (
	"fmt"
	"net/http"
)

func (app *application) createPersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this will create a new Person")
}

func (app *application) showPersonsHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)
	if  err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show details of single person %d\n", id)
}