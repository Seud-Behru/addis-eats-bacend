package main

import "net/http"

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.logger.Println(err) // full detail
	writeError(w, http.StatusInternalServerError, "internal error")
}

func (app *application) notFound(w http.ResponseWriter, msg string) {
	writeError(w, http.StatusNotFound, msg)
}

func (app *application) badRequest(w http.ResponseWriter, msg string) {
	writeError(w, http.StatusBadRequest, msg)
}