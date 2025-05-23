package main

import (
	"kreditplus/config"
	"kreditplus/controllers/consument"
	"kreditplus/controllers/dashboard"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", dashboard.Index)	

	http.HandleFunc("/consuments", consument.Index)	
	http.HandleFunc("/consuments/add", consument.Add)	
	http.HandleFunc("/consuments/delete", consument.Delete)	

	log.Println("Server running on port 8000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8000", nil);
}