package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yadav-shubh/go-magic-stream/config"
	"github.com/yadav-shubh/go-magic-stream/models"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type AuthService struct {
	db                 *mongo.Database
	kindeUtility       *utils.KindeUtility
	userSessionService *UserSessionService
}

func NewAuthService(db *mongo.Database, kindeUtility *utils.KindeUtility, userSessionService *UserSessionService) *AuthService {
	return &AuthService{db: db, kindeUtility: kindeUtility, userSessionService: userSessionService}
}

func (s *AuthService) AuthInfo() (*models.AuthInfoResponse, error) {

	ssoConfig := config.Get().SSO
	//	prepare url
	parse, err := url.Parse(ssoConfig.Issuer + ssoConfig.AuthUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid SSO configuration")
	}

	query := parse.Query()
	query.Add("client_id", ssoConfig.ClientId)
	query.Add("redirect_uri", ssoConfig.RedirectUri)
	query.Add("response_type", ssoConfig.ResponseType)
	query.Add("scope", ssoConfig.Scope)
	query.Add("state", uuid.New().String())

	parse.RawQuery = query.Encode()

	return &models.AuthInfoResponse{
		LoginUrl:  parse.String(),
		LogoutUrl: ssoConfig.LogoutUrl,
	}, nil
}

func (s *AuthService) Authenticate(context context.Context, code string) (*utils.ApiResponse, error) {
	utils.Log.Info("Authenticate", zap.String("code", code))

	apiResponse, err := s.kindeUtility.VerifyAuthCode(code)
	if err != nil {
		return nil, err
	}
	userEmail := s.kindeUtility.GetUserEmail(apiResponse.AccessToken)
	userEmail = strings.ToLower(userEmail)

	findOpts := options.FindOne().
		SetCollation(&options.Collation{
			Locale:   "en",
			Strength: 2,
		})

	var user *models.User
	err = s.db.Collection("users").FindOne(context, bson.D{{"email", userEmail}}, findOpts).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	tokenPair, err := utils.GenerateTokenPair(user.UserID, userEmail, user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	// create user session
	userSession := models.UserSession{
		ID:           primitive.NewObjectID(),
		UserID:       user.UserID,
		Token:        tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		IsActive:     true,
	}
	_, err = s.userSessionService.CreateUserSession(context, userSession)
	if err != nil {
		utils.Log.Error("Failed to create user session", zap.Error(err))
	}

	return utils.NewApiResponse([]interface{}{tokenPair}, "User Logged In Successfully", http.StatusOK), nil
}

func (s *AuthService) RefreshToken(c *gin.Context, req models.RefreshTokenRequest) (*utils.ApiResponse, error) {
	utils.Log.Info("RefreshToken", zap.String("refresh_token", req.RefreshToken))

	_, err := s.userSessionService.ValidateUserSessionByRefreshToken(c, req)
	if err != nil {
		utils.Log.Error("RefreshToken", zap.Error(err))
		return nil, err
	}

	token, err := utils.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		utils.Log.Error("RefreshToken", zap.Error(err))
		return nil, err
	}

	return utils.NewApiResponse([]interface{}{token}, "Refreshed Access Token Successfully", http.StatusOK), nil
}
