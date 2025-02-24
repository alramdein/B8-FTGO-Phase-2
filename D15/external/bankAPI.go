package external

type BankAPI interface {
	Transfer(from string, to string, amount int) error
}

type bankAPI struct {
	provider model.Bank
	bca
	bri
	mandiri
}

func NewBankAPI(bankAPI IExternalShippingService) IBankAPI {
	type provider model.Bank
	switch bankName {
	case "bca":
		provider = newBankBCA()
	case "bni":
		provider = newBankBRI()
	}

	return bankAPI{
		bca: newBankBCA(),
		// ...
	}
}
