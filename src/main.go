package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"api/src/router"
)

func init() {
	if envLoadError := godotenv.Load(); envLoadError != nil {
		log.Fatal("[ ERROR ] Failed to load .env file")
	}
}

func main() {
	var PORT string

	router := router.RegisterRoutes()

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "9090"
	}

	// Configure CORS middleware
	corsHandler := cors.Default().Handler(router)

	fmt.Printf("[ OK ] Server is Started and Listening on port: %v", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, corsHandler))
}
