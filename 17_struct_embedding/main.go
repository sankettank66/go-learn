package main

import (
	"fmt"
	"time"
)

type customer struct {
	name string
	id   int
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
	customer
}

func main() {
	customer1 := order{
		id:        "1",
		amount:    123.34,
		status:    "Not Paid",
		createdAt: time.Now(),
		customer:  customer{name: "Sanket", id: 1233},
	}

	fmt.Println(customer1)
}
