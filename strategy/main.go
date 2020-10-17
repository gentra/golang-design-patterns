package main

import (
	"log"

	"github.com/gentra/golang-design-patterns/strategy/internal/constant"
	"github.com/gentra/golang-design-patterns/strategy/internal/entity"
	"github.com/gentra/golang-design-patterns/strategy/internal/provide"
)

func main() {
	shippingPartner := constant.ShippingPartnerGojek
	shippingReader := provide.ShippingReaderGw(shippingPartner)

	rate, err := shippingReader.GetShippingRate(entity.ShipmentPlan{
		Origin:               "Jakarta",
		Destination:          "Bandung",
		ShipmentDurationDays: 3,
	})
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v", rate)
}
