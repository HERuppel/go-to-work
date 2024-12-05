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
	pool              *pgxpool.Pool
	authRepository    repositories.AuthRepositoryInterface
	addressRepository repositories.AddressRepositoryInterface
	userRepository    repositories.UserRepositoryInterface
	emailService      services.EmailService
}

func NewAuthUseCase(pool *pgxpool.Pool, authRepository repositories.AuthRepositoryInterface, addressRepository repositories.AddressRepositoryInterface, userRepository repositories.UserRepositoryInterface, emailService services.EmailService) *AuthUseCase {
	return &AuthUseCase{
		pool:              pool,
		authRepository:    authRepository,
		addressRepository: addressRepository,
		userRepository:    userRepository,
		emailService:      emailService,
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

	if err := user.Validate(); err != nil {
		return models.User{}, err
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hashedPassword)

	pinCode := security.GeneratePinCode()

	user.PinCode = new(string)

	*user.PinCode = strconv.Itoa(pinCode)

	user.Address, err = authUseCase.addressRepository.Create(ctx, tx, user.Address)
	if err != nil {
		return models.User{}, err
	}

	createdUser, err := authUseCase.authRepository.SignUp(ctx, tx, user)
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

	userToCompare, err := authUseCase.userRepository.GetUserByEmail(ctx, tx, email)
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

func (authUseCase *AuthUseCase) ConfirmAccount(ctx context.Context, email, pinCode string) error {
	tx, err := authUseCase.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("ERROR_STARTING_TRANSACTION: %w", err)
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

	user, err := authUseCase.userRepository.GetUserByEmail(ctx, tx, email)
	if err != nil {
		return err
	}

	if user.PinCode == nil {
		return errors.New("ACCOUNT_ALREADY_CONFIRMED")
	}

	if *user.PinCode != pinCode {
		return errors.New("INVALID_PIN_CODE")
	}

	err = authUseCase.authRepository.ConfirmAccount(ctx, tx, email)
	if err != nil {
		return errors.New("INTERNAL_ERROR")
	}

	return nil
}
