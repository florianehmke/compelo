package query

import (
	"compelo/event"
	"log"
)

type gameStatsHandler struct {
	*Compelo // embedded root query handler
}

func (h *gameStatsHandler) on(e event.Event) {
	log.Println("[query:game-stats] handling event", e.GetID(), e.EventType())
}
