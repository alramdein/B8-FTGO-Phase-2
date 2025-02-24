package repository

type IShippingRepository interface {
	GetShippingCost(origin string, destination string, weight int) float64
}

type shippingRepository struct{}

func NewShippingRepository() IShippingRepository {
	return &shippingRepository{}
}

func (s shippingRepository) GetShippingCost(origin string, destination string, weight int) float64 {
	// TODO: implement
	return 0
}
