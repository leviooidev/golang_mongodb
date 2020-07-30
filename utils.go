package main

import (
	"golang.org/x/crypto/bcrypt"
)

func mongodbConnectionPath() string {

	result := "mongodb://leviooi:1123956321@206.189.152.72:27017/?authSource=admin"

	return result
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
