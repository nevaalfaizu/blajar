package app

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"tes1/varglobal"

	"github.com/joho/godotenv"
)

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logFile, err := os.Create("app.log")
	if err != nil {
		log.Fatal("failed to create log file")
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("App started")

	// load environment variable form file and set to file

	// retrieve and print the environment variable
	// Akses variabel environment
	varglobal.DatabaseName = os.Getenv("DATABASE_NAME")
	fmt.Println("DATABASE_NAME:", varglobal.DatabaseName)
	varglobal.DatabaseUser = os.Getenv("DATABASE_USER")
	fmt.Println("DATABASE_USER:", varglobal.DatabaseUser)
	varglobal.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	fmt.Println("DATABASE_PASSWORD:", varglobal.DatabasePassword)
	varglobal.DatabaseHost = os.Getenv("DATABASE_HOST")
	fmt.Println("DATABASE_HOST:", varglobal.DatabaseHost)
	varglobal.DatabasePort = os.Getenv("DATABASE_PORT")
	fmt.Println("DATABASE_PORT:", varglobal.DatabasePort)
	// Default port, can be overridden by environment variable
	port := os.Getenv("MAIN_PORT")
	if port != "" {
		fmt.Println("MAIN_PORT:", port)
		varglobal.MainPort, _ = strconv.Atoi(port)
		// Optionally, you can parse and assign it to varglobal.MainPort if needed
	}

	debug := os.Getenv("APP_DEBUG")
	fmt.Println("DEBUG:", debug)

}
