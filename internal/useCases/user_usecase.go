package usecases

import (
	"context"
	"go-to-work/internal/database"
	"go-to-work/internal/models"
	"go-to-work/internal/repositories"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserUseCase struct {
	pool           *pgxpool.Pool
	userRepository repositories.UserRepositoryInterface
}

func NewUserUseCase(pool *pgxpool.Pool, userRepository repositories.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{
		pool:           pool,
		userRepository: userRepository,
	}
}

func (userUseCase *UserUseCase) GetUser(ctx context.Context, id uint64) (*models.User, error) {
	user, err := database.WithTransaction(userUseCase.pool, ctx, func(tx pgx.Tx) (*models.User, error) {
		user, err := userUseCase.userRepository.GetUser(ctx, tx, id)
		if err != nil {
			return nil, err
		}

		return user, nil
	})

	if err != nil {
		return &models.User{}, nil
	}

	return user, nil
}
