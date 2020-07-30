package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/insert_register_user", insertRegisterUser).Methods("POST")
	router.HandleFunc("/login_user", loginUser).Methods("POST")

	fmt.Println("Server has successfully run on port 80")
	http.ListenAndServe(":80", router)

}
