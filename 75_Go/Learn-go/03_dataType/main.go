package main

import "fmt"

func main() {
	// Integer types
	var age int = 30
	println("Age:", age)

	// Unsigned integer types => Positive numbers only
	var count uint = 100
	println("Count:", count)

	// Floating-point types
	var price float64 = 19.99
	println("Price:", price)

	// Boolean type
	var isAvailable bool = true
	println("Is Available:", isAvailable)

	// String type
	var name string = "Go Programming"
	println("Name:", name)

	// Zero values => Default values for uninitialized variables
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	println(zeroFloat)

	// String literals
	fmt.Printf("ZeroInt: %d, zeroFloat: %f, zeroBool: %t, zeroString: '%s'\n", zeroInt, zeroFloat, zeroBool, zeroString)

	// Sprintf for formatted output without printing to console
	output := fmt.Sprintf("ZeroInt: %d, zeroFloat: %f, zeroBool: %t, zeroString: '%s'\n", zeroInt, zeroFloat, zeroBool, zeroString)
	println(output)
	
	// Complex types
	var complexNum complex128 = 3 + 4i
	println("Complex Number:", complexNum)
}