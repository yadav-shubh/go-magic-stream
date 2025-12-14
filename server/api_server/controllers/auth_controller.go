package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/models"
	"github.com/yadav-shubh/go-magic-stream/service"
	"github.com/yadav-shubh/go-magic-stream/utils"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) RegisterRoutes(router *gin.Engine) {
	rg := router.Group("/api/auth")
	{
		rg.GET("/auth-info", c.AuthInfo)
		rg.GET("/authenticate", c.Authenticate)
		rg.POST("/refresh-token", c.RefreshToken)
	}
}

func (c *AuthController) AuthInfo(context *gin.Context) {
	authInfoResponse, err := c.authService.AuthInfo()

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewApiResponseNoData(err.Error(), http.StatusBadRequest))
		return
	}
	// Disable HTML escaping
	context.Writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(context.Writer)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(authInfoResponse)
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.NewApiResponseNoData(err.Error(), http.StatusInternalServerError))
		return
	}
}

func (c *AuthController) Authenticate(context *gin.Context) {
	code := context.Query("code")
	if code == "" {
		context.JSON(http.StatusBadRequest, utils.NewApiResponseNoData("code is required", http.StatusBadRequest))
		return
	}
	authenticate, err := c.authService.Authenticate(context, code)
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.NewApiResponseNoData(err.Error(), http.StatusInternalServerError))
		return
	}
	context.JSON(http.StatusOK, authenticate)
}

func (c *AuthController) RefreshToken(context *gin.Context) {
	var req models.RefreshTokenRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.authService.RefreshToken(context, req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.NewApiResponseNoData(err.Error(), http.StatusInternalServerError))
		return
	}
	context.JSON(http.StatusOK, response)
}
