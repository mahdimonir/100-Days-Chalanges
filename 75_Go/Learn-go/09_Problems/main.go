package main

import "fmt"

func main() {
	displayMenu := func() {
		fmt.Println("Welcome to the Grade Calculator!")
		fmt.Println("1) Calculate grade")
		fmt.Println("2) Check pass/fail status")
		fmt.Println("3) Exit")
		fmt.Print("Enter your option: ")
	}

	var choice int
	var score int

	running := true

	for running {
		displayMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter a score (0-100): ")
			fmt.Scan(&score)

			result := calculateGrade(score)
			if result == "Invalid score" {
				fmt.Println("Please enter a valid score between 0 and 100.")
			} else {
				fmt.Printf("Your grade is: %s\n", result)
			}
		case 2:
			fmt.Println("Enter a score (0-100): ")
			fmt.Scan(&score)
			status := checkPassFail(score)
			fmt.Printf("Status: %s\n", status)
		case 3:
			fmt.Println("Exiting program. Goodbye!")
			running = false
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func calculateGrade(score int) string {
	if score >= 90 && score <= 100 {
		return "A"
	} else if score >= 80 && score <= 89 {
		return "B"
	} else if score >= 70 && score <= 79 {
		return "C"
	} else if score >= 60 && score <= 69 {
		return "D"
	} else if score >= 0 && score < 60 {
		return "F"
	} else {
		return "Invalid score"
	}
}

func checkPassFail(score int) string {
	switch {
	case score >= 60 && score <= 100:
		return "Pass"
	case score >= 0 && score < 60:
		return "Fail"
	default:
		return "Invalid score"
	}
}