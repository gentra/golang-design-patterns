package internal

import "github.com/gentra/golang-design-patterns/chainofresponsibility/internal/entity"

type Checkout interface {
	Execute(cart *entity.Cart) error
	SetNext(co Checkout)
}
