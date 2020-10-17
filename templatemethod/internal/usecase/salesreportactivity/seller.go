package salesreportactivity

import (
	"log"
	"time"

	"github.com/gentra/legosamples/templatemethod/internal/entity"
)

type Seller struct {
}

func NewSeller() *Seller {
	return &Seller{}
}

func (s *Seller) FetchOrderData() ([]entity.Order, error) {
	log.Println("Fetching order-data from application database")
	return []entity.Order{{ID: "1"}}, nil
}

func (s *Seller) CompileReport(orders []entity.Order) (entity.SalesReport, error) {
	log.Println("Compiling report with generic format for seller")
	return entity.SalesReport{
		Header:    "Sales Report for Your Shop - Oct 2020",
		SubHeader: "October 2020",
		Content:   "Sales report content for Seller",
		Footer:    "Sales Report - ShopName",
		Orders:    orders,
		CreatedAt: time.Now(),
	}, nil
}

func (s *Seller) StoreReportData(report entity.SalesReport) error {
	log.Println("Storing report data to application's database")
	return nil
}

func (s *Seller) SendReport(report entity.SalesReport) error {
	log.Println("Checking seller's email preferences for sales report")
	log.Println("Sending report to seller's email")
	return nil
}
