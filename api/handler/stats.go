package handler

import (
	"net/http"
)

func (h *Handler) GetAllPlayerStats(w http.ResponseWriter, r *http.Request) {
	game := mustLoadGameFromContext(r)

	players, err := h.svc.LoadPlayerStatsByGameID(game.ID)
	if err == nil {
		writeJSON(w, http.StatusCreated, players)
	} else {
		writeError(w, http.StatusBadRequest, err)
	}
}
