package controllers

import (
	"log"
	"net/http"
	"service-bridge/models"
	"service-bridge/helpers"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Data any `json:"bin"`
}

type BinMgmtBaseHandler struct {
	BinMgmtRepo models.BinMgmtRepository
}

func NewBinMgmtBaseHandler(binMgmtRepo models.BinMgmtRepository) *BinMgmtBaseHandler {
	return &BinMgmtBaseHandler{
		BinMgmtRepo: binMgmtRepo,
	}
}

func (h *BinMgmtBaseHandler) CreateNewBin(w http.ResponseWriter, r *http.Request) {
	response, err := h.BinMgmtRepo.CreateNewBin()
	if err != nil {
		log.Fatalf("Error when calling GenerateNewBin: %s", err)
	}
	log.Printf("Response from server: %s", response.Id)

	payload := jsonResponse{
		Error: false,
		Data:  response,
	}
	statusCode := http.StatusAccepted

	helpers.WriteJSON(w, statusCode, payload)
}