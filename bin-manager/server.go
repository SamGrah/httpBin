package main

import (
	"fmt"
	"log"
	"net"

	db_service "bin-manager/db-service"
	"bin-manager/adapters"
	"bin-manager/generated/adapters"

	"google.golang.org/grpc"
)

const gpcPort = "65535"

// TODO: replace log.Fatal with a proper error handling
func main() {
	db_service.Init()

	// start grpc server and establish a tcp listener for it
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", gpcPort))
	if err != nil {
		log.Fatalf("failed to enable a listener for the grpc server: %v", err)
	}

	grpcServer := grpc.NewServer()

	srv := adapters.BinMgmtServer{}
	binManager.RegisterBinManagerServer(grpcServer, &srv)

	fmt.Printf("Starting Bin Management service on port %s\n", gpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
