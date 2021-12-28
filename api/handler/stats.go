package handler

import (
	"compelo/api/json"
	"net/http"
)

func (h *Handler) GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	game := MustLoadGameFromContext(r)

	players, err := h.q.GetPlayerStatsBy(game.ProjectGUID, game.GUID)
	if err == nil {
		json.WriteResponse(w, http.StatusOK, players)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetGameStats(w http.ResponseWriter, r *http.Request) {
	game := MustLoadGameFromContext(r)

	players, err := h.q.GetGameStatsBy(game.ProjectGUID, game.GUID)
	if err == nil {
		json.WriteResponse(w, http.StatusOK, players)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}
