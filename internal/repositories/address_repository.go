package repositories

import (
	"context"
	"go-to-work/internal/models"

	"github.com/jackc/pgx/v5"
)

type AddressRepository struct {
	tx pgx.Tx
}

func NewAddressRepository(tx pgx.Tx) *AddressRepository {
	return &AddressRepository{
		tx: tx,
	}
}

func (addressRepository *AddressRepository) Create(ctx context.Context, address models.Address) (models.Address, error) {
	query := `
		INSERT INTO addresses(country, uf, city, street, zipcode, district, complement)
			VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;
	`

	err := addressRepository.tx.QueryRow(
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
