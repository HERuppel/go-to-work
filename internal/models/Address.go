package models

type Address struct {
	ID         uint   `json:"id"`
	Country    string `json:"country"`
	Uf         string `json:"uf"`
	Street     string `json:"street"`
	Zipcode    string `json:"zipcode"`
	District   string `json:"district,omitempty"`
	Complement string `json:"complement,omitempty"`
}
