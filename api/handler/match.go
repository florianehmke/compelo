package handler

import (
	"net/http"
	"time"

	"compelo/internal/compelo"
)

func (h *Handler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	game := mustLoadGameFromContext(r)

	var param compelo.CreateMatchParameter
	if err := unmarshal(r.Body, &param); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	param.GameID = game.ID
	param.Date = time.Now()

	m, err := h.svc.CreateMatch(param)
	if err == nil {
		writeJSON(w, http.StatusCreated, m)
	} else {
		writeError(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllMatches(w http.ResponseWriter, r *http.Request) {
	game := mustLoadGameFromContext(r)

	matches, err := h.svc.LoadMatchesByGameID(game.ID)
	if err == nil {
		writeJSON(w, http.StatusOK, matches)
	} else {
		writeError(w, http.StatusBadRequest, err)
	}
}
