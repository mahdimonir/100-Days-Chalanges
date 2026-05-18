package main

import "fmt"

func main() {
	// Partial array initialization
	var numbers = [6]int{10, 20, 30, 40, 50, 60}
	fmt.Println("Original Array:", numbers)

	// [startIndex:endIndex] - length, capacity and pointer
	sliceNum := numbers[1:4]
	fmt.Println("Slice:", sliceNum)

	sliceNum[0] = 25
	sliceNum = append(sliceNum, 45)
	sliceNum = append(sliceNum, 47)
	sliceNum = append(sliceNum, 55) // It will not be in original array because of capacity
	fmt.Println("Modified Slice:", sliceNum)
	fmt.Println("Original Array after modification:", numbers)

	fmt.Printf("Length of slice: %d\n", len(sliceNum))
	fmt.Printf("Capacity of slice: %d\n", cap(sliceNum))
}