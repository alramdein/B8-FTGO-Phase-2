package usecase

import (
	"hacktiv/model"
	"hacktiv/repository"
)

type IShippingUsecase interface {
	GetShippingCost() float64
}

type shippingUsecase struct {
	shippingRepo        repository.IShippingRepository
	externalShippingAPI model.IExternalShippingService
}

func NewShippingUsecase(shippingRepo repository.IShippingRepository,
	externalShippingAPI model.IExternalShippingService) IShippingUsecase {
	return &shippingUsecase{
		shippingRepo: shippingRepo,
	}
}

func (s shippingUsecase) GetShippingCost() float64 {
	return s.externalShippingAPI.GetShippingCost("Jakarta", "Bandung", 100) // sesuaikan
}

// 