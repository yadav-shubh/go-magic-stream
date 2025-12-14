package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/models"
	"github.com/yadav-shubh/go-magic-stream/service"
	"github.com/yadav-shubh/go-magic-stream/utils"
)

type MovieController struct {
	movieService *service.MovieService
}

func NewMovieController(ms *service.MovieService) *MovieController {
	return &MovieController{
		movieService: ms,
	}
}

func (c *MovieController) RegisterRoutes(router *gin.Engine) {
	rg := router.Group("/api/movies")
	{
		rg.GET("", c.GetMovies)
		rg.GET("/:imdb_id", c.FindMovieByImdbID)
		rg.POST("", c.CreateMovie)
	}
}

func (c *MovieController) GetMovies(context *gin.Context) {
	movies, err := c.movieService.GetMovies()

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewApiResponseNoData(err.Error(), http.StatusBadRequest))
		return
	}
	response := utils.NewApiResponse(movies, "Movie list", http.StatusOK)
	context.JSON(http.StatusOK, response)
}

func (c *MovieController) FindMovieByImdbID(context *gin.Context) {
	imdbID := context.Param("imdb_id")
	movie, err := c.movieService.FindMovieByImdbID(imdbID)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewApiResponseNoData(err.Error(), http.StatusBadRequest))
		return
	}

	response := utils.NewApiResponse([]*models.MovieDTO{movie}, "Movie found", http.StatusOK)
	context.JSON(http.StatusOK, response)
}

func (c *MovieController) CreateMovie(context *gin.Context) {
	var movie models.MovieDTO
	if err := context.ShouldBindJSON(&movie); err != nil {
		context.JSON(http.StatusBadRequest, utils.NewApiResponseNoData(err.Error(), http.StatusBadRequest))
		return
	} else if err := utils.GetValidator().Struct(&movie); err != nil {
		context.JSON(http.StatusBadRequest, utils.NewApiResponseNoData("Validation failed", http.StatusBadRequest))
		return
	}

	response, err := c.movieService.CreateMovie(&movie)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewApiResponseNoData(err.Error(), http.StatusBadRequest))
		return
	}

	context.JSON(http.StatusOK, response)
}
