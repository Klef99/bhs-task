package v1

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response{msg})
}
