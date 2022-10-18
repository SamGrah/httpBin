package main

import (
	"fmt"
	"log"
	"net"

	"bin-manager/generated/adapters"

	"google.golang.org/grpc"
)

const gpcPort = "65535"

type binManagerServer struct {
	binManager.UnimplementedBinManagerServer
}

func main() {
	// essentially this is openning the port for the grpc server
	// and establishing a tcp listener for it
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", gpcPort))
	if err != nil {
		log.Fatalf("failed to enable a listener for the grpc server: %v", err)
	}

	grpcServer := grpc.NewServer()

	srv := binManagerServer{}
	binManager.RegisterBinManagerServer(grpcServer, &srv)

	fmt.Printf("Starting Bin Management service on port %s\n", gpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}