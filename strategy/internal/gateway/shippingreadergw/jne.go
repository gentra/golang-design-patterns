package shippingreadergw

import (
	"log"

	"github.com/gentra/legosamples/strategy/internal/entity"
)

type JNE struct {
}

func NewJNE() *JNE {
	return &JNE{}
}

func (j *JNE) GetShippingRate(route entity.ShipmentPlan) (*entity.ShippingRate, error) {
	// Let's just pretend we're fetching data from JNE here
	log.Println("Fetching shipping rate from JNE")
	return &entity.ShippingRate{
		Description: "JNE",
		Amount:      15000,
	}, nil
}
