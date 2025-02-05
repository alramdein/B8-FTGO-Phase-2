package main

// Open/Closed
// -- Bad
type PaymentService struct{}

func (o *PaymentService) ProcessPayment(method string, amount float64) {
	if method == "credit_card" {
		// Process credit card payment
	} else if method == "paypal" {
		// Process PayPal payment
	} else if method == "crypto" {
		// Process crypto payment
	}
	// Adding more method require adding more if statement
}

// -- Good
type PaymentMethod interface {
	Pay(amount float64)
}

type CreditCard struct{}

func (c *CreditCard) Pay(mount float64) {}

type Paypal struct{}

func (p *Paypal) Pay(amount float64) {}

type Crypto struct{}

func (c *Crypto) Pay(amount float64) {}

// ....

func main2() {
	cc := &CreditCard{}
	pp := &Paypal{}
	cr := &Crypto{}

	cc.Pay(100)
	pp.Pay(100)
	cr.Pay(100)
}
