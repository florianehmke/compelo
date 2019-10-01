package handler

import (
	"net/http"
	"time"

	"compelo/internal"
	"compelo/pkg/json"
)

func (h *Handler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	game := mustLoadGameFromContext(r)

	var param compelo.CreateMatchParameter
	if err := json.Unmarshal(r.Body, &param); err != nil {
		json.Error(w, http.StatusBadRequest, err)
		return
	}

	param.GameID = game.ID
	param.Date = time.Now()

	m, err := h.svc.CreateMatch(param)
	if err == nil {
		json.Write(w, http.StatusCreated, m)
	} else {
		json.Error(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllMatches(w http.ResponseWriter, r *http.Request) {
	game := mustLoadGameFromContext(r)

	matches, err := h.svc.LoadMatchesByGameID(game.ID)
	if err == nil {
		json.Write(w, http.StatusOK, matches)
	} else {
		json.Error(w, http.StatusBadRequest, err)
	}
}
