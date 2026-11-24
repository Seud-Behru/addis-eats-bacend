package main

import (
	"errors"
	"https/github.com/Seud-Behru/addis-eats-bacend/menu"
	"net/http"
)



// health is a simple "are we alive?" check
func (app *application) health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

// listDishes returns every dish
func (app *application) listDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := app.dishes.All()
	if err != nil {
		app.logger.Println(err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	writeJSON(w, http.StatusOK, dishes)
}

// showDish returns one dish by ID
func (app *application) showDish(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	dish, err := app.dishes.ByID(id)

	switch {
	case errors.Is(err, menu.ErrNotFound):
		writeError(w, http.StatusNotFound, "dish not found")
	case err != nil:
		app.logger.Println(err)
		writeError(w, http.StatusInternalServerError, "internal error")
	default:
		writeJSON(w, http.StatusOK, dish)
	}
}

