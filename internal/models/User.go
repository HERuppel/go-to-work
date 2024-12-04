package models

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
)

type RoleType string

const (
	Admin     RoleType = "ADMIN"
	Recruiter RoleType = "RECRUITER"
	Candidate RoleType = "CANDIDATE"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Cpf       string    `json:"cpf"`
	Birthdate time.Time `json:"birthdate"`
	PinCode   *string   `json:"pin_code,omitempty"`
	Role      RoleType  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Address   Address   `json:"address"`
}

func (user *User) Validate() error {
	if user.Name == "" {
		return errors.New("NAME_REQUIRED")
	}
	if user.Cpf == "" {
		return errors.New("CPF_REQUIRED")
	}
	if user.Birthdate.IsZero() {
		return errors.New("BIRTHDATE_REQUIRED")
	}
	if user.Password == "" {
		return errors.New("PASSWORD_REQURIED")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("INVALID_EMAIL_FORMAT")
	}
	if user.Role == "" {
		return errors.New("ROLE_REQUIRED")
	}

	if err := user.Address.Validate(); err != nil {
		return err
	}

	return nil
}
