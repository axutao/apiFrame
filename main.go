package main

import (
	"log"
	"net/http"
	"route"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":9091", router))
}