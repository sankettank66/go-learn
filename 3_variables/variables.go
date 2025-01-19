package main

import (
	"fmt"
	"math"
)

func main() {
	var name string = "Hello World"
	var age int = 62
	var price float32 = 17 / 3.0
	var isAdult = true
	max := math.Max(10, 20)

	fmt.Println(name, age, price, isAdult, max)
}
