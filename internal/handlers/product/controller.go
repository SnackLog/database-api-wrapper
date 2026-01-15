package product

import "go.mongodb.org/mongo-driver/v2/mongo"

type ProductController struct {
	DB *mongo.Client
}
