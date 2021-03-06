package user_repository

import (
	"context"
	models "zetting/pkg/user/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) AddProjectToUser(projectData models.AddProjectToUser, document string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	// ownerObjectId, err := primitive.ObjectIDFromHex(projectData.OwnerId.(string))
	// if err != nil {
	// 	return err
	// }

	projectObjectId, err := primitive.ObjectIDFromHex(projectData.ProjectId.(string))
	if err != nil {
		return err
	}

	userObjectId, err := primitive.ObjectIDFromHex(projectData.UserId.(string))
	if err != nil {
		return err
	}

	filter := bson.M{"_id": userObjectId}
	update := bson.M{
		"$push": bson.M{
			document: projectObjectId,
		},
	}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil

}
