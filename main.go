package main

import (
	"fmt"
	"strings"
	"time"
)

func Multiply(a int, b int) int {
	return (a * b) / 2
}

type DateNow struct{}

func (date DateNow) GetDate() string {
	if time.Now().Hour() < 10 {
		return strings.ToUpper("Good Morning")
	} else if time.Now().Hour() > 11 && time.Now().Hour() < 17 {

		return strings.ToUpper("Good Afternoon")
	} else if time.Now().Hour() > 17 {

		return strings.ToUpper("Good Evening")
	}

	return string("")

}

func (date DateNow) GetDateNow() string {
	now := time.Now()
	return strings.ToUpper(fmt.Sprintf("%s/%d/%d", now.Month().String(), now.Day(), now.Year()))
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

type TITLE struct {
	AppTitle string
}

func ServeTitles() {
	titles_app := TITLE{
		AppTitle: strings.ToUpper("Weekly Paycheck Balance"),
	}

	fmt.Println(titles_app.AppTitle)
}

func (title TITLE) GetTitle() {

	ServeTitles()
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

type PrintMessageln struct {
	PrinterMessage string
}

func (println PrintMessageln) GetPrintedLine() {
	fmt.Println(println.PrinterMessage)
}

type SLEEP struct {
	sleeping bool
	LENGTH   float64
}

func (sleep SLEEP) Sleep() {
	if sleep.sleeping {
		time.Sleep(time.Duration(sleep.LENGTH) * time.Second)
	}
}

func main() {

	objects := Objects{
		price:          0,
		hoursPaid:      0.0,
		isMoneyBalance: true,
		balance:        2.0,
	}

	SLEEP := SLEEP{
		sleeping: true,
		LENGTH:   float64(20.0 / 2),
	}

	if SLEEP.sleeping {
		SLEEP.Sleep()
	}

	date := DateNow{}

	date.GetDateNow()
	printDate := PrintMessageln{
		date.GetDateNow(),
	}

	printDate.GetPrintedLine()

	println := PrintMessageln{
		date.GetDate(),
	}

	println.GetPrintedLine()

	title := TITLE{}
	title.GetTitle()

	fmt.Printf("%s", "\n")

	SLEEP.LENGTH = float64(144/12 - 3)

	if !Running {

		if SLEEP.sleeping {
			SLEEP.Sleep()
		}

		for objects.balance < 3.0 {

			if objects.isMoneyBalance {
				objects.Get()
			}

			objects.balance += 1.0
		}

		ReadSystem()
	}

}
