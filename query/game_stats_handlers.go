package query

import (
	"compelo/event"
	"log"
)

type gameStatsHandler struct {
	c *Compelo
}

func (h *gameStatsHandler) on(e event.Event) {
	log.Println("[query:game-stats] handling event", e.GetID(), e.EventType())
}
