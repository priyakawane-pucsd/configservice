package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Host     string
	Database string
	Username string
	Password string
	Port     int
}

func getMongoConnection(ctx context.Context, conf *Config) *mongo.Client {
	// Build the MongoDB connection string
	uri := "mongodb://%s:%s@%s:%d/%s?retryWrites=true&loadBalanced=false&serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=admin"

	connectionString := fmt.Sprintf(uri,
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Create a new MongoDB client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	return client
}

type Repository struct {
	client *mongo.Client
}

func NewRepository(ctx context.Context, conf *Config) *Repository {
	client := getMongoConnection(ctx, conf)
	return &Repository{client: client}
}

func (repo *Repository) Ping(ctx context.Context) error {
	return repo.client.Ping(ctx, readpref.Primary())
}
