package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill: - ", b.name)
	return b
}

func propmptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip):", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number", err)
			propmptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		propmptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number", err)
			propmptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		propmptOptions(b)
	case "s":
		fmt.Println("you chose to save the file - ", b.name)
		b.save()
	default:
		fmt.Println("that what not a valid option...")
		propmptOptions(b)
	}

}

func main() {
	mybill := createBill()
	propmptOptions(mybill)
}
