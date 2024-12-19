package app

import (
	"go-to-work/internal/config"
	"go-to-work/internal/controllers"
	"go-to-work/internal/database"
	"go-to-work/internal/repositories"
	"go-to-work/internal/services"
	usecases "go-to-work/internal/useCases"
	"log"
	"path/filepath"
)

type AppContainer struct {
	UserController *controllers.UserController
	AuthController *controllers.AuthController
	JobController  *controllers.JobController
}

func NewAppContainer() (*AppContainer, error) {
	pool, err := database.NewDatabasePool()
	if err != nil {
		return nil, err
	}

	templatePath, err := filepath.Abs("internal/templates")
	if err != nil {
		log.Fatalf("FAILED_TO_RESOLVE_TEMPLATE_PATH: %v", err)
	}

	emailService := services.NewEmailService(
		config.SmtpHost,
		config.SmtpPort,
		config.SmtpEmail,
		config.SmtpPassword,
		config.SmtpEmail,
		templatePath,
	)

	authRepository := repositories.NewAuthRepository()
	addressRepository := repositories.NewAddressRepository()
	userRepository := repositories.NewUserRepository()

	// Auth
	authUseCase := usecases.NewAuthUseCase(pool, authRepository, addressRepository, userRepository, emailService)
	authController := controllers.NewAuthController(authUseCase)

	// User
	userUseCase := usecases.NewUserUseCase(pool, userRepository)
	userController := controllers.NewUserController(userUseCase)

	// Job
	jobRepository := repositories.NewJobRepository()
	jobUseCase := usecases.NewJobUseCase(pool, jobRepository)
	jobController := controllers.NewJobController(jobUseCase)

	return &AppContainer{
		UserController: userController,
		AuthController: authController,
		JobController:  jobController,
	}, nil
}
