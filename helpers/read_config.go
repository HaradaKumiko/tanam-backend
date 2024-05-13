package helpers

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting current working directory: %s\n", err)
	}
	log.Printf("Current working directory: %s\n", dir)

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading configuration file: %s\n", err)
	}

	return viper.GetString(key)
}
