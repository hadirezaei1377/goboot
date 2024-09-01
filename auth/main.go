package auth

import (
	"log"
	"net/http"

	"auth/config"
	"auth/routes"
)

func main() {
	config.LoadConfig()
	router := routes.SetupRouter()

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
