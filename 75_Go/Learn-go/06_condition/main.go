package main

import "fmt"

func ageStatus(age int) {}

func main() {
	age := 20

	if age >= 18 {
		fmt.Println("Adult!")
	}

	score := 80
	fmt.Printf("Outside score is: %d\n", score)
	if score := 70; score >= 80 {
		fmt.Println("You got Gold Medal and your score is", score)
	} else if score >= 70 {
		fmt.Println("You got Silver Medal and your score is", score)
	} else {
		fmt.Println("You got participation certificate and your score is", score)
	}
}