package main

import "fmt"

// reference by value
func changeNum(num int) {
	num = 5
	fmt.Println(num)
}

// reference by address
func changeNumber(num *int) {
	*num = 5
	fmt.Println(num)
}

func main() {
	num := 1
	fmt.Println(&num)
	changeNum(num)
	fmt.Println(num)
	changeNumber(&num)
	fmt.Println(num)
}
