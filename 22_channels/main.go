package main

import (
	"fmt"
	"time"
)

func ProcessNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing int", num)

		time.Sleep(time.Millisecond * 500)
	}
}

func sum(result chan int, num1 int, num2 int) {
	newNum := num1 + num2
	result <- newNum
}

func main() {
	// numChan := make(chan int)
	// go ProcessNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// time.Sleep(time.Second * 5)
	// msg := <-numChan
	// fmt.Println(msg)
	result := make(chan int)

	go sum(result, 4, 5)
	res := <-result
	fmt.Println(res)
}
