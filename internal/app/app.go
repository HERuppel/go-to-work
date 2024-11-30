package app

import (
	"go-to-work/internal/controllers"
	"go-to-work/internal/database"
	"go-to-work/internal/repositories"
	usecases "go-to-work/internal/useCases"
)

type AppContainer struct {
	UserController *controllers.UserController
	AuthController *controllers.AuthController
}

func NewAppContainer() (*AppContainer, error) {
	pool, err := database.NewDatabasePool()
	if err != nil {
		return nil, err
	}

	//Address
	addressRepository := repositories.NewAddressRepository(pool)

	// Auth
	authRepository := repositories.NewAuthRepository(pool)
	authUseCase := usecases.NewAuthUseCase(authRepository, addressRepository)
	authController := controllers.NewAuthController(authUseCase)

	// User
	userRepository := repositories.NewUserRepository(pool)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUseCase)

	return &AppContainer{
		UserController: userController,
		AuthController: authController,
	}, nil
}
