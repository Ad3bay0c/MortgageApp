package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Client *mongo.Client

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Minute)

	Client, _ = mongo.NewClient(options.Client().ApplyURI(URI))

	err := Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatalf("Database not connected: %v", err.Error())
	} else {
		log.Printf("Database Connected....")
	}
}
