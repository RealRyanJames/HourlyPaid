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
)

func main() {

	if !Running {

		var price float64 = 0
		var hoursPaid float64 = 0.0
		var isMoneyBalance bool = true
		var balance float64 = 2.0

		for balance < 3.0 {

			if isMoneyBalance {
				price = float64(Multiply(X, Y))

				hoursPaid = float64(MX_NUM)
				var paidCalculation float64 = price * hoursPaid

				fmt.Println("Price:", price, "Hours Worked:", hoursPaid, "Hours & Got",
					fmt.Sprintf("%s%0.2f", "$", paidCalculation))

			}

			balance += 1.0
		}
	}

}
