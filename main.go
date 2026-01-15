package main

import (
	"context"
	"fmt"

	"github.com/SnackLog/database-api-wrapper/internal/config"
	"github.com/SnackLog/database-api-wrapper/internal/database"
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
	fmt.Println("Config loaded successfully:", serviceConfigLib.GetConfig())

	err = config.LoadConfig()
	if err != nil {
		panic("error while parsing config: " + err.Error())
	}
	fmt.Println("Database config loaded successfully:", config.GetConfig().GetDatabaseConnectionString())
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

func setupRouter(*mongo.Client) *gin.Engine {
	r := gin.Default()
	return r
}
