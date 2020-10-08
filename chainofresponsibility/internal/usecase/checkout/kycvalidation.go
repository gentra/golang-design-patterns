package checkout

import (
	"log"

	"github.com/gentra/legosamples/chainofresponsibility/internal"
	"github.com/gentra/legosamples/chainofresponsibility/internal/entity"
)

type KYCValidation struct {
	next internal.Checkout
}

func NewKYCValidation() *KYCValidation {
	return &KYCValidation{}
}

func (v *KYCValidation) Execute(*entity.Cart) error {
	log.Println("Validating user's KYC...")

	if v.next == nil {
		return nil
	}
	return v.next.Execute(nil)
}

func (v *KYCValidation) SetNext(co internal.Checkout) {
	v.next = co
}
