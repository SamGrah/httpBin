package adapters

import (
	"context"
	"log"

	db_service "bin-manager/db-service"
	binManager "bin-manager/generated/adapters"
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

func (s *BinMgmtServer) LogRequestToBin(ctx context.Context, params *binManager.LogRequestParams) (*binManager.LogRequestResponse, error) {
	var httpRequestContents db_service.HttpRequestContents = params.RequestToLog
	err := services.LogRequestToBin(params.BinId, &httpRequestContents)

	if err != nil {
		log.Fatal("failed to log request to bin", err)
	}

	payload := binManager.LogRequestResponse{}
	return &payload, nil
}

func (s *BinMgmtServer) FetchBinContents(ctx context.Context, params *binManager.FetchBinContentsParams) (*binManager.FetchBinContentsResponse, error) {
	binId := params.BinId

	binContents, err := services.FetchRequestsFromBin(binId)
	if err != nil {
		log.Fatal("failed to fetch bin request history")
	}

	binContentsResponse := []*binManager.HttpRequest{}
	for _, requestContents := range *binContents {
		binContentsResponse = append(binContentsResponse, &binManager.HttpRequest{
			Contents: requestContents,
		})
	}

	payload := binManager.FetchBinContentsResponse{
		BinContents: binContentsResponse,
	}

	return &payload, nil
}
