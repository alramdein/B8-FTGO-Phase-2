package main

import (
	"hacktiv/external"
	"hacktiv/repository"
	"hacktiv/usecase"
)

func main() {
	bankAPIs := external.NewBankAPIs()
	rajaOngkir := external.NewRajaOngkirAPI()
	ratuOngkir := newn
	shippingRepo := repository.NewShippingRepository()

	shippingUsecase := usecase.NewShippingUsecase(shippingRepo, bankAPIs).GetShippingCost()
}
