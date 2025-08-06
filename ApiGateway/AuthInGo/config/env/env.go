package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func load() {
	err := godotenv.Load()
	if err != nil {
		// log the error if .env file is not found
		// or if there is an error loading it
		fmt.Println("Error loading .env file")
	}
}

// here GetString is used so that whenver we need a environment variable
// we just call the GetString function with key and we get the value
// if the key is not found, it returns the fallback value
func GetString(key string, fallback string) string {
	load() // ensure .env is loaded before accessing any keys
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
