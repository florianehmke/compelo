package query

import (
	"compelo/event"
	"log"
)

type playerStatsHandler struct {
	*Compelo // embedded root query handler
}

func (h *playerStatsHandler) on(e event.Event) {
	log.Println("[query:player-stats] handling event", e.GetID(), e.EventType())
}
