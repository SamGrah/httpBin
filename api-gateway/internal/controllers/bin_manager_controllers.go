package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/types/known/timestamppb"

	"api-gateway/internal/repositories"
	"api-gateway/pkg/utils"
	binManager "bin-manager/pkg/generated"
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
		log.Printf("Error when calling GenerateNewBin: %s", err)
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

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
	hostIp := r.RemoteAddr
	timestamp := timestamppb.Now()

	err := h.BinMgmtRepo.LogRequest(binId, &binManager.HttpRequest{
		HostIp:   hostIp,
		Recieved: timestamp,
		Contents: requestDetails,
	})
	if err != nil {
		log.Printf("Failed to log request to bin %s", binId)
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(nil)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to write http response for bin %s", binId)
		utils.LogError(errMsg, err)
	}
}

func (h *BinMgmtBaseHandler) FetchBinContents(w http.ResponseWriter, r *http.Request) {
	binId := chi.URLParam(r, "binId")
	binContents, err := h.BinMgmtRepo.FetchBinContents(binId)
	if err != nil {
		log.Printf("Failed to fetch bin contents for bin %s", binId)
		// w.WriteHeader(http.StatusBadRequest)
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	// _, err = w.Write(nil)
	// if err != nil {
	// 	log.Printf("Failed to write http response for bin %s", binId)
	// }

	payload := &jsonResponse{
		Error: false,
		Data:  binContents,
	}
	statusCode := http.StatusAccepted

	WriteJSON(w, statusCode, payload)
}
