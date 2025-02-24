package external

import (
	"hacktiv/model"
	"net/http"
)

type rajaOngkirAPI struct{}

func NewRajaOngkirAPI() model.IExternalShippingService {
	return &rajaOngkirAPI{}
}

func (r rajaOngkirAPI) GetShippingCost(origin string, destination string, weight int) float64 {
	// TODO: implement
	http.NewRequest()
	return 0
}
