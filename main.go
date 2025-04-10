package main

import (
	"fmt"
)

func main() {
	var numbers [5]int
	var sum int

	// Slice
	fmt.Println("Type five numbers:")

	for i := 0; i < 5; i++ {
		fmt.Printf("Number %d: ", i+1)
		fmt.Scanln(&numbers[i])
		sum += numbers[i]
	}

	// Slice
	fmt.Printf("The total sum is: %d\n", sum)
}