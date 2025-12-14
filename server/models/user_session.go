package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSession struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	UserID       string             `bson:"user_id" json:"user_id"`
	Token        string             `bson:"token" json:"token"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	IsActive     bool               `bson:"is_active" json:"is_active"`
}

type UserSessionDTO struct {
	UserID       string `bson:"user_id" json:"user_id"`
	Token        string `bson:"token" json:"token"`
	RefreshToken string `bson:"refresh_token" json:"refresh_token"`
	IsActive     bool   `bson:"is_active" json:"is_active"`
}
