package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Basic HTTP WEB SERVER IN GOLANG //

// Request Handler
func indexhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!!"))
}

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
	serverMux.HandleFunc("/home", indexhandler)
	log.Fatal(http.ListenAndServe(":"+port, serverMux))
}
