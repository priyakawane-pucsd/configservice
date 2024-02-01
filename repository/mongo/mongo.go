package mongo

import (
	"configservice/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Host       string
	Database   string
	Collection string
	Username   string
	Password   string
	Port       int
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
	fmt.Println("mongo connection established successful.....")
	return client
}

type Repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRepository(ctx context.Context, conf *Config) *Repository {
	client := getMongoConnection(ctx, conf)
	return &Repository{client: client,
		collection: client.Database(conf.Database).Collection(conf.Collection),
	}
}

func (repo *Repository) Ping(ctx context.Context) error {
	return repo.client.Ping(ctx, readpref.Primary())
}

func (repo *Repository) SavePingPong(ping *models.PingPong) error {
	response, err := repo.collection.InsertOne(context.TODO(), ping)
	if err != nil {
		return err
	}
	fmt.Println("response of repo$$$", response)
	return err
}

func (repo *Repository) GetAllPings(ctx context.Context) ([]models.PingPong, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var pings []models.PingPong
	err = cursor.All(ctx, &pings)
	if err != nil {
		return nil, err
	}
	return pings, nil
}

// GetPingByID retrieves a ping record by its ID.
func (repo *Repository) GetPingByID(ctx context.Context, id primitive.ObjectID) (*models.PingPong, error) {
	filter := bson.M{"_id": id}
	var ping models.PingPong
	err := repo.collection.FindOne(ctx, filter).Decode(&ping)
	if err != nil {
		return nil, err
	}
	return &ping, nil
}

// UpdatePingByID updates a ping record by its ID.
func (repo *Repository) UpdatePingByID(ctx context.Context, id primitive.ObjectID, update *models.PingPong) error {
	// Convert the update model to BSON for the update operation
	updateBSON, err := bson.Marshal(update)
	newupdate := bson.M{"$set": bson.Raw(updateBSON)}
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	updateResult, err := repo.collection.UpdateOne(ctx, filter, newupdate)
	if err != nil {
		return err
	}

	if updateResult.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// DeletePingByID deletes a ping record by its ID.
func (repo *Repository) DeletePingByID(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	deleteResult, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
