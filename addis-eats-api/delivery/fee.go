package delivery

import "errors"

var ErrUnknownArea = errors.New("unknown delivery area")

// Fee returns the delivery fee for an area in santim
func Fee(area string) (int, error) {
	switch area {
	case "bole":
		return 2000, nil
	case "piassa":
		return 3000, nil
	case "mexico":
		return 2500, nil
	default:
		return 0, ErrUnknownArea
	}
}