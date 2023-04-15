package config

import (
	"fmt"
	"os"
)

func GetDBURI() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db := os.Getenv("DB_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
}
