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

func (r *BinManagerRepo) CreateNewBin() (*models.BinId, error) {
	response, err := r.clientConn.GenerateNewBin(context.Background(), &binManager.Params{})
	if err != nil {
		log.Fatalf("Error when calling GenerateNewBin: %s", err)
		return nil, err
	}
	log.Printf("Response from server: %s", response.BinId)

	return &models.BinId{
		BinId:	response.BinId,
	}, nil
}

func (r *BinManagerRepo) LogRequest(binId *models.BinId, requestToLog *models.HttpRequest) error {
	httpRequestDetails := &binManager.LogRequestParams{
		BinId:        binId.BinId,
		RequestToLog: requestToLog.Contents,
	}

	_, err := r.clientConn.LogRequestToBin(context.Background(), httpRequestDetails)
	if err != nil {
		log.Fatalf("Error when calling LogRequest: %s", err)
		return err
	}
	return nil
}

func (r *BinManagerRepo) FetchBinContents(binId models.BinId) (*models.Bin, error) {
	binToFetch := &binManager.FetchBinContentsParams{
		BinId: binId.BinId,
	}

	binContents, err := r.clientConn.FetchBinContents(context.Background(), binToFetch)
	if err != nil {
		log.Fatal(err)
	}

	contents := make([]*models.HttpRequest, 0)
	for key, value := range binContents.BinContents {
		contents[key] = &models.HttpRequest{value.Contents}
	}

	return &models.Bin{
		BinId: models.BinId{binId.BinId},
		Contents: contents,
	}, nil	
}
