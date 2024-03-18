package main

import (
	"log"
	"net/http"
	"terraform-go-api/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
