package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id"`
	UserID          string             `bson:"user_id" json:"user_id"`
	FirstName       string             `bson:"first_name" json:"first_name"`
	LastName        string             `bson:"last_name" json:"last_name"`
	Email           string             `bson:"email" json:"email"`
	Role            string             `bson:"role" json:"role"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
	FavouriteGenres []Genre            `bson:"favourite_genres" json:"favourite_genres"`
}

type UserDTO struct {
	UserID          string  `bson:"user_id" json:"user_id" validate:"required,min=2"`
	FirstName       string  `bson:"first_name" json:"first_name" validate:"required,min=2"`
	LastName        string  `bson:"last_name" json:"last_name" validate:"required,min=2"`
	Email           string  `bson:"email" json:"email" validate:"required,email"`
	Role            string  `bson:"role" json:"role" validate:"required"`
	FavouriteGenres []Genre `bson:"favourite_genres" json:"favourite_genres" validate:"required,dive"`
}
