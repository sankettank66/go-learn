package main

import (
	"fmt"
	"maps"
)

func main() {
	// map declaration
	m := make(map[string]string)
	m2 := make(map[string]string)
	m["Key"] = "value"
	m["Key2"] = "value"
	m["Key"] = "null"
	fmt.Println(len(m))
	fmt.Println(m["form"])
	fmt.Println(m2["form"])
	for k := range m {
		delete(m, k)
	}
	m2["form"] = "value"
	fmt.Println(maps.Equal(m, m2))
	fmt.Println(m, m2)
}
