package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/api/v1/healthz", app.healthzHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/persons", app.createPersonHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/persons/:id", app.showPersonsHandler)

	return router
}