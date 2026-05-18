package main

import "fmt"

// Function definition
func makeCoffee(kind string, isSugar bool) (string, int) {
	price := 25
	coffee := fmt.Sprintf("Making a cup of %s , isSugar: %t,", kind, isSugar)
	return coffee, price
}

// Named return values
func makeTea(kind string, isLemon bool) (tea string, price int){
	price = 15
	tea = fmt.Sprintf("Making a cup of %s tea, isLemon: %t,", kind, isLemon)
	return
}

func main() {
	myCoffee, myBill := makeCoffee("latte", true)
	ranaCoffee, ranaBill := makeCoffee("cappuccino", false)
	tea, teaPrice := makeTea("green", true)
	fmt.Print(myCoffee, " and my bill is: ", myBill, "\n")
	fmt.Print(ranaCoffee, " and Rana's bill is: ", ranaBill, "\n")
	fmt.Print(tea, " and the tea price is: ", teaPrice, "\n")
	
	// Anonymous Function
	makeingCoffee := func() {
		fmt.Printf("Making coffee")
	}
	makeingCoffee()

	// IIFE => Immeditly Invoked Function Expression
	func (coffeeType string){
		fmt.Printf("Making hot %s...", coffeeType)
	}("Latte")
}
