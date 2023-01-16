package controllers

import (
	"io"
	"log"
	"net/http"
	"service-bridge/repositories"

	"github.com/go-chi/chi/v5"
)

type BinMgmtBaseHandler struct {
	BinMgmtRepo *repositories.BinManagerRepo
}

func NewBinMgmtBaseHandler(binMgmtRepo *repositories.BinManagerRepo) *BinMgmtBaseHandler {
	return &BinMgmtBaseHandler{
		BinMgmtRepo: binMgmtRepo,
	}
}

func (h *BinMgmtBaseHandler) CreateNewBin(w http.ResponseWriter, r *http.Request) {
	response, err := h.BinMgmtRepo.CreateNewBin()
	if err != nil {
		log.Fatalf("Error when calling GenerateNewBin: %s", err)
	}
	log.Printf("Response from server: %s", response.BinId)

	payload := &jsonResponse{
		Error: false,
		Data:  response,
	}
	statusCode := http.StatusAccepted

	WriteJSON(w, statusCode, payload)
}

func (h *BinMgmtBaseHandler) LogRequest(w http.ResponseWriter, r *http.Request) {
	requestDetails := map[string]string{}
	headersString := "body"
	body, _ := io.ReadAll(r.Body)
	requestDetails[headersString] = string(body)

	for name, headers := range r.Header {
		for _, hdr := range headers {
			requestDetails[name] = hdr
		}
	}

	binId := chi.URLParam(r, "binId")

	err := h.BinMgmtRepo.LogRequest(binId, requestDetails)
	if err != nil {
		log.Fatal(err)
	}

	payload := &jsonResponse{
		Error: false,
		Data:  requestDetails,
	}
	statusCode := http.StatusAccepted

	WriteJSON(w, statusCode, payload)
}
