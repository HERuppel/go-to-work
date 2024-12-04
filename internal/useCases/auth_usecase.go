package usecases

import (
	"context"
	"errors"
	"fmt"
	"go-to-work/internal/authentication"
	"go-to-work/internal/models"
	"go-to-work/internal/repositories"
	"go-to-work/internal/security"
	"go-to-work/internal/services"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthUseCase struct {
	pool         *pgxpool.Pool
	emailService services.EmailService
}

func NewAuthUseCase(pool *pgxpool.Pool, emailService services.EmailService) *AuthUseCase {
	return &AuthUseCase{
		pool:         pool,
		emailService: emailService,
	}
}

func (authUseCase *AuthUseCase) SignUp(ctx context.Context, user models.User) (models.User, error) {
	tx, err := authUseCase.pool.Begin(ctx)
	if err != nil {
		return models.User{}, fmt.Errorf("ERROR_STARTING_TRANSACTION: %w", err)
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

	authRepository := repositories.NewAuthRepository(tx)
	addressRepository := repositories.NewAddressRepository(tx)

	if err := user.Validate(); err != nil {
		return models.User{}, err
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hashedPassword)

	pinCode := security.GeneratePinCode()

	*user.PinCode = strconv.Itoa(pinCode)

	user.Address, err = addressRepository.Create(ctx, user.Address)
	if err != nil {
		return models.User{}, err
	}

	createdUser, err := authRepository.SignUp(ctx, user)
	if err != nil {
		return models.User{}, err
	}

	if err = authUseCase.emailService.SendConfirmEmail(createdUser.Email, createdUser.Name, *createdUser.PinCode); err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (authUseCase *AuthUseCase) SignIn(ctx context.Context, email, password string) (models.User, string, error) {
	tx, err := authUseCase.pool.Begin(ctx)
	if err != nil {
		return models.User{}, "", fmt.Errorf("ERROR_STARTING_TRANSACTION: %w", err)
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

	userRepository := repositories.NewUserRepository(tx)

	userToCompare, err := userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, "", err
	}

	if userToCompare.PinCode != nil {
		return models.User{}, "", errors.New("NEED_TO_CONFIRM_ACCOUNT")
	}

	if err = security.VerifyPassword(userToCompare.Password, password); err != nil {
		return models.User{}, "", errors.New("WRONG_PASSWORD_OR_EMAIL")
	}

	authToken, err := authentication.CreateToken(userToCompare.ID)
	if err != nil {
		return models.User{}, "", errors.New("INTERNAL_ERROR")
	}

	return *userToCompare, authToken, nil
}
