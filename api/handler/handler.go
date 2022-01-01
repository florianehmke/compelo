package handler

import (
	"compelo/command"
	"compelo/query"
)

type Handler struct {
	q *query.Service
	c *command.Service
}

func New(q *query.Service, c *command.Service) *Handler {
	return &Handler{q: q, c: c}
}
