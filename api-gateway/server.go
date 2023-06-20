package main

import (
	"log"
	"net/http"

	binManager "bin-manager/pkg"

	"api-gateway/internal/controllers"
	"api-gateway/internal/repositories"
	html "api-gateway/internal/endpoints"
)

const (
	httpServerPort   = ":8080"
	grpcListenerAddr = "0.0.0.0:65535"
	binMangerGrpcSvc = "bin-manager:65535"
)

func main() {
	binManagerSrv := binManager.ConnectToBinMgmtSrv()
	defer binManagerSrv.CloseConn()

	binManagerRepo := repositories.NewBinManagerRepo(*binManagerSrv.Client)

	binMgmtHandler := controllers.NewBinMgmtBaseHandler(binManagerRepo)
	handlersRegister := &html.Handlers{
		BinMgmtHandler: binMgmtHandler,
	}

	httpServer := &http.Server{
		Addr:    httpServerPort,
		Handler: html.Routes(handlersRegister),
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("api-gateway listening for HTTP requests on port %s\n", httpServerPort)
}
