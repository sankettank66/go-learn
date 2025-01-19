package main

import (
	"fmt"
	"time"
)

func main() {
	// Single condition switch
	age := 154
	switch age {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("DEFAULT")
	}
	// Multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is ", time.Saturday)
	case time.Sunday:
		fmt.Println("Today is ", time.Sunday)
	default:
		fmt.Println("Today is workday")
	}
	// type condition switch

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's an String")
		case bool:
			fmt.Println("It's an Boolean")
		default:
			fmt.Println(i)
		}

	}

	whoAmI(50)
	whoAmI(true)
	whoAmI(2.5)
}
