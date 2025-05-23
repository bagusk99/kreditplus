package main

import (
	"kreditplus/config"
	DashboardController "kreditplus/controllers"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", DashboardController.Index)	

	log.Println("Server running on port 8000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8000", nil);
}