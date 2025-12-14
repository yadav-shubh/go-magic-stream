package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	ImdbID      string             `bson:"imdb_id" json:"imdb_id"`
	Title       string             `bson:"title" json:"title"`
	PosterPath  string             `bson:"poster_path" json:"poster_path"`
	YoutubeID   string             `bson:"youtube_id" json:"youtube_id"`
	Genre       []Genre            `bson:"genre" json:"genre"`
	AdminReview string             `bson:"admin_review" json:"admin_review"`
	Ranking     Ranking            `bson:"ranking" json:"ranking"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type MovieDTO struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	ImdbID      string             `bson:"imdb_id" json:"imdb_id" validate:"required,min=2"`
	Title       string             `bson:"title" json:"title" validate:"required,min=2"`
	PosterPath  string             `bson:"poster_path" json:"poster_path" validate:"required,min=2"`
	YoutubeID   string             `bson:"youtube_id" json:"youtube_id" validate:"required,min=2"`
	Genre       []Genre            `bson:"genre" json:"genre" validate:"required,dive"`
	AdminReview string             `bson:"admin_review" json:"admin_review" validate:"required,min=2"`
	Ranking     Ranking            `bson:"ranking" json:"ranking" validate:"required"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
