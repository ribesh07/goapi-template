package main

import (
	"log"
	"net/http"

	"goapi/internal/app"
)

func main() {

	router := app.New()

	log.Println("Server running on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
