package main

import (
	"go-api/config"
	"go-api/routes"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
