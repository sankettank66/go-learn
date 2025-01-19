package main

import "fmt"

func printSlices[T string | int](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
}
func printStringSlices(items []string) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
}

type Stack[T string | int] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3}
	names := []string{"sanket", "Guddu", "prem"}
	printSlices(nums)
	printSlices(names)
	printStringSlices(names)
}
