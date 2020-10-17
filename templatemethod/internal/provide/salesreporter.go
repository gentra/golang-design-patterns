package provide

import (
	"github.com/gentra/golang-design-patterns/templatemethod/internal"
	"github.com/gentra/golang-design-patterns/templatemethod/internal/usecase/salesreporter"
)

func SalesReporter() internal.SalesReporter {
	return salesreporter.NewSalesReporter()
}
