package main

import "fmt"

func main() {

	var number [6]int

	number[0] = 10
	number[1] = 20
	number[2] = 30
	number[3] = 40
	number[4] = 50
	number[5] = 60

	// fmt.Println("Array:", number)
	// fmt.Println("Length of array:", len(number))

	for i := 0; i < len(number); i++ {
		fmt.Printf("Element at index %d: %d\n", i, number[i])
	}
}