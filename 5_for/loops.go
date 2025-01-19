package main

import "fmt"

func main() {
	// classic for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// while
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	// infinite loop
	for {
		break
	}
	// range
	var sum = 0
	for i := range 10 {
		sum += i
	}
	fmt.Println(sum)
}
