package main

import (
	"log"
	"net/http"

	"goapi/internal/app"
	"goapi/internal/config"
	"goapi/internal/database"
)

func main() {

	config.LoadEnv()
	cfg := config.New()
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	router := app.New(db)

	log.Println("Server running on :8080")
	log.Println(cfg.DatabaseURL)

	log.Fatal(http.ListenAndServe(":8080", router))
}
