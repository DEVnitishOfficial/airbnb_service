package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
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
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

// from env we always get string but suppose we want to get int then we can use below function
func GetInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	// convert string to int using strconv.Atoi(ASCII to integer), it's a built-in function in go
	intValue, err := strconv.Atoi(value)

	if err != nil {
		fmt.Printf("Error converting %s to int: %v\n", value, err)
		return fallback
	}
	return intValue

}

func GetBool(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Printf("Error converting %s to bool: %v\n", value, err)
		return fallback
	}
	return boolValue
}
