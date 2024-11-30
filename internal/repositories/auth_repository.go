package repositories

import (
	"context"
	"go-to-work/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		pool: pool,
	}
}

func (authRepository *AuthRepository) SignUp(ctx context.Context, user models.User) (models.User, error) {
	query := `
		INSERT INTO users (name, email, password, cpf, birthdate, pin_code, address_id, role)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;
	`

	err := authRepository.pool.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Cpf,
		user.Birthdate,
		user.PinCode,
		user.Address.ID,
		user.Role,
	).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
