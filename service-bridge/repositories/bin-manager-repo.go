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

func (r *BinManagerRepo) CreateNewBin() (*binManager.NewBinResponse, error) {
	response, err := r.clientConn.GenerateNewBin(context.Background(), &binManager.Params{})
	if err != nil {
		log.Fatalf("Error when calling GenerateNewBin: %s", err)
		return nil, err
	}
	log.Printf("Response from server: %s", response.BinId)

	return &binManager.NewBinResponse{
		BinId:               response.BinId,
	}, nil
}

func (r *BinManagerRepo) LogRequest(binId string, requestContents map[string]string) error {
	httpRequestDetails := &binManager.LogRequestParams{
		BinId:	binId,
		RequestToLog: requestContents,
	}

	_, err := r.clientConn.LogRequestToBin(context.Background(), httpRequestDetails)
	if err != nil {
		log.Fatalf("Error when calling LogRequest: %s", err)
		return err
	}
	return nil
}
