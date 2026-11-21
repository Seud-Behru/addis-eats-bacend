package main

import "net/http"

// routes returns the mux with every route registered
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", app.health)
	mux.HandleFunc("GET /menu", app.listDishes)
	mux.HandleFunc("GET /menu/{id}", app.showDish)
	mux.HandleFunc("POST /order", app.createOrder)
	// POST /menu → 405 automatically (no code needed)

	return mux
}