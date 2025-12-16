package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/models"
	"github.com/yadav-shubh/go-magic-stream/service"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.uber.org/zap"
)

type ContactController struct {
	contactService *service.ContactService
}

func NewContactController(contactService *service.ContactService) *ContactController {
	return &ContactController{contactService: contactService}
}

func (c *ContactController) RegisterRoutes(router *gin.Engine) {
	rg := router.Group("/api/contacts")
	{
		rg.POST("", c.CreateCreate)
	}
}

func (c *ContactController) CreateCreate(ctx *gin.Context) {
	utils.Log.Info("CreateCreate::invoked")
	var contact models.ContactMessageDTO
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewApiResponseNoData(err.Error(), http.StatusBadRequest))
		return
	} else if err := utils.GetValidator().Struct(&contact); err != nil {
		utils.Log.Info("CreateCreate::validation err", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, utils.NewApiResponseNoData("Validation failed", http.StatusBadRequest))
		return
	}

	clientIP := ctx.ClientIP()
	userAgent := ctx.Request.Header.Get("User-Agent")

	response, err := c.contactService.CreateContact(ctx, &contact, clientIP, userAgent)
	if err != nil {
		utils.Log.Info("CreateCreate::contact err", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, utils.NewApiResponseNoData(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
