package model

type IExternalShippingService interface {
	GetShippingCost(origin string, destination string, weight int) float64
}

type IBankAPI interface {
	SetBankProvider(bankAPI IExternalShippingService)
	GetShippingCost(origin string, destination string, weight int) float64
}
