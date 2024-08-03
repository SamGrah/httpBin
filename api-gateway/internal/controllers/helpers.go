package controllers

import (
	"encoding/json"
	"net/http"

	"api-gateway/pkg/utils"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Data any 
}

func WriteJSON(w http.ResponseWriter, status int, data *jsonResponse) error {
	out, err := json.Marshal(data)
	if err != nil {
		utils.LogError("Failed to marshal json", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		utils.LogError("Failed to write json response", err)
		return err
	}
	
	return nil
}

func WriteError(w http.ResponseWriter, status int, err error) error {
	payload := &jsonResponse{
		Error: true,
		Data:  err.Error(),
	}

	return WriteJSON(w, status, payload)
} 
