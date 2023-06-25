package repositories

import (
	"context"
	"log"

	"api-gateway/pkg/utils"
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
		utils.LogError("Error when calling GenerateNewBin", err)
		return nil, err
	}
	log.Printf("New bin created: %s", response.BinId)

	return response, nil
}

func (r *BinManagerRepo) LogRequest(binId string, httpRequest *binManagerGRPC.HttpRequest) error {
	httpRequestDetails := &binManagerGRPC.LogRequestParams{
		BinId:        binId,
		RequestToLog: httpRequest,
	}

	response , err := r.clientConn.LogRequestToBin(context.Background(), httpRequestDetails)
	if err != nil {
		utils.LogError("Error when calling LogRequest", err)
		return err
	}
	log.Printf("Request logged to bin %s: %+v", binId, response)

	return nil
}

func (r *BinManagerRepo) FetchBinContents(binId string) (*binManagerGRPC.FetchBinContentsResponse, error) {
	response, err := r.clientConn.FetchBinContents(context.Background(), &binManagerGRPC.FetchBinContentsParams{
		BinId: binId,
	})
	if err != nil {
		utils.LogError("Error when calling FetchBinContents", err)
		return nil, err
	}
	log.Printf("Bin contents fetched: %+v", response)

	return response, nil
}
