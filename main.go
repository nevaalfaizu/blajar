package main

import (
	"tes1/app"
	"tes1/dbku"
)

func main() {

	app.LoadConfig() // Load configuration from .env file
	dbku.InitDB()
	app.StartApi()

}
