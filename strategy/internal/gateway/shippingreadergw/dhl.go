package shippingreadergw

import (
	"log"

	"github.com/gentra/golang-design-patterns/strategy/internal/entity"
)

type DHL struct {
}

func NewDHL() *DHL {
	return &DHL{}
}

func (d *DHL) GetShippingRate(route entity.ShipmentPlan) (*entity.ShippingRate, error) {
	// Let's just pretend we're fetching data from JNE here
	log.Println("Fetching shipping rate from DHL")
	return &entity.ShippingRate{
		Description: "DHL",
		Amount:      19000,
	}, nil
}
