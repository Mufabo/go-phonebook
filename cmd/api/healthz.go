package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status availabe")
	fmt.Fprintf(w, "Environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
