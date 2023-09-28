package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Collection *mongo.Collection

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	Collection = client.Database("auth").Collection("tokens")

	fmt.Println("Connected to MongoDB!")

	/*err = client.Disconnect(context.TODO())

	if err != nil {
		fmt.Println("Connection to MongoDB closed.")
		log.Fatal(err)
	}*/

}
