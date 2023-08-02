package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func main() {
	err := initServer()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/add", addUserHandler)

	log.Println("Server started on http://localhost:3322")
	log.Fatal(http.ListenAndServe(":3322", nil))
}

func initServer() error {
	users = []User{}
	return nil
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
}
