package repositories

import (
	"context"
	"log"

	binManager "service-bridge/generated/adapters"
)

type BinManagerRepo struct {
	clientConn binManager.BinManagerClient
}

func NewBinManagerRepo(clientConn binManager.BinManagerClient) *BinManagerRepo {
	return &BinManagerRepo{
		clientConn: clientConn,
	}
}

func (r *BinManagerRepo) CreateNewBin() (*binManager.HttpRequestDetails, error) {
	response, err := r.clientConn.GenerateNewBin(context.Background(), &binManager.Params{})
	if err != nil {
		log.Fatalf("Error when calling GenerateNewBin: %s", err)
		return nil, err
	}
	log.Printf("Response from server: %s", response.BinId)

	return &binManager.HttpRequestDetails{
		BinId:               response.BinId,
		HttpRequestContents: nil,
	}, nil
}

func (r *BinManagerRepo) LogRequest(binId string, requestContents map[string]string) error {
	httpRequestDetails := &binManager.HttpRequestDetails{
		BinId:	binId,
		HttpRequestContents: requestContents,
	}

	_, err := r.clientConn.LogRequestToBin(context.Background(), httpRequestDetails)
	if err != nil {
		log.Fatalf("Error when calling LogRequest: %s", err)
		return err
	}
	return nil
}
