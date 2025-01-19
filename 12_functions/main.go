package main

import "fmt"

// Define the add function
func add(x int, y int) int {
	return x + y
}

// Define the sub function
func sub(x, y int) int {
	return x - y
}

// Define the min function (we can assume it subtracts the numbers)
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Define the getLanguages function
func getLanguages() (string, int, bool) {
	return "Golang", 10, true
}

// Define the calc function to use the passed function
func calc(x int, y int, fn func(a int, b int) int) int {
	return fn(x, y) // Use the passed function here
}

func main() {
	// Calling add function
	fmt.Println(add(10, 20)) // Output: 30

	// Calling min function
	fmt.Println(min(10, 20)) // Output: 10

	// Calling getLanguages function
	fmt.Println(getLanguages()) // Output: Golang 10 true

	// Calling calc with min function
	result := calc(5, 3, sub)
	fmt.Println(result)
}
