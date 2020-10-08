package provide

import (
	"github.com/gentra/legosamples/chainofresponsibility/internal"
	"github.com/gentra/legosamples/chainofresponsibility/internal/usecase/checkout"
)

type CheckoutType int

const (
	CheckoutConsumer    CheckoutType = 1
	CheckoutDistributor CheckoutType = 2
	CheckoutBigPromo    CheckoutType = 3
)

func Checkout(typ CheckoutType) internal.Checkout {
	switch typ {
	case CheckoutConsumer:
		pay := checkout.NewPaymentPreparation() // sequence 2

		product := checkout.NewProductValidation() // sequence 1
		product.SetNext(pay)

		return product
	case CheckoutDistributor:
		pay := checkout.NewPaymentPreparation() // sequence 6

		product := checkout.NewProductValidation() // sequence 5
		product.SetNext(pay)

		location := checkout.NewLocationValidation() // sequence 4
		location.SetNext(product)

		shopeligible := checkout.NewShopEligibility() // sequence 3
		shopeligible.SetNext(location)

		kyc := checkout.NewKYCValidation() // sequence 2
		kyc.SetNext(shopeligible)

		minqty := checkout.NewMinimumQtyValidation() // sequence 1
		minqty.SetNext(kyc)

		return minqty
	case CheckoutBigPromo:
		pay := checkout.NewPaymentPreparation() // sequence 4

		product := checkout.NewProductValidation() // sequence 3
		product.SetNext(pay)

		fraudcheck := checkout.NewFraudCheck() // sequence 2
		fraudcheck.SetNext(product)

		promocheck := checkout.NewPromoValidation() // sequence 1
		promocheck.SetNext(fraudcheck)
		return promocheck
	default:
		return nil
	}
}
