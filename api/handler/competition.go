package handler

import (
	"net/http"

	"compelo/api/json"
	"compelo/command"

	"github.com/go-chi/chi"
)

const (
	CompetitonGUID string = "competitionGUID"
)

type CreateCompetitionRequest struct {
	Rounds int                            `json:"rounds"`
	Name   string                         `json:"name"`
	Teams  []CreateCompetitionRequestTeam `json:"teams"`
}

type CreateCompetitionRequestTeam struct {
	PlayerGUIDs []string `json:"playerGuids" `
}

func (h *Handler) CreateCompetition(w http.ResponseWriter, r *http.Request) {
	var request CreateCompetitionRequest
	if err := json.Unmarshal(r.Body, &request); err != nil {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	c := command.CreateNewCompetitionCommand{
		GameGUID:    chi.URLParam(r, GameGUID),
		ProjectGUID: chi.URLParam(r, ProjectGUID),
		Name:        request.Name,
		Rounds:      request.Rounds,
	}
	for _, t := range request.Teams {
		c.Teams = append(c.Teams, struct {
			PlayerGUIDs []string
		}{
			PlayerGUIDs: t.PlayerGUIDs,
		})
	}

	p, err := h.c.CreateNewCompetition(c)
	if err == nil {
		json.WriteResponse(w, http.StatusCreated, p)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllCompetitions(w http.ResponseWriter, r *http.Request) {
	matches, err := h.q.GetCompetitionsBy(chi.URLParam(r, ProjectGUID), chi.URLParam(r, GameGUID))
	if err == nil {
		json.WriteResponse(w, http.StatusOK, matches)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}
