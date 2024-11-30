package usecases

import (
	"context"
	"go-to-work/internal/models"
	"go-to-work/internal/repositories"
	"go-to-work/internal/security"
	"strconv"
)

type AuthUseCase struct {
	authRepository    *repositories.AuthRepository
	addressRepository *repositories.AddressRepository
}

func NewAuthUseCase(authRepository *repositories.AuthRepository, addressRepository *repositories.AddressRepository) *AuthUseCase {
	return &AuthUseCase{
		authRepository:    authRepository,
		addressRepository: addressRepository,
	}
}

func (authUseCase *AuthUseCase) SignUp(ctx context.Context, user models.User) (models.User, error) {
	if err := user.Validate(); err != nil {
		return models.User{}, err
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hashedPassword)

	pinCode := security.GeneratePinCode()

	user.PinCode = strconv.Itoa(pinCode)

	user.Address, err = authUseCase.addressRepository.Create(ctx, user.Address)
	if err != nil {
		return models.User{}, err
	}

	return authUseCase.authRepository.SignUp(ctx, user)
}
