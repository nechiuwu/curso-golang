package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectMongoDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	log.Println("Exito ping")

	movie := &MovieDb{
		ID:       primitive.NewObjectID(),
		Name:     "totoro",
		Ano:      "2000",
		Director: "gibli",
	}

	createMovie(movie, client, ctx)
}

func createMovie(movie *MovieDb, client *mongo.Client, ctx context.Context) {
	collection := client.Database("movie").Collection("movies")
	res, err := collection.InsertOne(ctx, bson.D{{"name", movie.Name}, {"ano", movie.Ano}, {"director", movie.Director}})
	if err == nil {
		panic(err)
	}
	id := res.InsertedID
	log.Printf("Registrado con id: %s", id)
}
