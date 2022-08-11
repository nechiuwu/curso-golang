package main

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieDb struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Ano       string             `bson:"ano"`
	Director  string             `bson:"director"`
	Completed bool               `bson:"completed"`
}

func main() {
	// URLs amigables usando /
	router := NewRouter()

	// Levanta conexion a DB
	connectMongoDb()

	// crea un server y recibe router
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

	movie := &MovieDb{
		ID:       primitive.NewObjectID(),
		Name:     "totoro",
		Ano:      "2000",
		Director: "gibli",
	}
	createTask(movie)
}

func createTask(movie *MovieDb) error {
	/*d, err := collection.InsertOne(ctx, movie)
	log.Println(d)
	return err*/
	return nil
}
