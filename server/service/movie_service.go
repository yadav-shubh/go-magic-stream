package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/yadav-shubh/go-magic-stream/models"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieService struct {
	db *mongo.Database
}

func NewMovieService(db *mongo.Database) *MovieService {
	return &MovieService{db: db}
}

func (s *MovieService) GetMovies() ([]*models.MovieDTO, error) {
	ctx := context.Background()

	collection := s.db.Collection("movies")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error while getting all moview")
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	var movies []*models.MovieDTO
	for cursor.Next(ctx) {
		var movie *models.MovieDTO
		if err := cursor.Decode(&movie); err != nil {
			continue
		}
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error while getting all moview")
	}

	return movies, nil
}

func (s *MovieService) FindMovieByImdbID(id string) (*models.MovieDTO, error) {
	ctx := context.Background()

	collection := s.db.Collection("movies")
	cursor := collection.FindOne(ctx, bson.M{"imdb_id": id})
	if cursor.Err() != nil {
		return nil, fmt.Errorf("movie not found")
	}

	var movie *models.MovieDTO
	if err := cursor.Decode(&movie); err != nil {
		return nil, fmt.Errorf("error while getting movie by imdb id")
	}

	return movie, nil
}

func (s *MovieService) CreateMovie(m *models.MovieDTO) (*utils.ApiResponse, error) {
	// convert DTO to model
	movie := &models.Movie{
		ID:          primitive.NewObjectID(),
		ImdbID:      m.ImdbID,
		Title:       m.Title,
		PosterPath:  m.PosterPath,
		YoutubeID:   m.YoutubeID,
		Genre:       m.Genre,
		AdminReview: m.AdminReview,
		Ranking:     m.Ranking,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// insert into database
	collection := s.db.Collection("movies")
	_, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		return nil, fmt.Errorf("error while creating movie")
	}

	return utils.NewApiResponseNoData("Movie created", http.StatusCreated), nil
}
