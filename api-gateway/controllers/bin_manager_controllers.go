package controllers

import (
	"api-gateway/repositories"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	hostname := r.RemoteAddr 
	recieved := timestamppb.Now()
	err := h.BinMgmtRepo.LogRequest(binId, hostname, recieved, requestDetails)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *BinMgmtBaseHandler) FetchBinContents(w http.ResponseWriter, r *http.Request) {
	binId := chi.URLParam(r, "binId")
	binContents, err := h.BinMgmtRepo.FetchBinContents(binId)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(nil)
	if err != nil {
		log.Fatal(err)
	}

	payload := &jsonResponse{
		Error: false,
		Data:  binContents,
	}
	statusCode := http.StatusAccepted

	WriteJSON(w, statusCode, payload)
}
