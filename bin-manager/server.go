package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"bin-manager/generated/adapters"
	"bin-manager/services"
	"bin-manager/db-service"
	"google.golang.org/grpc"
)

const gpcPort = "65535"
const dbPort = "27017"
const dbIP = "mongo_db"

type BinMgmtServer struct {
	binManager.UnimplementedBinManagerServer
}


func main() {
	fmt.Printf("Connecting to MongoDB...")
	db_service.CreateDbConn(dbIP, dbPort)	

	// essentially this is openning the port for the grpc server
	// and establishing a tcp listener for it
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", gpcPort))
	if err != nil {
		log.Fatalf("failed to enable a listener for the grpc server: %v", err)
	}

	grpcServer := grpc.NewServer()

	srv := BinMgmtServer{}
	binManager.RegisterBinManagerServer(grpcServer, &srv)

	fmt.Printf("Starting Bin Management service on port %s\n", gpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func GenerateNewBin(ctx context.Context, params *binManager.Params) (*binManager.BinId, error) {
	binId, err := services.CreateNewBinId() 
	if err != nil {
		log.Fatal("failed to create a new bin id", err) 
	}
	
	payload := binManager.BinId{
		BinId: binId,
	}
	return &payload, nil 
}

