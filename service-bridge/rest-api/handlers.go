package adapters

import (
	"context"
	"fmt"
	"log"
	"net/http"

	binManager "service-bridge/generated/adapters"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const	binManagerGrpcPort = "65535"

func CreateBin(w http.ResponseWriter, r *http.Request) {
	binManagerPort := fmt.Sprintf("localhost:%s", binManagerGrpcPort)
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

	payload := jsonResponse{
		Error: false,
		Data:  response,
	}
	statusCode := http.StatusAccepted

	WriteJSON(w, statusCode, payload)
}
