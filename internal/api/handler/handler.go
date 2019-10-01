package handler

import (
	"compelo/internal"
)

type Handler struct {
	svc *compelo.Service
}

func New(svc *compelo.Service) *Handler {
	return &Handler{svc: svc}
}
