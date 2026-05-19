package main

import (
	"fmt"
)

func pointer() {
	bigArray := [5]int{10, 20, 30, 40, 50}

	modifyWithoutPointer(bigArray)
	fmt.Println("After modifyWithoutPointer:", bigArray)

	modifyWithPointer(&bigArray)
	fmt.Println("After modifyWithPointer:", bigArray)
}

func modifyWithoutPointer(arr [5]int) {
	arr[0] = 999
	fmt.Println("Inside without pointer:", arr)
}

func modifyWithPointer(arr *[5]int) {
	arr[0] = 888
	fmt.Println("Inside with pointer:", arr)
}

