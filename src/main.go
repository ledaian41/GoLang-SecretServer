package main

import (
	swagger "github.com/ledaian41/GoLang-SecretServerAPI/src/go"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")
	router := swagger.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
