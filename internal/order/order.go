package order

import "fmt"

func LineTotal(priceSantim, qty int) int {
	return priceSantim * qty
}

func ServiceFeeSantim(subtotal int) int {
	return int(float64(subtotal) * 0.05)
}

func DeliveryFeeSantim(area string) (int, error) {
    switch area {
    case "Bole":
        return 3000, nil
    case "Kazanchis":
        return 3500, nil
	}
	return 0, fmt.Errorf("unserved area %q", area)	
}


 



