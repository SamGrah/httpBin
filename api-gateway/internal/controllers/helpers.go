package controllers

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Data any 
}

func WriteJSON(w http.ResponseWriter, status int, data *jsonResponse) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	
	return nil
}
