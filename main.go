package main

import (
	"log"
	"net/http"
	"os"

	"whimsy-api/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/plants", handlers.JSONMiddleware(handlers.AllPlants))
	http.HandleFunc("/animals", handlers.JSONMiddleware(handlers.AllAnimals))
	http.HandleFunc("/colors", handlers.JSONMiddleware(handlers.AllColors))
	http.HandleFunc("/names", handlers.JSONMiddleware(handlers.AllNames))
	http.HandleFunc("/plants/random", handlers.JSONMiddleware(handlers.RandomPlants))
	http.HandleFunc("/animals/random", handlers.JSONMiddleware(handlers.RandomAnimals))
	http.HandleFunc("/colors/random", handlers.JSONMiddleware(handlers.RandomColors))
	http.HandleFunc("/names/random", handlers.JSONMiddleware(handlers.RandomNames))
	http.HandleFunc("/health", handlers.JSONMiddleware(handlers.Health))

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
