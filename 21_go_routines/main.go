package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Doing Task", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i+1, &wg)
	}
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }
	// time.Sleep(time.Second * 5)
	wg.Wait()
}
