package handler

import (
	"net/http"

	"compelo/api/json"
	"compelo/command"
)

type CreatePlayerRequest struct {
	Name string `json:"name"`
}

func (h *Handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	project := MustLoadProjectFromContext(r)

	var request CreatePlayerRequest
	if err := json.Unmarshal(r.Body, &request); err != nil {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.c.CreateNewPlayer(command.CreateNewPlayerCommand{
		ProjectGUID: project.GUID,
		Name:        request.Name,
	})
	if err == nil {
		json.WriteResponse(w, http.StatusCreated, p)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	project := MustLoadProjectFromContext(r)
	players, err := h.q.GetPlayersBy(project.GUID)

	if err == nil {
		json.WriteResponse(w, http.StatusOK, players)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}
