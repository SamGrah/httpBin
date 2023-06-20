package repositories

import (
	"context"
	"log"

	binManagerGRPC "bin-manager/pkg/generated"
)

type BinManagerRepo struct {
	clientConn binManagerGRPC.BinManagerClient
}

func NewBinManagerRepo(clientConn binManagerGRPC.BinManagerClient) *BinManagerRepo {
	return &BinManagerRepo{
		clientConn: clientConn,
	}
}

func (r *BinManagerRepo) CreateNewBin() (*binManagerGRPC.NewBinResponse, error) {
	response, err := r.clientConn.GenerateNewBin(context.Background(), &binManagerGRPC.Params{})
	if err != nil {
		log.Fatalf("Error when calling GenerateNewBin: %s", err)
		return nil, err
	}
	log.Printf("Response from server: %s", response.BinId)

	return response, nil
}

func (r *BinManagerRepo) LogRequest(binId string, httpRequest *binManagerGRPC.HttpRequest) error {
	httpRequestDetails := &binManagerGRPC.LogRequestParams{
		BinId:        binId,
		RequestToLog: httpRequest,
	}

	_, err := r.clientConn.LogRequestToBin(context.Background(), httpRequestDetails)
	if err != nil {
		log.Fatalf("Error when calling LogRequest: %s", err)
		return err
	}
	return nil
}

func (r *BinManagerRepo) FetchBinContents(binId string) (*binManagerGRPC.FetchBinContentsResponse, error) {
	response, err := r.clientConn.FetchBinContents(context.Background(), &binManagerGRPC.FetchBinContentsParams{
		BinId: binId,
	})
	if err != nil {
		log.Fatalf("Error when calling FetchBinContents: %s", err)
		return nil, err
	}
	log.Printf("Bin Contents: %+v", response)

	return response, nil
}
