package product

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetProductByID(db *mongo.Client, id string) (string, error) {
	bsonResult, err := db.Database("off").Collection("products").FindOne(context.TODO(), map[string]string{"_id": id}).Raw()

	if err == mongo.ErrNoDocuments {
		return "", nil
	}

	if err != nil {
		return "", fmt.Errorf("failed to find product: %v", err)
	}

	jsonBytes, err := bson.MarshalExtJSON(bsonResult, true, false)

	if err != nil {
		return "", fmt.Errorf("failed to marshal product to JSON: %v", err)
	}

	return string(jsonBytes), nil
}
