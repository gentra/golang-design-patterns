package main

import (
	"log"

	"github.com/gentra/legosamples/chainofresponsibility/internal/entity"
	"github.com/gentra/legosamples/chainofresponsibility/internal/provide"
)

func main() {
	// Let's just pretent we get this checkout-type from user-type, config, or somewhere else
	checkoutType := provide.CheckoutDistributor
	co := provide.Checkout(checkoutType)

	err := co.Execute(&entity.Cart{})
	if err != nil {
		log.Println(err)
	}
}
