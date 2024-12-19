package repositories

import (
	"context"
	"go-to-work/internal/models"

	"github.com/jackc/pgx/v5"
)

type AddressRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, address models.Address) (models.Address, error)
}

type AddressRepository struct{}

func NewAddressRepository() AddressRepositoryInterface {
	return &AddressRepository{}
}

func (addressRepository *AddressRepository) Create(ctx context.Context, tx pgx.Tx, address models.Address) (models.Address, error) {
	query := `
		INSERT INTO addresses(country, uf, city, street, zipcode, district, complement)
			VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;
	`

	err := tx.QueryRow(
		ctx,
		query,
		address.Country,
		address.Uf,
		address.City,
		address.Street,
		address.Zipcode,
		address.District,
		address.Complement,
	).Scan(&address.ID)
	if err != nil {
		return models.Address{}, err
	}

	return address, nil
}
