package order

import (
	"errors"
	"https/github.com/Seud-Behru/addis-eats-bacend/menu"
)
type Order struct {
	ID      string
	Area    string
	DishIDs []string
	Phone   string
	Status  string
}

var (
	ErrNotFound      = errors.New("order not found")
	ErrInvalidArea   = errors.New("invalid delivery area")
	ErrNoDishes      = errors.New("order must have at least one dish")
)


func Total(dishes []menu.Dish, area string) (int, error) {
	if len(dishes) == 0 {
		return 0, ErrNoDishes
	}

	total := 0
	for _, d := range dishes {
		total += d.PriceSantim
	}

	fee, err := deliveryFee(area)
	if err != nil {
		return 0, err
	}

	return total + fee, nil
}

func deliveryFee(area string) (int, error) {
	switch area {
	case "bole":
		return 2000, nil
	case "piassa":
		return 3000, nil
	case "mexico":
		return 2500, nil
	default:
		return 0, ErrInvalidArea
	}
}


 



