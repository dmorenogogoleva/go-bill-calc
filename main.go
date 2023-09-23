package main

import "fmt"

func main() {
	// struct

	mybill := newBill("mario's bill")

	mybill.updateTip(10)
	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 8.95)
	mybill.addItem("toffe pudding", 4.95)
	mybill.addItem("coffe", 3.25)
	fmt.Println(mybill.format())
}
