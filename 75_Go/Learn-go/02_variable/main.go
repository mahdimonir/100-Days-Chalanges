package main

func main() {
	// var name string
	// name = "Go"

	// var name string = "Go"

	// var name = "Go"

	// name := "Go"

	// println("Hello", name)

	var (
		name string = "John"
		age  int    = 22
	)
	println("My name is", name, ", I'm", age, "years old.")

	// Multiple variable declaration
	// var firstName, lastName string
	// firstName = "John"
	// lastName = "Doe"
	var firstName, lastName string = "John", "Doe"
	println("My name is", firstName, lastName)

	const pi = 3.1416
	println("The value of pi is", pi)
}