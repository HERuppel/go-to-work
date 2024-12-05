package repositories

import (
	"context"
	"errors"
	"go-to-work/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepositoryInterface interface {
	GetUser(ctx context.Context, tx pgx.Tx, id uint64) (*models.User, error)
	GetUserByEmail(ctx context.Context, tx pgx.Tx, email string) (*models.User, error)
}

type UserRepository struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (userRepository *UserRepository) GetUser(ctx context.Context, tx pgx.Tx, id uint64) (*models.User, error) {
	var user models.User

	query := `SELECT 
			u.id,
			u.name,
			u.email,
			u.cpf,
			u.birthdate,
			u.role,
			u.created_at,
			u.updated_at,
			a.id AS address_id,
			a.country,
			a.uf,
			a.city,
			a.street,
			a.zipcode,
			a.district,
			a.complement
		FROM users u 
		INNER JOIN addresses a ON u.address_id = a.id 
		WHERE u.id = $1`
	err := tx.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Cpf,
		&user.Birthdate,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Address.ID,
		&user.Address.Country,
		&user.Address.Uf,
		&user.Address.City,
		&user.Address.Street,
		&user.Address.Zipcode,
		&user.Address.District,
		&user.Address.Complement,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (userRepository *UserRepository) GetUserByEmail(ctx context.Context, tx pgx.Tx, email string) (*models.User, error) {
	var user models.User

	query := `SELECT 
			u.id,
			u.name,
			u.email,
			u.password,
			u.cpf,
			u.birthdate,
			u.pin_code,
			u.role,
			u.created_at,
			u.updated_at,
			a.id AS address_id,
			a.country,
			a.uf,
			a.city,
			a.street,
			a.zipcode,
			a.district,
			a.complement
		FROM users u 
		INNER JOIN addresses a ON u.address_id = a.id 
		WHERE u.email = $1`
	err := tx.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Cpf,
		&user.Birthdate,
		&user.PinCode,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Address.ID,
		&user.Address.Country,
		&user.Address.Uf,
		&user.Address.City,
		&user.Address.Street,
		&user.Address.Zipcode,
		&user.Address.District,
		&user.Address.Complement,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
