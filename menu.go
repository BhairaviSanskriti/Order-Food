package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	//itemNo  itemName   itemPrice
	{1, "Adrakh Chai", 20},
	{2, "Filter Coffee", 25},
	{3, "Chhas", 35.50},
	{4, "Lassi", 30},
	{5, "Mango Lassi", 60},
	{6, "Kadi Chaawal", 50},
	{7, "Chhole Bhature", 45},
	{8, "Khasta Kachori", 30},
	{9, "Raj Kachori", 30},
	{10, "Veg. Sandwich", 20},
	{11, "Veg. Masala Maggi", 60.00},
	{12, "Samosa", 20},
	{13, "Cream Roll", 15},
}

func printMenu() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "S.No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf(" %-7d %.2f    %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}
