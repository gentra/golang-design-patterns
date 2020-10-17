package provide

import (
	"github.com/gentra/golang-design-patterns/templatemethod/internal"
	"github.com/gentra/golang-design-patterns/templatemethod/internal/constant"
	"github.com/gentra/golang-design-patterns/templatemethod/internal/usecase/salesreportactivity"
)

func SalesReporterActivity(recipient constant.SalesReportRecipient) internal.SalesReportActivity {
	switch recipient {
	case constant.SalesReportForSeller:
		return salesreportactivity.NewSeller()
	case constant.SalesReportForBusinessDevDept:
		return salesreportactivity.NewBusinessDev()
	case constant.SalesReportForFinanceDept:
		return salesreportactivity.NewFinance()
	case constant.SalesReportForInvestor:
		return salesreportactivity.NewInvestor()
	default:
		return nil
	}

}
