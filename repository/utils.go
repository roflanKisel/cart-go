package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func idFilter(id string) (bson.M, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return bson.M{}, err
	}

	return bson.M{"_id": objectID}, nil
}
