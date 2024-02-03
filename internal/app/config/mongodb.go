package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() *mongo.Client {
	mongoURI := viper.GetString("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		panic(err)
	}

	// Ping the MongoDB server to ensure connectivity
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Printf("Could not ping to MongoDB: %v\n", err)
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
