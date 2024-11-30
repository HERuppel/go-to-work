package models

import (
	"errors"
)

type Address struct {
	ID         uint64 `json:"id"`
	Country    string `json:"country"`
	Uf         string `json:"uf"`
	City       string `json:"city"`
	Street     string `json:"street"`
	Zipcode    string `json:"zipcode"`
	District   string `json:"district,omitempty"`
	Complement string `json:"complement,omitempty"`
}

func (address *Address) Validate() error {
	if address.Country == "" {
		return errors.New("COUNTRY_REQUIRED")
	}
	if address.Uf == "" {
		return errors.New("UF_REQUIRED")
	}
	if address.City == "" {
		return errors.New("CITY_REQUIRED")
	}
	if address.Street == "" {
		return errors.New("STREET_REQUIRED")
	}
	if address.Zipcode == "" {
		return errors.New("ZIPCODE_REQUIRED")
	}

	return nil
}
