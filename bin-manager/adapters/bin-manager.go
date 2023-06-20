package adapters

import (
	"context"
	"log"

<<<<<<< HEAD:bin-manager/internal/adapters/bin-manager.go
	db_service "bin-manager/internal/db-service"
	"bin-manager/internal/services"
	binManager "bin-manager/pkg/generated"
=======
	db_service "bin-manager/db-service"
	binManager "bin-manager/generated/adapters"
	"bin-manager/services"
>>>>>>> 0f101632433d7e39c16f7177f06a16a39a70e8fa:bin-manager/adapters/bin-manager.go
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
	var httpRequest *db_service.HttpRequest = &db_service.HttpRequest{
		HostIp: params.RequestToLog.HostIp,
		Recieved: params.RequestToLog.Recieved,
		Contents: params.RequestToLog.Contents,
	}

	err := services.LogRequestToBin(params.BinId, httpRequest)

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

	binContentsResponse := make([]*binManager.HttpRequest, 0)
	for _, requestContents := range *binContents {
		binContentsResponse = append(binContentsResponse, &binManager.HttpRequest{
			HostIp: requestContents.HostIp,
			Recieved: requestContents.Recieved,
			Contents: requestContents.Contents,
		})
	}

	payload := binManager.FetchBinContentsResponse{
		BinContents: binContentsResponse,
	}

	return &payload, nil
}
