package main

import (
	"fmt"
	"log"
	"net/http"

	html "service-bridge/rest-api"
)

const serverPort = "8080"

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", serverPort),
		Handler: html.Routes(),
	}

	log.Printf("Service-Bridge listening for HTTP requests on port %s\n", serverPort)
	
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}