package entity

import "time"

type SalesReport struct {
	Header    string
	SubHeader string
	Content   string
	Footer    string
	Orders    []Order
	CreatedAt time.Time
}
