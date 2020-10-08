package internal

import "github.com/gentra/legosamples/chainofresponsibility/internal/entity"

type Checkout interface {
	Execute(cart *entity.Cart) error
	SetNext(co Checkout)
}
