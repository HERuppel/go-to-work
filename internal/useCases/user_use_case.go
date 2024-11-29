package usecases

import (
	"context"
	"go-to-work/internal/models"
	"go-to-work/internal/repositories"
)

type UserUseCase struct {
	userRepository *repositories.UserRepository
}

func NewUserUseCase(userRepository *repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (userUseCase *UserUseCase) GetUser(ctx context.Context, id uint64) (*models.User, error) {
	user, err := userUseCase.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
