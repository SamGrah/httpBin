package endpoints 

import (
	"net/http"

	"api-gateway/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Handlers struct {
	BinMgmtHandler *controllers.BinMgmtBaseHandler
}

func Routes(h *Handlers) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))
	
	router.Post("/new-bin", h.BinMgmtHandler.CreateNewBin)
	router.Post("/bin/{binId}", h.BinMgmtHandler.LogRequest)
	router.Get("/view-bin/{binId}", h.BinMgmtHandler.FetchBinContents)

	return router
}