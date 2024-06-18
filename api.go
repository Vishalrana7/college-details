package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func(s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetCollegeDetail)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *ApiServer) handleGetCollegeDetail(w http.ResponseWriter, r *http.Request) {
	// Extract the 'country' parameter from the request
	country := r.URL.Query().Get("country")
	if country == "" {
		country = "india" // Default value
	}

	// Call the service method with the country parameter
	name, err := s.svc.GetCollegeDetail(context.Background(), country)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, name)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}






