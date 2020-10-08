package checkout

import (
	"log"

	"github.com/gentra/legosamples/chainofresponsibility/internal"
	"github.com/gentra/legosamples/chainofresponsibility/internal/entity"
)

type FraudCheck struct {
	next internal.Checkout
}

func NewFraudCheck() *FraudCheck {
	return &FraudCheck{}
}

func (v *FraudCheck) Execute(*entity.Cart) error {
	log.Println("Checking transaction for fraud...")

	if v.next == nil {
		return nil
	}
	return v.next.Execute(nil)
}

func (v *FraudCheck) SetNext(co internal.Checkout) {
	v.next = co
}
