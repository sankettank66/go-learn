package main

import "fmt"

type OrderStatusType string

const (
	FULFILLED OrderStatusType = "fullfilled"
	PENDING                   = "pending"
	REJECTED                  = "rejected"
)

func changeOrderStatus(status OrderStatusType) {
	fmt.Println("Changing Order Status to", status)
}

func main() {
	changeOrderStatus(FULFILLED)
}
