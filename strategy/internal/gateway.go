package internal

import "github.com/gentra/legosamples/strategy/internal/entity"

type ShippingReader interface {
	GetShippingRate(route entity.ShipmentPlan) (*entity.ShippingRate, error)
}
