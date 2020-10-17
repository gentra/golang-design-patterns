package constant

type SalesReportRecipient int

const (
	SalesReportForSeller          SalesReportRecipient = 1
	SalesReportForBusinessDevDept SalesReportRecipient = 2
	SalesReportForFinanceDept     SalesReportRecipient = 3
	SalesReportForInvestor        SalesReportRecipient = 4
)
