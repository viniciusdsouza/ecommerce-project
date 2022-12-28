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
