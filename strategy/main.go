package main

import (
	"log"

	"github.com/gentra/legosamples/strategy/internal/constant"
	"github.com/gentra/legosamples/strategy/internal/entity"
	"github.com/gentra/legosamples/strategy/internal/provide"
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
