package config

import (
	"fmt"
	"os"
)

const (
	connectionString = "mongodb://localhost:27017"
	databaseName     = "cart_go"
	port             = ":8080"
)

// AppPort returns application HTTP port.
func AppPort() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}

	return port
}

// DB returns database name which will be used in application.
func DB() string {
	if dbName := os.Getenv("DB"); dbName != "" {
		return dbName
	}

	return databaseName
}

// ConnectionString returns connection string to mongodb.
func ConnectionString() string {
	if dbCS := os.Getenv("CONNECTION_STRING"); dbCS != "" {
		return fmt.Sprintf(":%s", dbCS)
	}

	return connectionString
}
