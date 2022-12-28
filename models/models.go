package models

import "time"

type User struct {
	Id             string        `json:"id" validate:"uuid"`
	FirstName      *string       `json:"first_name" validate:"required,gte=0,lte=255"`
	LastName       *string       `json:"last_name" validate:"required,gte=0,lte=255"`
	Password       *string       `json:"password" validate:"required,gte=6"`
	Email          *string       `json:"email" validate:"email,required"`
	Phone          *string       `json:"phone" validate:"required"`
	Token          string        `json:"token" validate:"required"`
	RefreshToken   string        `json:"refresh_token"`
	CreationDate   time.Time     `json:"creation_date"`
	UpdateDate     time.Time     `json:"update_date"`
	UserId         string        `json:"user_id"`
	UserCart       []ProductUser `json:"user_cart"`
	AddressDetails []Address     `json:"address_details"`
	OrderStatus    []Order       `json:"order_status"`
}

type Address struct {
	AddressId string `json:"id" validate:"uuid"`
	Number    *uint `json:"number" validate:"required"`
	Street    *string `json:"street" validate:"required,gte=0,lte=255"`
	City      *string `json:"city" validate:"required,gte=0,lte=255"`
	ZipCode   *string `json:"zipcode" validate:"required"`
}

type Order struct {
	OrderId       string `json:"id" validate:"uuid"`
	OrderCart     []ProductUser `json:"order_list"`
	OrderedDate   time.Time `json:"ordered_date"`
	Price         *float64 `json:"total_price"`
	Discount      *float64 `json:"discount"`
	PaymentMethod Payment `json:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital"`
	Cod     bool `json:"cod"`
}

type Product struct {
	ProductId   string `json:"id" validate:"uuid"`
	ProductName *string `json:"product_name" validate:"required,gte=0,lte=255"`
	Price       *float64 `json:"price"`
	Rating      *uint8 `json:"rating"`
	Image       *string `json:"image"`
}

type ProductUser struct {
	ProductId   string `json:"id" validate:"uuid"`
	ProductName *string `json:"product_name" validate:"required,gte=0,lte=255"`
	Price       *float64 `json:"price"`
	Rating      *uint8 `json:"rating"`
	Image       *string `json:"image"`
}
