package config_test

import (
	"os"
	"testing"

	"github.com/roflanKisel/cart-go/config"
)

var (
	ports = []struct {
		p        string
		expected string
	}{
		{"3000", ":3000"},
		{"3001", ":3001"},
		{"3002", ":3002"},
		{"", ":8080"},
	}

	dbs = []struct {
		name     string
		expected string
	}{
		{"test", "test"},
		{"test2", "test2"},
		{"", "cart_go"},
	}

	connectionStrings = []struct {
		c        string
		expected string
	}{
		{"mongodb://db:27017", "mongodb://db:27017"},
		{"", "mongodb://localhost:27017"},
	}
)

func TestAppPort(t *testing.T) {
	for _, p := range ports {
		os.Setenv("PORT", p.p)
		if port := config.AppPort(); port != p.expected {
			t.Errorf("AppPort(): expected: %v actual: %v", p.expected, port)
		}
	}
}

func TestDB(t *testing.T) {
	for _, db := range dbs {
		os.Setenv("DB", db.name)
		if name := config.DB(); name != db.expected {
			t.Errorf("DB(): expected: %v actual: %v", db.expected, name)
		}
	}
}

func TestConnectionString(t *testing.T) {
	for _, cs := range connectionStrings {
		os.Setenv("CONNECTION_STRING", cs.c)
		if c := config.ConnectionString(); c != cs.expected {
			t.Errorf("DB(): expected: %v actual: %v", cs.expected, c)
		}
	}
}
