package main

import "fmt"

func main() {
	// initilization
	// var nums [4]int
	// declaration
	var nums = [4]int{1, 2, 3, 4}

	//length
	fmt.Println(len(nums))

	// for i := 0; i < len(nums); i++ {
	// 	nums[i] = i
	// }
	fmt.Println(nums)

	var strArray [3]string
	strArray[0] = "first"

	fmt.Println(len(strArray[0]))

	nums2 := [5]int{1, 2, 4, 5}
	fmt.Println(nums2)

	// Dynamic Arrays

}
