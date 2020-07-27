package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Check is it the user is already registered
func checkRegisterUser(username string) int {
	//connectionPath := mongodbConnectionPath()
	//mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbConnectionPath()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("lfu_db")
	usersCollection := quickstartDatabase.Collection("users")

	filterCursor, err := usersCollection.Find(ctx, bson.M{"user": username})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(len(episodesFiltered))
	return len(episodesFiltered)
}

//Insert user info into db
func insertMongoUser(username string, password string) string {

	result := ""

	if checkRegisterUser(username) == 0 {

		//mongodb client
		client, err := mongo.NewClient(options.Client().ApplyURI(mongodbConnectionPath()))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		quickstartDatabase := client.Database("lfu_db")
		usersCollection := quickstartDatabase.Collection("users")

		usersResult, err := usersCollection.InsertOne(ctx, bson.D{
			{Key: "user", Value: username},
			{Key: "password", Value: hashingPassword(password)},
		})

		if err != nil {
			log.Fatal(err)
		}

		result = fmt.Sprintf("Successful: Inserted a single document: %d.", usersResult.InsertedID)

	} else {
		result = "Failed: username already registered."
	}
	return result
}

func checkLoginUser(username string, password string) {

	//connectionPath := mongodbConnectionPath()
	//mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbConnectionPath()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("lfu_db")
	usersCollection := quickstartDatabase.Collection("users")

	filterCursor, err := usersCollection.Find(ctx, bson.M{"user": username})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}

	vPassword := episodesFiltered
	fmt.Println(vPassword)

}

//Login
func loginUser(username string, password string) string {
	// 1. check the username is it exists in db
	// 2. get the info
	// 3. validate the password
	return ""
}
