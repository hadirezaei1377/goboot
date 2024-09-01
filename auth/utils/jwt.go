package utils

// this file used for generate, proccess and update jwt tokens

var jwtKey = []byte("my_secret_key")

func GenerateJWT(userID uint) (string, error) {}

func VerifyJWT(tokenStr string) (uint, error) {}
