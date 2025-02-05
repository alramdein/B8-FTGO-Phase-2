package main

// Liskov Substitution
// -- Good
type PaymentMethodWallet interface {
	Pay(amount float64)
}

type Dana struct{}

func (c *Dana) Pay(mount float64) {}

func NewDana(method string) PaymentMethodWallet {
	return &Dana{}
}

// -- Bad
type Gopay struct{}

func (p *Gopay) ProcessPayment(amount float64) {}

func NewGopay(method string) PaymentMethodWallet {
	return &Gopay{}
}

// -- Bad
type Ovo struct{}

func (p *Ovo) Pay(amount float32) {}

func NewOVO(method string) PaymentMethodWallet {
	return &Ovo{}
}

// -- Good (no problem different name of parameter)
type SuperBank struct{}

func (p *SuperBank) Pay(asdasd float64) {}

func NewSuperBank(method string) PaymentMethodWallet {
	return &SuperBank{}
}

func main3() {
	dana := &Dana{}
	gopay := &Gopay{}
	ovo := &Ovo{}

	dana.Pay(111)
	gopay.Pay(33)
	ovo.Pay(22)
}
