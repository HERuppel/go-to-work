package usecases

import (
	"context"
	"fmt"
	"go-to-work/internal/models"
	"go-to-work/internal/repositories"

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
	tx, err := userUseCase.pool.Begin(ctx)
	if err != nil {
		return &models.User{}, fmt.Errorf("ERROR_STARTING_TRANSACTION: %w", err)
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				fmt.Println("ROLLBACK_ERROR")
			}
		} else {
			if commitErr := tx.Commit(ctx); commitErr != nil {
				fmt.Println("COMMIT_ERROR")
			}
		}
	}()

	user, err := userUseCase.userRepository.GetUser(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
