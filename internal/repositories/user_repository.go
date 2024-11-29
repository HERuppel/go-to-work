package repositories

import (
	"context"
	"errors"
	"go-to-work/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (userRepository *UserRepository) GetUser(ctx context.Context, id uint64) (*models.User, error) {
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
			a.street,
			a.zipcode,
			a.district,
			a.complement
		FROM users u 
		INNER JOIN addresses a ON u.address_id = a.id WHERE u.id = $1`
	err := userRepository.pool.QueryRow(ctx, query, id).Scan(
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
