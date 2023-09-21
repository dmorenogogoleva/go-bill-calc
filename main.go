package main

import "fmt"

func main() {
	// struct

	mybill := newBill("mario's bill")

	fmt.Println(mybill.format())
	fmt.Println("hello world")
}
