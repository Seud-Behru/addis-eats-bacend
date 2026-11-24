package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"https/github.com/Seud-Behru/addis-eats-bacend/internal/delivery"
	"https/github.com/Seud-Behru/addis-eats-bacend/internal/order"
	"https/github.com/Seud-Behru/addis-eats-bacend/menu"
	"net/http"
)

type createOrderRequest struct {
	Area    string   `json:"area"`
	DishIDs []string `json:"dish_ids"`
	Phone   string   `json:"phone"`
}

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1_048_576)

	var in createOrderRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "malformed JSON body")
		return
	}
 items := make([]menu.Dish, 0, len(in.DishIDs))
    for _, id := range in.DishIDs {
        d, err := app.dishes.ByID(id)
        if err != nil {              // unknown id, or store failure
            app.badRequest(w, "unknown dish "+id)
            return
        }
        items = append(items, d)     // price from the STORE
    }
 
    o := order.New(in.Area, in.Phone, items)
    total, err := o.ComputeTotal()   // domain rules, not handler rules
    if errors.Is(err, delivery.ErrUnservedArea) {
        app.badRequest(w, "we do not deliver to "+in.Area)
        return
    }
 fmt.Println(total)
}