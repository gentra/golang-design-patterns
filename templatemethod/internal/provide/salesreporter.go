package provide

import (
	"github.com/gentra/legosamples/templatemethod/internal"
	"github.com/gentra/legosamples/templatemethod/internal/usecase/salesreporter"
)

func SalesReporter() internal.SalesReporter {
	return salesreporter.NewSalesReporter()
}
