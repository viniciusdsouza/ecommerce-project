package models

type ProductUser struct {
	ProductId   string   `json:"id" validate:"uuid"`
	ProductName *string  `json:"product_name" validate:"required,gte=0,lte=255"`
	Price       *float64 `json:"price"`
	Rating      *uint8   `json:"rating"`
	Image       *string  `json:"image"`
}
