package adapters

import (
	"context"
	"log"

	"bin-manager/generated/adapters"
	"bin-manager/services"
)

type BinMgmtServer struct {
	binManager.UnimplementedBinManagerServer
}

func (s *BinMgmtServer) GenerateNewBin(ctx context.Context, params *binManager.Params) (*binManager.NewBinResponse, error) {
	binId, err := services.CreateNewBin() 
	if err != nil {
		log.Fatal("failed to create a new bin id", err) 
	}
	
	payload := binManager.NewBinResponse{
		BinId: binId,
	}
	return &payload, nil 
}

func (s *BinMgmtServer) LogRequestToBin (ctx context.Context, request *binManager.HttpRequestDetails) (*binManager.LogRequestResponse, error) {
	err := services.LogRequestToBin(request.BinId, request.HttpRequestContents)

	if err != nil {
		log.Fatal("failed to log request to bin", err)
	}

	payload := binManager.LogRequestResponse{}
	return &payload, err 
}