package users

import (
	"context"
	"fmt"
	"time"

	database "github.com/mstevenacero/application-core-blues/database"
	"github.com/mstevenacero/application-core-blues/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsersMongo() ([]models.User, error) {
	collection := database.GetMongoClient().Database("blues_core").Collection("users")

	var users []models.User
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("error", err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			fmt.Println("Error decoding product:", err)
		}
		users = append(users, user)
	}
	fmt.Println(users)
	return users, nil
}

func AddUsersMongo(data models.UserJson) (*mongo.InsertOneResult, error) {
	collection := database.GetMongoClient().Database("blues_core").Collection("users")

	user := bson.M{
		"name":     data.Name,
		"email":    data.Email,
		"password": data.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return nil, err
	}
	fmt.Printf("User added successfully with ID: %v\n", result.InsertedID)

	return result, nil
}
