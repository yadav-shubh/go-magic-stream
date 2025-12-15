package controllers

import (
	"net/http"
	"strconv"

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

func (c *MovieController) GetMovies(ctx *gin.Context) {
	utils.Log.Info("GetMovies invoked")
	// pagination (1 and 20 if not present), search and genre filter
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "20"))
	search := ctx.Query("search")
	genre := ctx.Query("genre")

	movies, err := c.movieService.GetMovies(ctx.Request.Context(), page, size, search, genre)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewApiResponseNoData(err.Error(), http.StatusBadRequest))
		return
	}
	response := utils.NewApiResponse(movies, "Movie list", http.StatusOK)
	ctx.JSON(http.StatusOK, response)
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

func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var movie models.MovieDTO
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewApiResponseNoData(err.Error(), http.StatusBadRequest))
		return
	} else if err := utils.GetValidator().Struct(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewApiResponseNoData("Validation failed", http.StatusBadRequest))
		return
	}

	response, err := c.movieService.CreateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewApiResponseNoData(err.Error(), http.StatusBadRequest))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
