package controllers

import (
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	//register user
}

func Login(w http.ResponseWriter, r *http.Request) {
	// log in by user and generate jwt token
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//log out by user and delete session
}
