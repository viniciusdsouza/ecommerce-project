package models

import "time"

type Order struct {
	OrderId       string        `json:"id" validate:"uuid"`
	OrderCart     []ProductUser `json:"order_list"`
	OrderedDate   time.Time     `json:"ordered_date"`
	Price         *float64      `json:"total_price"`
	Discount      *float64      `json:"discount"`
	PaymentMethod Payment       `json:"payment_method"`
}
