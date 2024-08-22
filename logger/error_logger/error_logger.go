package errorlogger

import (
	"encoding/json"
	"os"
)

type ErrorLog struct {
	Message string `json:"message"`
}

func LogError(message string) {
	log := ErrorLog{Message: message}

	file, err := os.OpenFile("error_logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(log); err != nil {
		panic(err)
	}
}
