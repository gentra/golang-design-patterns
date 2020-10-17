package checkout

import (
	"log"

	"github.com/gentra/golang-design-patterns/chainofresponsibility/internal"
	"github.com/gentra/golang-design-patterns/chainofresponsibility/internal/entity"
)

type PromoValidation struct {
	next internal.Checkout
}

func NewPromoValidation() *PromoValidation {
	return &PromoValidation{}
}

func (v *PromoValidation) Execute(*entity.Cart) error {
	log.Println("Validating promotion...")

	if v.next == nil {
		return nil
	}
	return v.next.Execute(nil)
}

func (v *PromoValidation) SetNext(co internal.Checkout) {
	v.next = co
}
