package api_server

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/api_server/controllers"
	"github.com/yadav-shubh/go-magic-stream/api_server/middleware"
	"github.com/yadav-shubh/go-magic-stream/di"
	"github.com/yadav-shubh/go-magic-stream/service"
)

var container *di.Container
var once sync.Once

type GinServer struct {
	movieController    *controllers.MovieController
	authController     *controllers.AuthController
	userSessionService *service.UserSessionService
	contactController  *controllers.ContactController
}

func NewGinServer() *GinServer {
	once.Do(func() {
		container = di.NewContainer()
	})
	return &GinServer{
		movieController:    container.MovieController,
		authController:     container.AuthController,
		userSessionService: container.UserSessionService,
		contactController:  container.ContactController,
	}
}

func (s *GinServer) GetHandler() http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.JWTAuth())
	router.Use(middleware.RateLimit())
	router.Use(middleware.SecurityHeaders())

	s.RegisterRoutes(router)

	return router
}

func (s *GinServer) RegisterRoutes(router *gin.Engine) {
	s.movieController.RegisterRoutes(router)
	s.authController.RegisterRoutes(router)
	s.contactController.RegisterRoutes(router)
}
