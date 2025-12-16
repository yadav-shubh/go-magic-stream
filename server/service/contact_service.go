package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/models"
	"github.com/yadav-shubh/go-magic-stream/repository"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.uber.org/zap"
)

type ContactService struct {
	contactRepository *repository.ContactRepository
}

func NewContactService(contactRepository *repository.ContactRepository) *ContactService {
	return &ContactService{contactRepository: contactRepository}
}

func (s *ContactService) CreateContact(ctx *gin.Context, contact *models.ContactMessageDTO, clientIp string, userAgent string) (*utils.ApiResponse, error) {
	contactMessage := models.ContactMessage{
		Name:      contact.Name,
		Email:     contact.Email,
		Phone:     contact.Phone,
		Query:     contact.Query,
		IPAddress: clientIp,
		UserAgent: userAgent,
		CreatedAt: time.Now(),
	}

	err := s.contactRepository.Create(ctx, contactMessage)
	if err != nil {
		utils.Log.Error("Error while creating the contact message", zap.Error(err))
		return nil, err
	}
	utils.Log.Info("Created contact message", zap.Any("contactMessage", contactMessage))
	return utils.NewApiResponseNoData("Created contact message", http.StatusOK), nil
}
