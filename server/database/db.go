package database

import (
	"context"
	"sync"
	"time"

	"github.com/yadav-shubh/go-magic-stream/config"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	client     *mongo.Client
	dbInstance *mongo.Database
	once       sync.Once
)

func ConnectMongo() *mongo.Database {
	once.Do(func() {
		mongoCfg := config.Get().DB

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		opts := options.Client().
			ApplyURI(mongoCfg.Uri).
			SetMaxPoolSize(100).
			SetMinPoolSize(5).
			SetMaxConnIdleTime(30 * time.Second)

		c, err := mongo.Connect(ctx, opts)
		if err != nil {
			utils.Log.Error("MongoDB connection failed", zap.Error(err))
			panic("MongoDB connection failed")
		}

		if err = c.Ping(ctx, nil); err != nil {
			utils.Log.Error("MongoDB ping failed", zap.Error(err))
			panic("MongoDB ping failed")
		}

		utils.Log.Info("MongoDB connected successfully")
		client = c
		dbInstance = c.Database(mongoCfg.Database)
	})

	return dbInstance
}

func GetMongoClient() *mongo.Client {
	return client
}

func GetCollectionClient(collectionName string) *mongo.Collection {
	return dbInstance.Collection(collectionName)
}
