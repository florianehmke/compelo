package handler

import (
	"net/http"

	"compelo/pkg/json"
)

type CreatePlayerRequest struct {
	Name string `json:"name"`
}

func (h *Handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	project := MustLoadProjectFromContext(r)
	var body CreatePlayerRequest
	if err := json.Unmarshal(r.Body, &body); err != nil {
		json.Error(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.svc.CreatePlayer(project.ID, body.Name)
	if err == nil {
		json.Write(w, http.StatusCreated, p)
	} else {
		json.Error(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	project := MustLoadProjectFromContext(r)
	players := h.svc.LoadPlayersByProjectID(project.ID)
	json.Write(w, http.StatusOK, players)
}
