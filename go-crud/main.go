package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app, err := InitializeApp()
	if err != nil { log.Fatal(err) }
	app.Start(":8080")
}
