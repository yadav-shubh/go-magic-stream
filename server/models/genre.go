package models

type Genre struct {
	GenreID   int64  `bson:"genre_id" json:"genre_id"`
	GenreName string `bson:"genre_name" json:"genre_name"`
}

type GenreDTO struct {
	GenreID   int64  `bson:"genre_id" json:"genre_id" validate:"required"`
	GenreName string `bson:"genre_name" json:"genre_name" validate:"required,min=2"`
}
