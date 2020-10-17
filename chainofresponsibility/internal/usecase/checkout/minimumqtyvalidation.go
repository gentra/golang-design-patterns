package checkout

import (
	"log"

	"github.com/gentra/golang-design-patterns/chainofresponsibility/internal"
	"github.com/gentra/golang-design-patterns/chainofresponsibility/internal/entity"
)

type MinimumQtyValidation struct {
	next internal.Checkout
}

func NewMinimumQtyValidation() *MinimumQtyValidation {
	return &MinimumQtyValidation{}
}

func (v *MinimumQtyValidation) Execute(*entity.Cart) error {
	log.Println("Validating minimum quantity of purchase...")

	if v.next == nil {
		return nil
	}
	return v.next.Execute(nil)
}

func (v *MinimumQtyValidation) SetNext(co internal.Checkout) {
	v.next = co
}
