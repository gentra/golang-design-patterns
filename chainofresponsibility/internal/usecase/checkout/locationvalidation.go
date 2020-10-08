package checkout

import (
	"log"

	"github.com/gentra/legosamples/chainofresponsibility/internal"
	"github.com/gentra/legosamples/chainofresponsibility/internal/entity"
)

type LocationValidation struct {
	next internal.Checkout
}

func NewLocationValidation() *LocationValidation {
	return &LocationValidation{}
}

func (v *LocationValidation) Execute(*entity.Cart) error {
	log.Println("Validating shop location is within range of the buyer...")

	if v.next == nil {
		return nil
	}
	return v.next.Execute(nil)
}

func (v *LocationValidation) SetNext(co internal.Checkout) {
	v.next = co
}
