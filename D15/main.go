package main

import (
	"hacktiv/external"
	"hacktiv/repository"
	"hacktiv/usecase"
)

func main() {

	rajaOngkir := external.NewRajaOngkirAPI()
	shippingRepo := repository.NewShippingRepository()

	shippingUsecase := usecase.NewShippingUsecase(shippingRepo, rajaOngkir).GetShippingCost()

	println(shippingUsecase)
}
