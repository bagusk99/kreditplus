package main

import (
	"kreditplus/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	log.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil);
}