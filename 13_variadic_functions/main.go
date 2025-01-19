package main

import "fmt"

func sum(nums ...int) int {
	var sum = 0
	for _, i := range nums {
		sum += i
	}
	return sum
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...), "HELLO")
}
