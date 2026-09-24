package main

import (
	"fmt"
)

func Multiply(a int, b int) int {
	return (a * b) / 2
}

const (
	MIN_NUMBER = 8
	Running    = false
	MX_NUM     = MIN_NUMBER + 2
	X          = 5.0
	Y          = 6.0
	Z          = (X * Y) * (MX_NUM * 6)
	MONTHLY    = (Y + 9) * MX_NUM * 4
	WEEKLY     = (150 * 4) / 4
)

type Objects struct {
	price          float64
	hoursPaid      float64
	balance        float64
	isMoneyBalance bool
}

func (objects Objects) Get() {

	objects.price = float64(Multiply(X, Y))

	objects.hoursPaid = float64(MX_NUM)

	var paidCalculation float64 = objects.price * objects.hoursPaid

	fmt.Println("Price:", objects.price, "Hours Worked:", objects.hoursPaid, "Hours & Got",
		fmt.Sprintf("%s%0.2f", "$", paidCalculation))

	fmt.Printf("Monthly: $%0.2f\n", float64(WEEKLY*4))
	fmt.Printf("Weekly: $%0.2f", float64(WEEKLY))

}

func ReadSystem() {
	fmt.Scanln()
}

func main() {

	objects := Objects{
		price:          0,
		hoursPaid:      0.0,
		isMoneyBalance: true,
		balance:        2.0,
	}

	if !Running {

		for objects.balance < 3.0 {

			if objects.isMoneyBalance {
				objects.Get()
			}

			objects.balance += 1.0
		}

		ReadSystem()
	}

}
