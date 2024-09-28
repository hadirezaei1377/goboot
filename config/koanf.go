package config

import (
	"fmt"
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")

func main() {

	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	dbHost := k.String("database.host")
	dbPort := k.Int("database.port")

	fmt.Printf("Connecting to database at %s:%d\n", dbHost, dbPort)
}
