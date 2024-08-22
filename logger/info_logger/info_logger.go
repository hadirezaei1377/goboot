package infologger

import (
	"os"

	"gopkg.in/yaml.v2"
)

type InfoLog struct {
	Message string `yaml:"message"`
}

func LogInfo(message string) {
	log := InfoLog{Message: message}

	file, err := os.OpenFile("info_logs.yaml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(log); err != nil {
		panic(err)
	}
}
