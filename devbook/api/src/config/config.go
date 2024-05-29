package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LOAD THE VARIABLES AS PUBLIC VALUES FROM LoadEnvs function
var (
	Port = 0
	// the database connection string
	ConnectionStr = ""
)

// Load the envs with a env package
func LoadEnvs() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal("Error loading .env file")
	}
	// Converting the port to 9000
	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 9000
	}

	ConnectionStr = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

}
