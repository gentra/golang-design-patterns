package checkout

import (
	"log"

	"github.com/gentra/legosamples/chainofresponsibility/internal"
	"github.com/gentra/legosamples/chainofresponsibility/internal/entity"
)

type ShopEligibility struct {
	next internal.Checkout
}

func NewShopEligibility() *ShopEligibility {
	return &ShopEligibility{}
}

func (v *ShopEligibility) Execute(*entity.Cart) error {
	log.Println("Checking if user is eligible to purchase on this shop...")

	if v.next == nil {
		return nil
	}
	return v.next.Execute(nil)
}

func (v *ShopEligibility) SetNext(co internal.Checkout) {
	v.next = co
}
