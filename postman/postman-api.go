package postman

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var inventory []Item

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range inventory {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = "ID" + string(len(inventory)+1)
	inventory = append(inventory, item)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range inventory {
		if item.ID == params["id"] {
			inventory = append(inventory[:i], inventory[i+1:]...)
			var updatedItem Item
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.ID = params["id"]
			inventory = append(inventory, updatedItem)
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range inventory {
		if item.ID == params["id"] {
			inventory = append(inventory[:i], inventory[i+1:]...)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/item/{id}", getItem).Methods("GET")
	router.HandleFunc("/item", createItem).Methods("POST")
	router.HandleFunc("/item/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
