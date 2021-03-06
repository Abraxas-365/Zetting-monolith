package workRequest_repository

import (
	"context"
	models "zetting/pkg/workRequest/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r mongoRepository) GetWorkRequest(workRequestId interface{}) (*models.WorkRequest, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	projectObjectId, err := primitive.ObjectIDFromHex(workRequestId.(string))
	if err != nil {
		return nil, err
	}

	var workRequest models.WorkRequest
	filter := bson.M{"_id": projectObjectId}
	if err := collection.FindOne(ctx, filter).Decode(&workRequest); err != nil {
		return nil, err
	}
	return &workRequest, nil

}
