package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Conexão com o banco

var MONGODB_URL = "MONGODB_URI"
var ctx = context.TODO()
var Collection *mongo.Collection

func InitConnection() {
	mongodb_uri := os.Getenv(MONGODB_URL)
	clientOptions := options.Client().ApplyURI(mongodb_uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected Mongo")

	// // Teste de insert
	// coll := client.Database("Lucas").Collection("Test")
	// newUser := models.User{
	// 	Nome:  "Teste2",
	// 	Email: "Teste@teste.com",
	// }
	// result, err := coll.InsertOne(context.TODO(), newUser)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}
