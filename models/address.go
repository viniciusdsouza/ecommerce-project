package models

type Address struct {
	AddressId string  `json:"id" validate:"uuid"`
	Number    *uint   `json:"number" validate:"required"`
	Street    *string `json:"street" validate:"required,gte=0,lte=255"`
	City      *string `json:"city" validate:"required,gte=0,lte=255"`
	ZipCode   *string `json:"zipcode" validate:"required"`
}
