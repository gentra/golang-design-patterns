package main

import (
	"log"

	"github.com/gentra/golang-design-patterns/chainofresponsibility/internal/constant"
	"github.com/gentra/golang-design-patterns/chainofresponsibility/internal/entity"
	"github.com/gentra/golang-design-patterns/chainofresponsibility/internal/provide"
)

func main() {
	// Let's just pretend we get this checkout-type from user-type, config, or somewhere else
	checkoutType := constant.CheckoutDistributor
	co := provide.Checkout(checkoutType)

	err := co.Execute(&entity.Cart{})
	if err != nil {
		log.Println(err)
	}
}
