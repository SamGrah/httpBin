package binManager

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	binManager "bin-manager/pkg/generated"
)

type BinMgmtConn struct {
	Conn   *grpc.ClientConn
	Client *binManager.BinManagerClient
}

const binManagerSrvAddr = "bin-manager:65535"

// ConnectToBinMgmtSrv provides a function for clients to connect to this bin-manager service
func ConnectToBinMgmtSrv() *BinMgmtConn {
	conn, err := grpc.Dial(binManagerSrvAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to: %v", err)
	}

	client := binManager.NewBinManagerClient(conn)
	return &BinMgmtConn{
		Conn:   conn,
		Client: &client,
	}
}

func (s *BinMgmtConn) CloseConn() {
	s.Conn.Close()
}
