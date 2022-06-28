package main

import (
	"fmt"
)

//make subTotalBill as global variable to make it easily accessible in case customer modifies her order.
// subTotalBill means total bill but excluding taxes.
var subTotalBill float64

//make a map of customerOrder in which we will store the items ordered as "key" and no. of plates as "value".
var customerOrder = make(map[string]uint, 0)

func main() {
	var customerName string
	fmt.Println("What is your first name?")
	fmt.Scan(&customerName)

	greet(customerName)
	orderItems()
	displayGeneratingBill() //just displays that "generating bill" in a fancy manner.
	generateBill()
	modifyOrder()
	printFinalBill()
	sayTata(customerName)
}
