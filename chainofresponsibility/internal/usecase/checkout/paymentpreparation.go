package checkout

import (
	"log"

	"github.com/gentra/legosamples/chainofresponsibility/internal"
	"github.com/gentra/legosamples/chainofresponsibility/internal/entity"
)

type PaymentPreparation struct {
	next internal.Checkout
}

func NewPaymentPreparation() *PaymentPreparation {
	return &PaymentPreparation{}
}

func (v *PaymentPreparation) Execute(*entity.Cart) error {
	log.Println("Preparing cart for payment...")

	if v.next == nil {
		return nil
	}
	return v.next.Execute(nil)
}

func (v *PaymentPreparation) SetNext(co internal.Checkout) {
	v.next = co
}
