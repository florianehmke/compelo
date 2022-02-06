package handler

import (
	"compelo/api/json"
	"compelo/command"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	MatchGUID string = "matchGUID"
)

type CreateMatchRequest struct {
	Teams []CreateMatchRequestTeam `json:"teams"`
}

type CreateMatchRequestTeam struct {
	PlayerGUIDs []string `json:"playerGuids" `
	Score       int      `json:"score"`
}

func (h *Handler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	var request CreateMatchRequest
	if err := json.Unmarshal(r.Body, &request); err != nil {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	c := command.CreateNewMatchCommand{
		GameGUID:    chi.URLParam(r, GameGUID),
		ProjectGUID: chi.URLParam(r, ProjectGUID),
	}
	for _, t := range request.Teams {
		c.Teams = append(c.Teams, struct {
			PlayerGUIDs []string
			Score       int
		}{
			PlayerGUIDs: t.PlayerGUIDs,
			Score:       t.Score,
		})
	}

	m, err := h.c.CreateNewMatch(c)
	if err == nil {
		json.WriteResponse(w, http.StatusCreated, m)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	c := command.DeleteMatchCommand{
		ProjectGUID: chi.URLParam(r, ProjectGUID),
		GameGUID:    chi.URLParam(r, GameGUID),
		GUID:        chi.URLParam(r, MatchGUID),
	}

	m, err := h.c.DeleteMatch(c)
	if err == nil {
		json.WriteResponse(w, http.StatusOK, m)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllMatches(w http.ResponseWriter, r *http.Request) {
	matches, err := h.q.GetMatchesBy(chi.URLParam(r, ProjectGUID), chi.URLParam(r, GameGUID))
	if err == nil {
		json.WriteResponse(w, http.StatusOK, matches)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}
