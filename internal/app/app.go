package app

import (
	"go-to-work/internal/controllers"
	"go-to-work/internal/database"
	"go-to-work/internal/repositories"
	usecases "go-to-work/internal/useCases"
)

type AppContainer struct {
	UserController *controllers.UserController
}

func NewAppContainer(databaseUrl string) (*AppContainer, error) {
	pool, err := database.NewDatabasePool()
	if err != nil {
		return nil, err
	}

	userRepository := repositories.NewUserRepository(pool)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUseCase)

	return &AppContainer{
		UserController: userController,
	}, nil
}
