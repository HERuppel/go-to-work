package models

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
