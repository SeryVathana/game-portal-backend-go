package core

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var database *mongo.Database

func ConnectToDB() {
	fmt.Println("Connecting to the database...")
	config := LoadConfig()

	// Create backgroud context with 10 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB Database with input connection string (dev: local database)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBConnectionString))
	if err != nil {
		fmt.Printf("Unable to connect to database %s\n", config.DBConnectionString)
		panic(err)
	}

	mongoClient = client

	// Test connection to make sure it's connected
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Unable to ping to database")
		panic(err)
	}

	database = mongoClient.Database(config.DBName)

	userCol := database.Collection(TBL_USER)

	//Create user db indexes
	uniqueEmailIndex := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	userCol.Indexes().CreateOne(context.Background(), uniqueEmailIndex)

	fmt.Println("Connected to the database!")
}

func DisconnectFromDB() {
	fmt.Println("Cleaning up connections...")
	if mongoClient != nil {
		err := mongoClient.Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	}
}

func GetDatabaseConnection() *mongo.Database {
	return database
}
