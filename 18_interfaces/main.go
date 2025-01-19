package main

import "fmt"

type PaymentGateway interface {
	pay(amount float32)
	refund(amount float32, accountNumber string)
}

type payment struct {
	gateway PaymentGateway
}

func (p payment) makePayment(amount float32) {
	// razorpayPayment := razorpay{}
	// razorpayPayment.pay(amount)
	// stripePayment := stripe{}
	// stripePayment.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

type paypal struct{}

func (r paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}
func (r paypal) refund(amount float32, accountNumber string) {
	fmt.Println(r)
}
func main() {
	// razorpayGW := razorpay{}
	// stripeGW := stripe{}
	paypalGW := paypal{}
	// newPayment := payment{
	// 	gateway: razorpayGW,
	// }
	// newPayment2 := payment{
	// 	gateway: stripeGW,
	// }
	newPayment3 := payment{
		gateway: paypalGW,
	}
	// newPayment.makePayment(1000)
	// newPayment2.makePayment(1000)
	newPayment3.makePayment(1000)
	newPayment3.gateway.refund(123, "Hello")
}
