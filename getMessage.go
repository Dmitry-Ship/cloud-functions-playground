// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Message string
}
type Response struct {
	Data string `json:"data"`
}

func processRequest(r Request) Response {
	if r.Message == "" {
		return Response{
			Data: "empty message",
		}
	}

	return Response{
		Data: r.Message,
	}
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var request Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	data := processRequest(request)

	json.NewEncoder(w).Encode(&data)
}
