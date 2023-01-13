package adapters

import (
	"log"

	"service-bridge/generated/adapters"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BinMgmtConn struct {
	Conn *grpc.ClientConn
	Client *binManager.BinManagerClient
}

const binManagerSrvAddr = "bin-manager:65535"

func ConnectToBinMgmtSrv() *BinMgmtConn {
	conn, err := grpc.Dial(binManagerSrvAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to: %v", err)
	}

	client := binManager.NewBinManagerClient(conn)
	return &BinMgmtConn{
		Conn: conn,
		Client: &client,
	}
}

func (s *BinMgmtConn) CloseConn() {
	s.Conn.Close()
}

