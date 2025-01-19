package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("demo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo, err := f.Stat()

	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	buff := make([]byte, fileInfo.Size())
	d, err := f.Read(buff)
	if err != nil {
		panic(err)
	}
	// for i := 0; i < len(buff); i++ {
	// 	fmt.Println(d, "d", string(buff[i]))
	// }
	fmt.Println(d, "d", string(buff))
	// f2, err2 := os.ReadFile("demo.txt")
	// if err2 != nil {
	// 	panic(err2)
	// }
	// fmt.Println(string(f2))

	// read folders
	f2, err2 := os.Open("../")
	if err2 != nil {
		panic(err2)
	}
	defer f2.Close()
	fileInfo2, err := f2.ReadDir(-4)

	for _, file := range fileInfo2 {
		fmt.Println(file.Name(), file.IsDir())
	}

	// create file

	f3, err := os.Create("demo2.txt")
	if err != nil {
		panic(err)
	}
	f3.WriteString("Hii Golang")
}
