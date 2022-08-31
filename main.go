package main

import (
	"encoding/json"
	"fmt"
	"golangTraining/service"
	"net/http"
)

var PORT = ":8080"

var userSvc = service.NewUserService()

func main() {

	http.HandleFunc("/", greet)
	http.HandleFunc("/user", getUser)
	http.HandleFunc("/register", registerUser)

	http.ListenAndServe(PORT, nil)

}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello Word"
	fmt.Fprint(w, msg)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		resHasil := userSvc.GetUser()
		json.NewEncoder(w).Encode(resHasil)
		return
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("name")

		res := userSvc.Register(&service.User{Nama: name})
		fmt.Fprint(w, res)
		return
	}
}
