package external

import (
	"hacktiv/model"
)

type rajaOngkirAPI struct{}

func NewRajaOngkirAPI() model.IExternalShippingService {
	return &rajaOngkirAPI{}
}

func (r rajaOngkirAPI) GetShippingCost(origin string, destination string, weight int) float64 {
	// TODO: implement call API http disini
	return 0
}
