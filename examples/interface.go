package examples

import (
	"fmt"
	"math"
)

type Car interface {
	Price() int

	Open()
	Close()

	GoForward(km int) (int, error)
}

func BuyCar(c Car, money int) error {
	if money < 0 {
		return fmt.Errorf("money cannot be negative")
	}

	if c.Price() < money {
		return fmt.Errorf("cannot buy a car")
	}

	return nil
}

type Volvo struct {
	isOpen bool
	fuel   int
	price  int
}

func (v Volvo) Price() int {
	return v.price
}

func (v Volvo) Open() {
	v.isOpen = true
}

func (v Volvo) Close() {
	v.isOpen = false
}

func (v Volvo) GoForward(km int) (int, error) {
	if !v.isOpen {
		return 0, fmt.Errorf("car is closed")
	}

	needFuel := int(math.Ceil(float64(km / 100)))
	if needFuel > v.fuel {
		v.fuel -= needFuel
		return needFuel, nil
	}

	return 0, fmt.Errorf("fuel not enough")
}
