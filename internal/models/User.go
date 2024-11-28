package models

import "time"

type RoleType string

const (
	Admin     RoleType = "ADMIN"
	Recruiter RoleType = "RECRUITER"
	Candidate RoleType = "CANDIDATE"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Cpf       string    `json:"cpf"`
	Birthdate time.Time `json:"birthdate"`
	PinCode   string    `json:"pin_code,omitempty"`
	AddressID uint      `json:"address_id"`
	Role      RoleType  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
