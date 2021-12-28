package handler

import (
	"net/http"

	"compelo/api/json"
	"compelo/command"

	"github.com/go-chi/chi"
)

const (
	GameGUID string = "gameGUID"
)

type CreateGameRequest struct {
	Name string `json:"name"`
}

func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	var request CreateGameRequest
	if err := json.Unmarshal(r.Body, &request); err != nil {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.c.CreateNewGame(command.CreateNewGameCommand{
		ProjectGUID: chi.URLParam(r, ProjectGUID),
		Name:        request.Name,
	})
	if err == nil {
		json.WriteResponse(w, http.StatusCreated, p)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	games, err := h.q.GetGamesBy(chi.URLParam(r, ProjectGUID))

	if err == nil {
		json.WriteResponse(w, http.StatusOK, games)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}
