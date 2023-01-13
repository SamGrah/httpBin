package repositories

import (
	"context"
	"log"
	binManager "service-bridge/generated/adapters"
	"service-bridge/models"
)

type BinManagerRepo struct {
	clientConn binManager.BinManagerClient
}

func NewBinManagerRepo(clientConn binManager.BinManagerClient) *BinManagerRepo {
	return &BinManagerRepo{
		clientConn: clientConn,
	}
}

func (r *BinManagerRepo) CreateNewBin() (*models.Bin, error) {
	response, err := r.clientConn.GenerateNewBin(context.Background(), &binManager.Params{})
	if err != nil {
		log.Fatalf("Error when calling GenerateNewBin: %s", err)
		return nil, err
	}
	log.Printf("Response from server: %s", response.BinId)

	return &models.Bin{
		Id: response.BinId,
		Contents: nil,
	}, nil 
}





