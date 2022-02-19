package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)
const mongourl = "mongodb://username:password@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019/admin?replicaSet=mongos&authSource=admin"

func Getclient() *mongo.Client {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongourl))
	if err != nil {
		log.Println(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Println(err)
	}

return client
}

func Insert(papers Papers) {
	db := Getclient()
	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			log.Println(err)
		}
	}()
	c := db.Database("bingpapers").Collection("papers")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	_, err := c.InsertOne(ctx, &papers)

	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully insert data in mongodb.")
}