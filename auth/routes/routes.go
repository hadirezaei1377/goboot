package routes

import (
	"auth/controllers"
	"auth/middlewares"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	authRoutes := router.PathPrefix("/auth").Subrouter()
	authRoutes.Use(middlewares.AuthMiddleware)
	authRoutes.HandleFunc("/profile", controllers.GetProfile).Methods("GET")
	authRoutes.HandleFunc("/logout", controllers.Logout).Methods("POST")

	return router
}
