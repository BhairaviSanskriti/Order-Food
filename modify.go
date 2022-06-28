package main

import "fmt"

func modifyOrder() {
	for {
		var isOrderOK string
		fmt.Println("Do you want to change your order?[y/n]")
		fmt.Scan(&isOrderOK)
		if isOrderOK != "y" {
			return
		}
		var serialNo uint
		var modifyType uint

		fmt.Println("Please enter the respective no. to proceed further: ")
		fmt.Println("Press '1' to update item quantity.")
		fmt.Println("Press '2' to delete an item from the order list.")
		fmt.Println("Press '3' to add item(s) in the order list.")
		fmt.Scan(&modifyType)

		switch modifyType {
		case 1:
			printMenu()
			fmt.Println("Please enter the S.No. of the item to be updated: ")
			fmt.Scan(&serialNo)
			updateQuantity(serialNo)

		case 2:
			printMenu()
			fmt.Println("Please enter the S.No. of the item to be deleted: ")
			fmt.Scan(&serialNo)
			delFromOrder(serialNo)

		case 3:
			insertIntoOrder()
		default:
			return
		}
		displayGeneratingBill()
		generateBill()
	}
}

func updateQuantity(serialNo uint) {

	var newQuantity uint

	for _, targetItem := range menu {

		if serialNo == targetItem.itemNo {

			itemName := targetItem.itemName        //name of the item whose quantity to be updated
			oldQuantity := customerOrder[itemName] //current quantity of that item

			fmt.Printf("Current quantity of %v is %v.\n", itemName, oldQuantity)
			fmt.Printf("Now, enter the updated quantity of %v to be ordered: \n", itemName)
			fmt.Scan(&newQuantity)

			//in case new quantity is '0' then delete that item from the order
			if newQuantity == 0 {
				delFromOrder(serialNo)
				return
			}
			fmt.Printf("")

			//update quantity of item
			customerOrder[targetItem.itemName] = newQuantity
			fmt.Printf("Updated the quantity of %v from %v to %v.\n", itemName, oldQuantity, newQuantity)

			//update bill
			subTotalBill -= float64(oldQuantity) * float64(targetItem.itemPrice) //delete the item price for old quantity
			subTotalBill += float64(newQuantity) * float64(targetItem.itemPrice) //add the item price for updated quantity

			break
		}
	}

}

func delFromOrder(serialNo uint) {

	for _, targetItem := range menu {
		if serialNo == targetItem.itemNo {
			itemName := targetItem.itemName //name of the item who is to be removed from the list
			//update bill; customerOrder[itemName] -> it results in the quantity of that item
			subTotalBill -= float64(customerOrder[itemName]) * float64(targetItem.itemPrice)
			//delete item
			delete(customerOrder, itemName)
			fmt.Printf("%v removed from the order list.\n", itemName)
			break
		}
	}
}
func insertIntoOrder() {
	//add an item in the order
	orderItems()
}
