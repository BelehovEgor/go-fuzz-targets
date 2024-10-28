package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// calculateTotal function takes three parameters: price, quantity, and taxRate
func calculateTotal(price float64, quantity int, taxRate float64) (float64, error) {
	// Validate inputs
	if price < 0 {
		return 0, errors.New("price should not be negative")
	}
	if quantity < 0 {
		return 0, errors.New("quantity should not be negative")
	}
	if taxRate < 0 {
		return 0, errors.New("taxRate should not be negative")
	}

	// Calculate the subtotal
	subtotal := price * float64(quantity)

	// Calculate the tax
	tax := subtotal * taxRate

	// Calculate the total cost
	total := subtotal + tax

	return total, nil
}

// Product struct represents a product with name, price, quantity, and tax rate
type Product struct {
	Name     string
	Price    float64
	Quantity int
	TaxRate  float64
}

// calculateTotal function takes a Product struct as an argument and calculates the total cost
func calculateTotalForProduct(product Product) (float64, error) {
	// Validate inputs
	if product.Price < 0 {
		return 0, errors.New("price should not be negative")
	}
	if product.Quantity < 0 {
		return 0, errors.New("quantity should not be negative")
	}
	if product.TaxRate < 0 {
		return 0, errors.New("taxRate should not be negative")
	}

	// Calculate the subtotal
	subtotal := product.Price * float64(product.Quantity)

	// Calculate the tax
	tax := subtotal * product.TaxRate

	// Calculate the total cost
	total := subtotal + tax

	return total, nil
}

// generateJSON function takes a variable number of arguments and generates a JSON string
func generateJSON(args ...interface{}) (string, error) {
	// Create a map to hold the arguments
	data := make(map[string]interface{})

	// Assuming the arguments are provided in key-value pairs
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)
			if !ok {
				return "", fmt.Errorf("key at index %d is not a string", i)
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

func main() {
	// Calling the generateJSON function with some sample arguments
	jsonString, err := generateJSON("name", "Alice", "age", 25, "isStudent", true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("JSON String:", jsonString)
	}
}
