package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"test/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	app.StartApp()
}
