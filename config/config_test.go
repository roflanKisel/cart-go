package config_test

import (
	"os"
	"testing"

	"github.com/roflanKisel/cart-go/config"
)

func TestAppPort(t *testing.T) {
	ports := []struct {
		name     string
		port     string
		expected string
	}{
		{"Passed port as ENV variable", "3000", ":3000"},
		{"Default port", "", ":8080"},
	}

	for _, p := range ports {
		t.Run(p.name, func(t *testing.T) {
			os.Setenv("PORT", p.port)
			if port := config.AppPort(); port != p.expected {
				t.Fatalf("expected: %v actual: %v", p.expected, port)
			}
		})
	}
}

func TestDB(t *testing.T) {
	dbs := []struct {
		name     string
		dbName   string
		expected string
	}{
		{"Passed DB name as ENV variable", "test", "test"},
		{"Default DB name", "", "cart_go"},
	}

	for _, db := range dbs {
		t.Run(db.name, func(t *testing.T) {
			os.Setenv("DB", db.dbName)
			if dbName := config.DB(); dbName != db.expected {
				t.Fatalf("expected: %v actual: %v", db.expected, dbName)
			}
		})
	}
}

func TestConnectionString(t *testing.T) {
	connectionStrings := []struct {
		name             string
		connectionString string
		expected         string
	}{
		{"Passed connection string as ENV variable", "mongodb://db:27017", "mongodb://db:27017"},
		{"Default connection string", "", "mongodb://localhost:27017"},
	}

	for _, cs := range connectionStrings {
		t.Run(cs.name, func(t *testing.T) {
			os.Setenv("CONNECTION_STRING", cs.connectionString)
			if c := config.ConnectionString(); c != cs.expected {
				t.Fatalf("DB(): expected: %v actual: %v", cs.expected, c)
			}
		})
	}
}
