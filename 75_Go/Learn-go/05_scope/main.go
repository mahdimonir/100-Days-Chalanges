package main

import "fmt"

func main() {
	outerItem := "sugar"
	makeCoffee := func ()  {
		innerItem := "Cappuccino"
		fmt.Printf("Making %s with %s", innerItem, outerItem)
	}
	makeCoffee()

	// variable shadowing
	makeCoffee2 := func ()  {
		innerItem := "Cappuccino"
		outerItem := "sugar"
		fmt.Printf("Making %s with %s", innerItem, outerItem)
	}
	makeCoffee2()
}