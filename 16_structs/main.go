package main

import (
	"fmt"
	"time"
)

// int => 0
// float => 0
// string => ""
// boolean => false
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder
}
func main() {
	myOrder := newOrder("123", 12.34, "Paid")

	myOrder2 := order{
		id:        "124",
		amount:    234,
		status:    "delivered",
		createdAt: time.Now(),
	}

	myOrder.changeStatus("PAID")
	fmt.Println(myOrder.getAmount())
	fmt.Println(myOrder)
	fmt.Println(myOrder2)

	language := struct {
		name   string
		isGood bool
	}{"Golang", true}
	fmt.Println(language)
}
