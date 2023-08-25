package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	} else {
		port = "5000"
	}

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello World!!"))
	})
	log.Fatal(http.ListenAndServe(":"+port, serverMux))
}
