package shippingreadergw

import (
	"log"

	"github.com/gentra/golang-design-patterns/strategy/internal/entity"
)

type Gojek struct {
}

func NewGojek() *Gojek {
	return &Gojek{}
}

func (g *Gojek) GetShippingRate(route entity.ShipmentPlan) (*entity.ShippingRate, error) {
	// Let's just pretend we're fetching data from JNE here
	log.Println("Fetching shipping rate from Gojek")
	return &entity.ShippingRate{
		Description: "Gojek",
		Amount:      20000,
	}, nil
}
