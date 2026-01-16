package product

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetProductByID(db *mongo.Client, id string) (string, error) {
	bsonResult, err := db.Database("off").Collection("products").FindOne(context.TODO(), bson.M{"_id": id}).Raw()

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

func SearchProductByName(db *mongo.Client, query string, limit int) (string, error) {
	filter := bson.M{
		"$text": bson.M{
			"$search": strings.TrimSpace(query),
		},
	}

	filterOptions := options.Find()
	filterOptions.SetProjection(bson.M{
		"score": bson.M{
			"$meta": "textScore",
		},
	})
	filterOptions.SetSort(bson.M{
		"score": bson.M{
			"$meta": "textScore",
		},
	})
	filterOptions.SetLimit(int64(limit))

	cursor, err := db.Database("off").Collection("products").Find(context.TODO(), filter, filterOptions)
	if err != nil {
		return "", fmt.Errorf("failed to search products: %v", err)
	}
	defer cursor.Close(context.TODO())

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return "", fmt.Errorf("failed to decode products: %v", err)
	}

	resultsBson := bson.M{
		"results": results,
	}

	jsonBytes, err := bson.MarshalExtJSON(resultsBson, true, false)
	if err != nil {
		return "", fmt.Errorf("failed to marshal products to JSON: %v", err)
	}

	return string(jsonBytes), nil
}
