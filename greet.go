package main

import (
	"fmt"
)

func greet(customerName string) {
	fmt.Printf("%52s %s\n", "Namaste ", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to Jaipur Bhojanalya! _/\\_ ")
	fmt.Println()
}

func sayTata(customerName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Printf("_/\\_ Thank you %v for visiting Jaipur Bhojanalya! _/\\_\n", customerName)
	fmt.Printf("%55s\n", "We hope to see you again!")
	fmt.Printf("%51s\n", "Have a nice day! ")
}
