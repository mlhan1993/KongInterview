package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mlhan1993/KongInterview/pkg/config"
	v1Handlers "github.com/mlhan1993/KongInterview/pkg/handlers/v1"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/mlhan1993/KongInterview/pkg/db"
	"github.com/mlhan1993/KongInterview/pkg/middlewares"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	dbURI := config.GetDBURI()
	// Connect to the MySQL database
	conn, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	serviceDb, err := db.NewKong(conn)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new HTTP server
	server := mux.NewRouter()

	// Middlewares
	server.Use(middlewares.RequestIDMiddleware)
	server.Use(middlewares.LogRequestMiddleware)

	// Register the v1 handler functions
	v1 := v1Handlers.NewV1(serviceDb)
	server.HandleFunc("/v1/service", v1.PostRetrieveServices)
	server.HandleFunc("/v1/version", v1.PostRetrieveVersions)

	// Start the server
	log.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}

}
