package usecases

import (
	"context"
	"errors"
	"go-to-work/internal/authentication"
	"go-to-work/internal/database"
	"go-to-work/internal/models"
	"go-to-work/internal/repositories"
	"go-to-work/internal/security"
	"go-to-work/internal/services"
	"strconv"

	"github.com/jackc/pgx/v5"
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

type SignInResponse struct {
	User      models.User `json:"user"`
	AuthToken string      `json:"authToken"`
}

func (authUseCase *AuthUseCase) SignUp(ctx context.Context, user models.User) (models.User, error) {
	createdUser, err := database.WithTransaction(authUseCase.pool, ctx, func(tx pgx.Tx) (models.User, error) {
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
	})

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (authUseCase *AuthUseCase) SignIn(ctx context.Context, email, password string) (SignInResponse, error) {
	user, err := database.WithTransaction(authUseCase.pool, ctx, func(tx pgx.Tx) (SignInResponse, error) {
		userToCompare, err := authUseCase.userRepository.GetUserByEmail(ctx, tx, email)
		if err != nil {
			return SignInResponse{}, err
		}

		if userToCompare == nil {
			return SignInResponse{}, errors.New("USER_NOT_FOUND")
		}

		if userToCompare.PinCode != nil {
			return SignInResponse{}, errors.New("NEED_TO_CONFIRM_ACCOUNT")
		}

		if err = security.VerifyPassword(userToCompare.Password, password); err != nil {
			return SignInResponse{}, errors.New("WRONG_PASSWORD_OR_EMAIL")
		}

		authToken, err := authentication.CreateToken(userToCompare.ID)
		if err != nil {
			return SignInResponse{}, errors.New("INTERNAL_ERROR")
		}

		return SignInResponse{
			User:      *userToCompare,
			AuthToken: authToken,
		}, nil
	})

	if err != nil {
		return SignInResponse{}, err
	}

	return user, nil
}

func (authUseCase *AuthUseCase) ConfirmAccount(ctx context.Context, email, pinCode string) error {
	_, err := database.WithTransaction(authUseCase.pool, ctx, func(tx pgx.Tx) (interface{}, error) {
		user, err := authUseCase.userRepository.GetUserByEmail(ctx, tx, email)
		if err != nil {
			return nil, err
		}

		if user.PinCode == nil {
			return nil, errors.New("ACCOUNT_ALREADY_CONFIRMED")
		}

		if *user.PinCode != pinCode {
			return nil, errors.New("INVALID_PIN_CODE")
		}

		err = authUseCase.authRepository.ConfirmAccount(ctx, tx, email)
		if err != nil {
			return nil, errors.New("INTERNAL_ERROR")
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}
