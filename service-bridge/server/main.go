package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"service-bridge/generated/adapters"
	html "service-bridge/rest-api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverPort = "8080"
	binManagerGrpcPort = "9000"
)


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

	binManagerPort := fmt.Sprintf(":%s", binManagerGrpcPort)
	conn, err := grpc.Dial(binManagerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to: %v", err)
	}
	defer conn.Close()


	log.Printf("Service-Bridge connected to bin manager grpc server on port %s", binManagerGrpcPort)
	client := binManager.NewBinManagerClient(conn)

	response, err := client.GenerateNewBin(context.Background(), &binManager.Params{})	
	if err != nil {
		log.Fatalf("Error when calling GenerateNewBin: %s", err)
	}
	log.Printf("Response from server: %s", response.BinId)
}