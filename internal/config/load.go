package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	MongoHost     string
	MongoPort     int64
	MongoUser     string
	MongoPass     string
	MongoDatabase string
}

var AppConfig Config

func LoadConfig() error {
	config := Config{
		MongoHost:     os.Getenv("MONGO_HOST"),
		MongoPort:     0, // Will be set after parsing
		MongoUser:     os.Getenv("MONGO_USER"),
		MongoPass:     os.Getenv("MONGO_PASS"),
		MongoDatabase: os.Getenv("MONGO_DATABASE"),
	}

	err := parsePort(&config)
	if err != nil {
		return fmt.Errorf("error parsing MONGO_PORT: %v", err)
	}

	if err := validateConfig(config); err != nil {
		return fmt.Errorf("config validation failed: %v", err)
	}

	AppConfig = config
	return nil
}

func parsePort(config *Config) error {
	port, err := strconv.ParseInt(os.Getenv("MONGO_PORT"), 10, 0)
	if err != nil {
		return fmt.Errorf("invalid MONGO_PORT: %v", err)
	}
	config.MongoPort = port
	return nil
}

func validateConfig(config Config) error {
	if config.MongoHost == "" {
		return fmt.Errorf("MONGO_HOST is required")
	}
	if config.MongoPort <= 0 {
		return fmt.Errorf("MONGO_PORT is required and must be a positive integer")
	}
	if config.MongoUser == "" {
		return fmt.Errorf("MONGO_USER is required")
	}
	if config.MongoPass == "" {
		return fmt.Errorf("MONGO_PASS is required")
	}
	if config.MongoDatabase == "" {
		return fmt.Errorf("MONGO_DATABASE is required")
	}

	return nil
}

func GetConfig() Config {
	return AppConfig
}

func SetConfig(config Config) {
	AppConfig = config
}

func (c Config) GetDatabaseConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/", c.MongoUser, c.MongoPass, c.MongoHost, c.MongoPort)
}
