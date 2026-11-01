package main

import "fmt"

func main() {
	var totalBill float64
	fmt.Println("Welcome to the tip calculator")
	fmt.Println("What was the total bill?")
	fmt.Scanln(&totalBill)

	var tipValue int
	fmt.Println("What percentage tip would you give? 10, 12 o 15")
	fmt.Scanln(&tipValue)

	var totalPeople int
	fmt.Println("How many people to split the bill? ")
	fmt.Scanln(&totalPeople)

	totalBill *= 1 + float64(tipValue)/100

	eachPay := totalBill / float64(totalPeople)

	fmt.Printf("Each one pays: $%.2f\n", eachPay)
}
