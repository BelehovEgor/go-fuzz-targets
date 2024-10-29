package examples

import (
	"encoding/json"
	"errors"
	"fmt"
)

func CalculateTotal(price float64, quantity int, taxRate float64) (float64, error) {
	if price < 0 {
		return 0, errors.New("price should not be negative")
	}
	if quantity < 0 {
		return 0, errors.New("quantity should not be negative")
	}
	if taxRate < 0 {
		return 0, errors.New("taxRate should not be negative")
	}

	subtotal := price * float64(quantity)

	tax := subtotal * taxRate

	total := subtotal + tax

	return total, nil
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
	TaxRate  float64
}

func CalculateTotalForProduct(product Product) (float64, error) {
	if product.Price < 0 {
		return 0, errors.New("price should not be negative")
	}
	if product.Quantity < 0 {
		return 0, errors.New("quantity should not be negative")
	}
	if product.TaxRate < 0 {
		return 0, errors.New("taxRate should not be negative")
	}

	subtotal := product.Price * float64(product.Quantity)

	tax := subtotal * product.TaxRate

	total := subtotal + tax

	return total, nil
}

func GenerateJSON(args ...interface{}) (string, error) {
	data := make(map[string]interface{})

	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)
			if !ok {
				return "", fmt.Errorf("key is not a string")
			}
			data[key] = args[i+1]
		} else {
			return "", fmt.Errorf("odd number of arguments")
		}
	}

	// Marshal the map into a JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
