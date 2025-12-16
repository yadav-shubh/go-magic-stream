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
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MovieService struct {
	db *mongo.Database
}

func NewMovieService(db *mongo.Database) *MovieService {
	return &MovieService{db: db}
}

func (s *MovieService) GetMovies(
	ctx context.Context,
	page int,
	size int,
	search string,
	genre string,
) ([]*models.MovieDTO, error) {

	collection := s.db.Collection("movies")

	filter := bson.M{}
	if search != "" {
		filter["title"] = bson.M{
			"$regex":   search,
			"$options": "i",
		}
	}

	if genre != "" {
		filter["genre.genre_name"] = bson.M{
			"$regex":   genre,
			"$options": "i",
		}
	}

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}

	skip := int64((page - 1) * size)
	limit := int64(size)

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("error fetching movies: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			utils.Log.Error("error closing cursor", zap.Error(err))
		}
	}(cursor, ctx)

	var movies []*models.MovieDTO
	if err := cursor.All(ctx, &movies); err != nil {
		return nil, fmt.Errorf("error decoding movies: %w", err)
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
