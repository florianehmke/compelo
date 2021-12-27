package handler

import (
	"compelo/command"
	"compelo/query"
)

type Handler struct {
	q *query.Compelo
	c *command.Compelo
}

func New(q *query.Compelo, c *command.Compelo) *Handler {
	return &Handler{q: q, c: c}
}
