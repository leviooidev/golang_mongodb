package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Result   string `json:"result"`
	//password string `json:"password"`
}

func insertRegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		// in case of any error
		return
	}

	// Use the r.Form.Get() method to retrieve the relevant data fields
	// from the r.Form map.
	vusername := r.Form.Get("username")
	vpassword := r.Form.Get("password")

	result := insertMongoUser(vusername, vpassword)

	json.NewEncoder(w).Encode(result)
}

func loginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		// in case of any error
		return
	}

	// Use the r.Form.Get() method to retrieve the relevant data fields
	// from the r.Form map.
	vusername := r.Form.Get("username")
	vpassword := r.Form.Get("password")

	result, rusername := checkLoginUser(vusername, vpassword)

	res := &User{
		Username: rusername,
		Result:   result,
	}

	json.NewEncoder(w).Encode(res)

}
