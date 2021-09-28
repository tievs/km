package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDocument(ctx context.Context, collection *mongo.Collection, filter interface{}, result interface{}) error {
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func CreateDocument(ctx context.Context,collection *mongo.Collection, document interface{}) (interface{} ,error) {
	insertResult, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func UpdateDocument(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) (interface{} ,error) {
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult.UpsertedID, nil
}

func DeleteDocument(ctx context.Context,collection *mongo.Collection, filter interface{}) (interface{} ,error) {
	deleteCollection, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return deleteCollection.DeletedCount, err
}

func GetDocumentList(ctx context.Context,collection *mongo.Collection, filter interface{}, results interface{}, options *options.FindOptions) error {
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		return err
	}
	if err = cursor.All(ctx, results); err != nil {
		return err
	}
	return nil
}

