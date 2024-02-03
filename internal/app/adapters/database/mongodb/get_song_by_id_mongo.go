package mongodb

import (
	"back-vozes-cifras/internal/app/application/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongodb) GetSongByIDRepository(id int) (*model.Song, error) {
	collection := m.client.Database("vozes-chord").Collection("song")
	filter := bson.D{{"_id", id}}
	result := collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var song song
	err := result.Decode(&song)
	if err != nil {
		return nil, err
	}

	return song.toModel(), nil
}
