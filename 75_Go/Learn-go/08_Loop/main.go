package main

func makeCoffee(coffeeNo int) {
	println("Making coffee...", coffeeNo)
}

func main() {
	// for loop
	// for i := 1; i <= 5; i++ {
	// 	makeCoffee(i)
	// }

	// while loop
	// j := 1
	// for j <= 5 {
	// 	makeCoffee(j)
	// 	j++
	// }

	// infinite loop
	// for {
	// 	println("This is an infinite loop")
	// }

	// break and continue
	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			continue // skip even numbers
		}
		// if k == 7 {
		// 	break // stop when we reach 7
		// }
		makeCoffee(k)
	}
}