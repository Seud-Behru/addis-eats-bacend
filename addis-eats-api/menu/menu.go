package menu

import "errors"

// Dish is what the kitchen knows about
type Dish struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PriceSantim int    `json:"price_santim"`
	IsFasting   bool   `json:"is_fasting"`
	Note        string `json:"note,omitempty"`
	internal    string // lowercase = never sent to the customer
}

// ErrNotFound is a sentinel error — the kitchen's way of saying "we don't have that"
var ErrNotFound = errors.New("dish not found")

// DishReader is a job poster: "I need someone who can do these two things"
type DishReader interface {
	All() ([]Dish, error)
	ByID(id string) (Dish, error)
}
