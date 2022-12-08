package main

import (
	"fmt"
	"net/http"

	"mufabo.github.io/go-phonebook/internal/data"
)

func (app *application) createPersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this will create a new Person")
}

func (app *application) showPersonsHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)
	if  err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	person := data.Person{
		ID: id,
		Name: "Harry",
		Number: 1234,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"person":person}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}