package product

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetProductByID(db *mongo.Client, id string) (*Product, error) {
	var product *Product = new(Product)
	err := db.Database("off").Collection("products").FindOne(context.TODO(), bson.M{"_id": id}).Decode(product)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find product: %v", err)
	}

	return product, nil
}

func SearchProductByName(db *mongo.Client, query string, limit int) (*[]Product, error) {
	filter := bson.M{
		"$text": bson.M{
			"$search": strings.TrimSpace(query),
		},
		"states_tags": "en:nutrition-facts-completed",
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
		return nil, fmt.Errorf("failed to search products: %v", err)
	}
	defer cursor.Close(context.TODO())

	var results []Product
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, fmt.Errorf("failed to decode products: %v", err)
	}

	return &results, nil
}
