package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func mongodbConnectionPath() string {

	result := "mongodb://leviooi:1123956321@206.189.152.72:27017/?authSource=admin"

	return result
}

func hashingPassword(password string) []byte {
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

func passwordValidation(hashPassword1 []byte, hashPassword2 []byte) string {
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
