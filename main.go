package main

import (
	"kreditplus/config"
	"kreditplus/controllers/consument"
	"kreditplus/controllers/dashboard"
	"kreditplus/controllers/transaction"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", dashboard.Index)	

	http.HandleFunc("/consuments", consument.Index)	
	http.HandleFunc("/consuments/add", consument.Add)	
	http.HandleFunc("/consuments/delete", consument.Delete)	
	http.HandleFunc("/consuments/edit", consument.Edit)	

	http.HandleFunc("/transactions", transaction.Index)	
	http.HandleFunc("/transactions/add", transaction.Add)	
	http.HandleFunc("/transactions/delete", transaction.Delete)	
	http.HandleFunc("/transactions/edit", transaction.Edit)	

	log.Println("Server running on port 8000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8000", nil);
}