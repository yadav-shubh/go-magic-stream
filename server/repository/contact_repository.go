package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactRepository struct {
	db *mongo.Database
}

func NewContactRepository(db *mongo.Database) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) Create(ctx *gin.Context, message models.ContactMessage) error {
	_, err := r.db.Collection("contact-messages").InsertOne(ctx, message)
	if err != nil {
		return fmt.Errorf("Error while creating the contact message")
	}

	return nil
}
