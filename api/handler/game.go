package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"compelo/api/json"
	"compelo/command"
	"compelo/query"
)

const (
	GameGUID string     = "gameGUID"
	GameKey  ContextKey = "game"
)

type CreateGameRequest struct {
	Name string `json:"name"`
}

func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	project := MustLoadProjectFromContext(r)
	var request CreateGameRequest
	if err := json.Unmarshal(r.Body, &request); err != nil {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.c.CreateNewGame(command.CreateNewGameCommand{
		ProjectGUID: project.GUID,
		Name:        request.Name,
	})
	if err == nil {
		json.WriteResponse(w, http.StatusCreated, p)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	project := MustLoadProjectFromContext(r)
	games, err := h.q.GetGamesBy(project.GUID)

	if err == nil {
		json.WriteResponse(w, http.StatusOK, games)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GameCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		project := MustLoadProjectFromContext(r)
		guid := chi.URLParam(r, GameGUID)
		if guid == "" {
			json.WriteErrorResponse(w, http.StatusBadRequest, errors.New("no game guid provided"))
			return
		}
		game, err := h.q.GetGameBy(project.GUID, guid)
		if err != nil {
			msg := fmt.Sprintf("could not set game with guid %s in context", guid)
			json.WriteErrorResponse(w, http.StatusNotFound, fmt.Errorf("%s: %v", msg, err))
			return
		}
		ctx := context.WithValue(r.Context(), GameKey, game)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func MustLoadGameFromContext(r *http.Request) query.Game {
	game, ok := r.Context().Value(GameKey).(query.Game)
	if !ok {
		panic("game must be set in context")
	}
	return game
}
