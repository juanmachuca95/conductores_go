package web

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Results interface{} `json:"resulst,omitempty"`
	Status  int         `json:"status,omitempty"`
	Success bool        `json:"success,omitempty"`
}

func (r *Response) NewResponse(results interface{}, status int) *Response {
	return &Response{
		Results: results,
		Status:  status,
		Success: true,
	}
}

func (r *Response) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}
