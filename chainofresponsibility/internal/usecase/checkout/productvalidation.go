package checkout

import (
	"log"

	"github.com/gentra/golang-design-patterns/chainofresponsibility/internal"
	"github.com/gentra/golang-design-patterns/chainofresponsibility/internal/entity"
)

type ProductValidation struct {
	next internal.Checkout
}

func NewProductValidation() *ProductValidation {
	return &ProductValidation{}
}

func (v *ProductValidation) Execute(*entity.Cart) error {
	log.Println("Validating product's eligibility...")

	if v.next == nil {
		return nil
	}
	return v.next.Execute(nil)
}

func (v *ProductValidation) SetNext(co internal.Checkout) {
	v.next = co
}
