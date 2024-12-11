package repositories

import (
	"context"
	"errors"
	"go-to-work/internal/models"

	"github.com/jackc/pgx/v5"
)

type AuthRepositoryInterface interface {
	SignUp(ctx context.Context, tx pgx.Tx, user models.User) (models.User, error)
	ConfirmAccount(ctx context.Context, tx pgx.Tx, email string) error
}

type AuthRepository struct {
}

func NewAuthRepository() AuthRepositoryInterface {
	return &AuthRepository{}
}

func (authRepository *AuthRepository) SignUp(ctx context.Context, tx pgx.Tx, user models.User) (models.User, error) {
	query := `
		INSERT INTO users (name, email, password, cpf, birthdate, pin_code, address_id, role)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;
	`

	err := tx.QueryRow(
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

func (authRepository *AuthRepository) ConfirmAccount(ctx context.Context, tx pgx.Tx, email string) error {
	query := `
		UPDATE users
			SET pin_code = NULL
		WHERE email = $1;
	`

	cmdTag, err := tx.Exec(ctx, query, email)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("FAILED_TO_UPDATED")
	}

	return nil
}
