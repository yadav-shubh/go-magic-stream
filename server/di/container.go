package di

import (
	"github.com/yadav-shubh/go-magic-stream/api_server/controllers"
	"github.com/yadav-shubh/go-magic-stream/api_server/repository"
	"github.com/yadav-shubh/go-magic-stream/database"
	"github.com/yadav-shubh/go-magic-stream/service"
	"github.com/yadav-shubh/go-magic-stream/utils"
)

type Container struct {
	MovieController    *controllers.MovieController
	AuthController     *controllers.AuthController
	UserSessionService *service.UserSessionService
}

func NewContainer() *Container {
	db := database.ConnectMongo()

	// Utilities
	kindeUtility := utils.NewKindeUtility()

	// Repositories
	userSessionRepository := repository.NewUserSessionRepository(db)

	// Services
	userSessionService := service.NewUserSessionService(userSessionRepository)
	authService := service.NewAuthService(db, kindeUtility, userSessionService)
	movieService := service.NewMovieService(db)

	// Controllers
	authController := controllers.NewAuthController(authService)
	movieController := controllers.NewMovieController(movieService)

	return &Container{
		MovieController:    movieController,
		AuthController:     authController,
		UserSessionService: userSessionService,
	}
}
