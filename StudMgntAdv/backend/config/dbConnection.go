package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"StudentManagementAdvance/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(filePath string) (*model.DatabaseConfiguration, error) {
	byteData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error while reading file: %v", err)
	}

	var dbConfig model.DatabaseConfiguration
	if err := json.Unmarshal(byteData, &dbConfig); err != nil {
		return nil, fmt.Errorf("Error while unmarshalling JSON: %v", err)
	}

	err = ConnectMongo(dbConfig.MongoConnection)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to MongoDB: %v", err)
	}

	return &dbConfig, nil
}

var DB *mongo.Database

func ConnectMongo(dbConfig model.DatabaseConfig) error {
    // MongoDB URI format for Atlas
    uri := fmt.Sprintf(
        "mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
        dbConfig.Username,
        dbConfig.Password,
        dbConfig.ServerIP,  
        dbConfig.Database,
    )

    clientOptions := options.Client().ApplyURI(uri)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("Error connecting to MongoDB:", err)
        return err
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Could not ping MongoDB:", err)
        return err
    }

    fmt.Println("Connected to MongoDB!")

    DB = client.Database(dbConfig.Database)


    return nil
}
