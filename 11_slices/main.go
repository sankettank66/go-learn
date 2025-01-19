package main

import (
	"fmt"
	"sort"
)

func main() {

	//Slices Dynamic Arrays
	// uninitiated slices are nil
	var arr []int
	fmt.Println(arr)
	fmt.Println(arr == nil)
	fmt.Println(len(arr))

	var arr2 = make([]int, 2)
	fmt.Println(cap(arr2))
	fmt.Println(len(arr2))
	arr2 = append(arr2, 1)
	arr2 = append(arr2, 2)
	arr2 = append(arr2, 3)
	fmt.Println(arr2)
	sort.Slice(arr2, func(i, j int) bool {
		return i < j
	})
	fmt.Println(arr2)
	fmt.Println(arr2[len(arr2)-1])
	fmt.Printf("arr2: %v\n", arr2)

	//second way to declare slices
	arr3 := []int{}
	fmt.Println(arr3)
	for x, y := range arr2 {
		fmt.Println(x, y)
	}
}
