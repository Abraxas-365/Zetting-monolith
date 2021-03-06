package project_repository

import (
	"context"
	models "zetting/pkg/project/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mongoRepository) GetProjectByProjectId(projectId interface{}) (*models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)
	projectObjecId, err := primitive.ObjectIDFromHex(projectId.(string))
	if err != nil {
		return nil, err
	}
	var project models.Project
	filter := bson.M{"_id": projectObjecId}
	if err := collection.FindOne(ctx, filter).Decode(&project); err != nil {
		return nil, err
	}
	return &project, nil
}
