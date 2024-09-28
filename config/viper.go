package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")

	fmt.Printf("Connecting to database at %s:%d\n", dbHost, dbPort)
}
