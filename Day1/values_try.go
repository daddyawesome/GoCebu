package main

import "fmt"

func main() {
	// Declare a slice of strings
	colors := []string{"Red", "Green", "Blue", "Yellow"}

	// Loop over the slice with index and value
	for i, color := range colors {
		fmt.Printf("Color at index %d is %s\n", i, color)
	}
}
