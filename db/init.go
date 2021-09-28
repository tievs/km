package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Client *mongo.Client
var ctx = context.TODO()
var Doctor *mongo.Collection
var Patient *mongo.Collection
var Drug *mongo.Collection
var Nurse *mongo.Collection
func Init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("DB_URL")
	clientOptions := options.Client().ApplyURI(url)
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	Doctor = Client.Database("km").Collection("doctor")
	Patient = Client.Database("km").Collection("patient")
	Drug = Client.Database("km").Collection("drug")
	Nurse = Client.Database("km").Collection("nurse")
}

