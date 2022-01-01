package handler

import (
	"compelo/api/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	players, err := h.q.GetPlayerStatsBy(chi.URLParam(r, ProjectGUID), chi.URLParam(r, GameGUID))
	if err == nil {
		json.WriteResponse(w, http.StatusOK, players)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetGameStats(w http.ResponseWriter, r *http.Request) {
	players, err := h.q.GetGameStatsBy(chi.URLParam(r, ProjectGUID), chi.URLParam(r, GameGUID))
	if err == nil {
		json.WriteResponse(w, http.StatusOK, players)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}
