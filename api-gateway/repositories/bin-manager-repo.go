package repositories

import (
	"context"
	"log"

<<<<<<< HEAD:api-gateway/internal/repositories/bin-manager-repo.go
	binManagerGRPC "bin-manager/pkg/generated"

	"google.golang.org/protobuf/types/known/timestamppb"
=======
	binManager "api-gateway/generated/adapters"
>>>>>>> 0f101632433d7e39c16f7177f06a16a39a70e8fa:api-gateway/repositories/bin-manager-repo.go
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

	return response, nil
}

<<<<<<< HEAD:api-gateway/internal/repositories/bin-manager-repo.go
func (r *BinManagerRepo) LogRequest(binId string, hostname string, recieved *timestamppb.Timestamp, contents map[string]string) error {
	httpRequestDetails := &binManagerGRPC.LogRequestParams{
		BinId: binId,
		RequestToLog: &binManagerGRPC.HttpRequest{
			HostIp: hostname,
			Recieved: recieved,
			Contents: contents,
		},
=======
func (r *BinManagerRepo) LogRequest(binId string, requestContents map[string]string) error {
	httpRequestDetails := &binManager.LogRequestParams{
		BinId:        binId,
		RequestToLog: requestContents,
>>>>>>> 0f101632433d7e39c16f7177f06a16a39a70e8fa:api-gateway/repositories/bin-manager-repo.go
	}

	_, err := r.clientConn.LogRequestToBin(context.Background(), httpRequestDetails)
	if err != nil {
		log.Fatalf("Error when calling LogRequest: %s", err)
		return err
	}
	return nil
}

func (r *BinManagerRepo) FetchBinContents(binId string) (*binManager.FetchBinContentsResponse, error) {
	response, err := r.clientConn.FetchBinContents(context.Background(), &binManager.FetchBinContentsParams{
		BinId: binId,
	})
	if err != nil {
		log.Fatalf("Error when calling FetchBinContents: %s", err)
		return nil, err
	}
	log.Printf("Bin Contents: %+v", response)

	return response, nil
}
