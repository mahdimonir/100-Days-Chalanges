package main

func main() {
	day := "Sat"

	// if day == "Sat" {
	// 	println("It's Saturday!")
	// } else {
	// 	println("Work day")
	// }

	switch day {
	case "Friday", "Sat":
		println("It's weekend!")
	default:
		println("Work day")
	}
}