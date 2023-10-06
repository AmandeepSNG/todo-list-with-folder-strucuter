package db

import (
	"context"
	"sync"
	constants "todolist-app/common/shared-constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDbClient() (*mongo.Client, error) {
	var mongoClientInstance *mongo.Client
	var mongoOnce sync.Once
	var clientInstanceError error
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(string(constants.MONGO_DB_URL))

		client, err := mongo.Connect(context.TODO(), clientOptions)
		mongoClientInstance = client
		clientInstanceError = err
	})
	return mongoClientInstance, clientInstanceError
}
