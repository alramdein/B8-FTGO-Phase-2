package main

// Dependecy Inversion
// -- Good
type PaymentMethod2 interface {
	Pay(amount float64)
}

type CreditCard1 struct{}

func (cc CreditCard1) Pay(amount float64) {}

type Paypal1 struct{}

func (p Paypal1) Pay(amount float64) {}

type PaymentService2 struct {
	method PaymentMethod2
}

func (ps *PaymentService2) SetPaymentMethod(method PaymentMethod2) {
	ps.method = method
}

func (ps PaymentService2) ProcessPayment(amount float64) {
	ps.method.Pay(amount)
}

func main() {
	cc := CreditCard1{}
	pp := Paypal1{}

	paymentSvc := PaymentService2{}

	// paymentSvc.ProcessPayment(100) // error

	paymentSvc.SetPaymentMethod(cc) // inject dulu
	paymentSvc.ProcessPayment(100)
	paymentSvc.ProcessPayment(42123)
	paymentSvc.ProcessPayment(232)
	paymentSvc.ProcessPayment(23423)
	paymentSvc.ProcessPayment(5313)

	paymentSvc.SetPaymentMethod(pp) // inject dulu
	paymentSvc.ProcessPayment(100)
	paymentSvc.ProcessPayment(42123)
	paymentSvc.ProcessPayment(232)
	paymentSvc.ProcessPayment(23423)
	paymentSvc.ProcessPayment(5313)
}
