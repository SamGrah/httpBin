package adapters

import "net/http"

func CreateBin(w http.ResponseWriter, r *http.Request) {
	binId := "vax7pq"

	payload := jsonResponse{
		Error: false,
		Data: binId,
	}

	statusCode := http.StatusAccepted

	WriteJSON(w, statusCode, payload)
}