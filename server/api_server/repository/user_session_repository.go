package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/models"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSessionRepository struct {
	db *mongo.Database
}

func (r *UserSessionRepository) FindAllByUserId(ctx context.Context, id string) (*[]models.UserSession, error) {
	var sessions *[]models.UserSession
	cursor, err := r.db.Collection("user_sessions").Find(ctx, bson.M{"user_id": id, "is_active": true})
	if err != nil {
		return nil, fmt.Errorf("failed to find user sessions")
	}
	err = cursor.All(ctx, sessions)
	if err != nil {
		return nil, fmt.Errorf("failed to find user sessions")
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	return sessions, nil
}

func (r *UserSessionRepository) UpdateAll(ctx context.Context, id string) {
	_, err := r.db.Collection("user_sessions").UpdateMany(ctx,
		bson.M{"user_id": id},
		bson.M{"$set": bson.M{"is_active": false, "updated_at": time.Now()}},
	)
	if err != nil {
		utils.Log.Info("failed to update all user sessions")
	} else {
		utils.Log.Info("all user sessions updated successfully")
	}
}

func (r *UserSessionRepository) Create(ctx context.Context, session models.UserSession) error {
	_, err := r.db.Collection("user_sessions").InsertOne(ctx, session)
	if err != nil {
		return fmt.Errorf("failed to create user session")
	}
	return nil
}

func (r *UserSessionRepository) FindByRefreshToken(c *gin.Context, refreshToken string) (*models.UserSession, error) {
	var userSession *models.UserSession
	err := r.db.Collection("user_sessions").FindOne(c, bson.M{"refresh_token": refreshToken, "is_active": true}).Decode(&userSession)
	if err != nil {
		return nil, fmt.Errorf("invalid or expired refresh token")
	}
	return userSession, nil
}

func NewUserSessionRepository(db *mongo.Database) *UserSessionRepository {
	return &UserSessionRepository{db: db}
}
