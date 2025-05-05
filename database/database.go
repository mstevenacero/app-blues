package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectToMongo() {
	// URL de conexión de MongoDB
	uri := "mongodb://localhost:27017" // Cambia esto si usas MongoDB Atlas o un servidor remoto

	// Conexión a MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Verifica la conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a MongoDB!")

}

func GetMongoClient() *mongo.Client {
	return client
}
