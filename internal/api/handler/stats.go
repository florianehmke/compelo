package handler

import (
	"net/http"

	"compelo/pkg/json"
)

func (h *Handler) GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	game := MustLoadGameFromContext(r)

	players, err := h.svc.LoadPlayerStatsByGameID(game.ID)
	if err == nil {
		json.Write(w, http.StatusOK, players)
	} else {
		json.Error(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetGameStats(w http.ResponseWriter, r *http.Request) {
	game := MustLoadGameFromContext(r)

	players, err := h.svc.LoadGameStats(game.ID)
	if err == nil {
		json.Write(w, http.StatusOK, players)
	} else {
		json.Error(w, http.StatusBadRequest, err)
	}
}
