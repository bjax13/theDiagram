package main

import (
	"os"

	"github.com/bjax13/theDiagram/api"
)

func main() {

	port := getenvOrDefault("PORT", "1337")

	api := api.API{
		Port: port,
	}

	api.RunServer()
}

func getenvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
