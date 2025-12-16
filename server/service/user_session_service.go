package service

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/models"
	"github.com/yadav-shubh/go-magic-stream/repository"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.uber.org/zap"
)

type UserSessionService struct {
	userSessionRepository *repository.UserSessionRepository
}

func NewUserSessionService(userSessionRepository *repository.UserSessionRepository) *UserSessionService {
	return &UserSessionService{userSessionRepository: userSessionRepository}
}

func (s *UserSessionService) CreateUserSession(ctx context.Context, session models.UserSession) (*utils.ApiResponse, error) {
	s.DeactivateAllUserSessions(ctx, session.UserID)
	err := s.userSessionRepository.Create(ctx, session)
	if err != nil {
		return nil, err
	}
	return utils.NewApiResponse([]interface{}{session}, "User session created successfully", http.StatusOK), nil

}

func (s *UserSessionService) DeactivateAllUserSessions(ctx context.Context, userId string) {
	s.userSessionRepository.UpdateAll(ctx, userId)
}

func (s *UserSessionService) FindAllByUserId(ctx context.Context, userId string) (*[]models.UserSession, error) {
	return s.userSessionRepository.FindAllByUserId(ctx, userId)
}

func (s *UserSessionService) ValidateUserSessionByRefreshToken(c *gin.Context, req models.RefreshTokenRequest) (*models.UserSession, error) {
	utils.Log.Info("GetUserSession", zap.String("refresh_token", req.RefreshToken))
	userSession, err := s.userSessionRepository.FindByRefreshToken(c, req.RefreshToken)
	if err != nil {
		utils.Log.Error("GetUserSession", zap.Error(err))
		return nil, err
	}
	return userSession, nil
}
