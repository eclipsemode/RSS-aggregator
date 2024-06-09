package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server listening on port %s", port)
	err = srv.ListenAndServe()
	if err != nil {
		return
	}
}
