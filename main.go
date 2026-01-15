package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SnackLog/auth-lib"
	"github.com/SnackLog/database-api-wrapper/internal/config"
	"github.com/SnackLog/database-api-wrapper/internal/database"
	"github.com/SnackLog/database-api-wrapper/internal/handlers/product"
	serviceConfigLib "github.com/SnackLog/service-config-lib"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	loadConfigs()
	db := connectDatabase()
	defer disconnectDatabase(db)
	router := setupRouter(db)

	router.Run()
}

func loadConfigs() {
	err := serviceConfigLib.LoadConfig()
	if err != nil {
		panic("error while parsing service config: " + err.Error())
	}

	err = config.LoadConfig()
	if err != nil {
		panic("error while parsing config: " + err.Error())
	}
}

func connectDatabase() *mongo.Client {
	client, err := database.Connect(config.GetConfig().GetDatabaseConnectionString())
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		panic("Failed to ping database: " + err.Error())
	}

	return client
}

func disconnectDatabase(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic("Failed to disconnect from database: " + err.Error())
	}
}

func setupRouter(db *mongo.Client) *gin.Engine {
	r := gin.Default()
	products := r.Group("/products")

	setupEndpoints(products, db)

	return r
}

func setupEndpoints(router *gin.RouterGroup, db *mongo.Client) {
	productController := &product.ProductController{DB: db}

	if serviceConfigLib.GetConfig().DebugBypassAuthMiddleware {
		log.Println("\033[1;31mWARNING: DEBUG_BYPASS_AUTH_MIDDLEWARE is enabled. Setting up endpoints without authentication!\033[0m")
		router.GET("/product", productController.Get)
		router.GET("/product/:id", productController.GetID)
		return
	}
	router.GET("/product", authlib.Authentication, productController.Get)
	router.GET("/product/:id", authlib.Authentication, productController.GetID)
}
