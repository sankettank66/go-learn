package main

import "fmt"

func main() {
	var sum = 0
	fmnums := [5]int{1, 2, 3, 5, 4}
	for _, v := range fmnums {
		sum += v
	}
	fmt.Println(sum)
	for k, v := range "GOLANG" {
		fmt.Println(k, string(v))
	}
}
