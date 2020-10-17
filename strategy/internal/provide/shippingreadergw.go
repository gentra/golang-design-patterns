package provide

import (
	"github.com/gentra/golang-design-patterns/strategy/internal"
	"github.com/gentra/golang-design-patterns/strategy/internal/constant"
	"github.com/gentra/golang-design-patterns/strategy/internal/gateway/shippingreadergw"
)

func ShippingReaderGw(partner constant.ShippingPartner) internal.ShippingReader {
	switch partner {
	case constant.ShippingPartnerJNE:
		return shippingreadergw.NewJNE()
	case constant.ShippingPartnerDHL:
		return shippingreadergw.NewDHL()
	case constant.ShippingPartnerGojek:
		return shippingreadergw.NewGojek()
	default:
		return nil
	}
}
