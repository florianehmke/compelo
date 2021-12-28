package handler

import (
	"net/http"

	"compelo/api/json"
	"compelo/command"

	"github.com/go-chi/chi"
)

type CreatePlayerRequest struct {
	Name string `json:"name"`
}

func (h *Handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var request CreatePlayerRequest
	if err := json.Unmarshal(r.Body, &request); err != nil {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.c.CreateNewPlayer(command.CreateNewPlayerCommand{
		ProjectGUID: chi.URLParam(r, ProjectGUID),
		Name:        request.Name,
	})
	if err == nil {
		json.WriteResponse(w, http.StatusCreated, p)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := h.q.GetPlayersBy(chi.URLParam(r, ProjectGUID))

	if err == nil {
		json.WriteResponse(w, http.StatusOK, players)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}
