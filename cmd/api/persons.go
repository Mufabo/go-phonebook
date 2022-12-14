package main

import (
	"fmt"
	"net/http"

	"mufabo.github.io/go-phonebook/internal/data"
)

func (app *application) createPersonHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
		Number int32 `json:"number"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Dump the contents of the input struct in a HTTP response.
    fmt.Fprintf(w, "%+v\n", input)
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