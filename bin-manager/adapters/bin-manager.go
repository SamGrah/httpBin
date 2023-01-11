package adapters

import (
	"context"
	"log"

	"bin-manager/services"
	"bin-manager/generated/adapters"
)

type BinMgmtServer struct {
	binManager.UnimplementedBinManagerServer
}

func (s *BinMgmtServer) GenerateNewBin(ctx context.Context, params *binManager.Params) (*binManager.BinId, error) {
	binId, err := services.CreateNewBin() 
	if err != nil {
		log.Fatal("failed to create a new bin id", err) 
	}
	
	payload := binManager.BinId{
		BinId: binId,
	}
	return &payload, nil 
}