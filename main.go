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

	http.HandleFunc("/plants", handlers.AllPlants)
	http.HandleFunc("/animals", handlers.AllAnimals)
	http.HandleFunc("/colors", handlers.AllColors)
	http.HandleFunc("/names", handlers.AllNames)
	http.HandleFunc("/plants/random", handlers.RandomPlants)
	http.HandleFunc("/animals/random", handlers.RandomAnimals)
	http.HandleFunc("/colors/random", handlers.RandomColors)
	http.HandleFunc("/names/random", handlers.RandomNames)
	http.HandleFunc("/health", handlers.Health)

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
