package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func user_service() string {
	// fmt.Println("db service")
	return "Here"
}

func hashing_password(password string) []byte {
	//Declare username variable
	//username := "Yhishuang"
	//userPassword1 := "some user-provided password"

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	//fmt.Println("Hash to store:", string(hash))
	// Store this "hash" somewhere, e.g. in your database
	return hash

}

func password_validation(hashPassword1 []byte, hashPassword2 []byte) string {
	// After a while, the user wants to log in and you need to check the password he entered
	// userPassword2 := "some user-provided password"
	// hashFromDatabase := []byte("$2a$10$7Yu83J03Lt8RGBFdnT5rKu3T1K8UD3c/Pzp/Ijt1haPOsJYNNN.AS")

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword(hashPassword1, hashPassword2); err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}

	//fmt.Println("Password was correct!")
	return "Password was correct!"

}

func insert_user(username string, password string) {

	//mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://leviooi:1123956321@206.189.152.72:27017/?authSource=admin"))
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
		{Key: "password", Value: string(password)},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v documents into episode collection!\n", usersResult)
}