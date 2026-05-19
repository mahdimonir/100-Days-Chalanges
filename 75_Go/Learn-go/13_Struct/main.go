package main

import "fmt"

type User struct {
	name  string
	email string
}

func main() {

	John := User{"John Doe", "john@example.com"} // positional
	Abraham := User{name: "Abraham Lincoln", email: "abraham@example.com"} // key value

	John.name = "John Smith"
	
	fmt.Printf("User1: %+v \n", John)
	fmt.Printf("User2: %+v \n", Abraham)
}