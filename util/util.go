package util

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data interface{} `json:"data"`
}

func ErrorResp(w http.ResponseWriter, message interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	msg := message.(string)
	json.NewEncoder(w).Encode(APIResponse{msg})
}

func SuccessResp(w http.ResponseWriter, message interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch message.(type) {
	case string:
		json.NewEncoder(w).Encode(APIResponse{message.(string)})
		break

	default:
		json.NewEncoder(w).Encode(APIResponse{message})
	}

}
