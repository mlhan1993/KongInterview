package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	v1Handlers "github.com/mlhan1993/KongInterview/pkg/handlers/v1"
)

func main() {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new HTTP server
	server := http.NewServeMux()

	// Register the v1 handler functions
	v1 := v1Handlers.NewV1(db)
	server.HandleFunc("/v1/service/overview", v1.PostRetrieveServiceOverview)
	server.HandleFunc("/v1/service/details", v1.PostRetrieveServiceDetails)

	// Start the server
	log.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}
