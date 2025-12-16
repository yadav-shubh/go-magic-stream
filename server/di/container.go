package di

import (
	"github.com/yadav-shubh/go-magic-stream/api_server/controllers"
	"github.com/yadav-shubh/go-magic-stream/database"
	"github.com/yadav-shubh/go-magic-stream/repository"
	"github.com/yadav-shubh/go-magic-stream/service"
	"github.com/yadav-shubh/go-magic-stream/utils"
)

type Container struct {
	MovieController    *controllers.MovieController
	AuthController     *controllers.AuthController
	UserSessionService *service.UserSessionService
	ContactController  *controllers.ContactController
}

func NewContainer() *Container {
	db := database.ConnectMongo()

	// Utilities
	kindeUtility := utils.NewKindeUtility()

	// Repositories
	userSessionRepository := repository.NewUserSessionRepository(db)
	contactRepository := repository.NewContactRepository(db)

	// Services
	userSessionService := service.NewUserSessionService(userSessionRepository)
	authService := service.NewAuthService(db, kindeUtility, userSessionService)
	movieService := service.NewMovieService(db)
	contactService := service.NewContactService(contactRepository)

	// Controllers
	authController := controllers.NewAuthController(authService)
	movieController := controllers.NewMovieController(movieService)
	contactController := controllers.NewContactController(contactService)

	return &Container{
		MovieController:    movieController,
		AuthController:     authController,
		UserSessionService: userSessionService,
		ContactController:  contactController,
	}
}
