package internal

import "github.com/gentra/golang-design-patterns/strategy/internal/entity"

type ShippingReader interface {
	GetShippingRate(route entity.ShipmentPlan) (*entity.ShippingRate, error)
}
