package repositories

import (
	"context"
	"log"

	binManagerGRPC "bin-manager/pkg/generated"

	"google.golang.org/protobuf/types/known/timestamppb"
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

func (r *BinManagerRepo) LogRequest(binId string, hostname string, recieved *timestamppb.Timestamp, contents map[string]string) error {
	httpRequestDetails := &binManagerGRPC.LogRequestParams{
		BinId: binId,
		RequestToLog: &binManagerGRPC.HttpRequest{
			HostIp: hostname,
			Recieved: recieved,
			Contents: contents,
		},
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
