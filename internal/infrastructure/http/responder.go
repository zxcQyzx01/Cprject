package http

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	JSON(w http.ResponseWriter, status int, data interface{})
	Error(w http.ResponseWriter, status int, message string)
}

type HTTPResponder struct{}

func NewHTTPResponder() *HTTPResponder {
	return &HTTPResponder{}
}

func (r *HTTPResponder) JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (r *HTTPResponder) Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
