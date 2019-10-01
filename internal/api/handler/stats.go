package handler

import (
	"net/http"

	"compelo/pkg/json"
)

func (h *Handler) GetAllPlayerStats(w http.ResponseWriter, r *http.Request) {
	game := mustLoadGameFromContext(r)

	players, err := h.svc.LoadPlayerStatsByGameID(game.ID)
	if err == nil {
		json.Write(w, http.StatusOK, players)
	} else {
		json.Error(w, http.StatusBadRequest, err)
	}
}
