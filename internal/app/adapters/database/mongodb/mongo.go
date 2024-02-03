package mongodb

import (
	"back-vozes-cifras/internal/app/application/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	client *mongo.Client
}

func NewMongoRepository(client *mongo.Client) repository.Repository {
	return &mongodb{client: client}
}
