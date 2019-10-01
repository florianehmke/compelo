package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"compelo/internal"
)

type Handler struct {
	svc *compelo.Service
}

func New(svc *compelo.Service) *Handler {
	return &Handler{svc: svc}
}

func unmarshal(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(&v)
}

func marshal(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(&v)
}

func writeJSON(w http.ResponseWriter, code int, c interface{}) {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(b)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func writeError(w http.ResponseWriter, code int, err error) {
	body := map[string]string{
		"code":    fmt.Sprint(code),
		"message": err.Error(),
	}
	writeJSON(w, code, body)
}
