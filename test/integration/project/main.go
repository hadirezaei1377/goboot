package project

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Ali"},
	{ID: 2, Name: "bardia"},
}

func main() {
	http.HandleFunc("/users", getUsersHandler)
	http.ListenAndServe(":8080", nil)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
