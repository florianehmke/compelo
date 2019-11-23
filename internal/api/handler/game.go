package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"compelo/internal/db"
	"compelo/pkg/json"
)

const (
	GameID  = "gameID"
	GameKey = "game"
)

type CreateGameRequest struct {
	Name string `json:"name"`
}

func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	project := MustLoadProjectFromContext(r)
	var body CreateGameRequest
	if err := json.Unmarshal(r.Body, &body); err != nil {
		json.Error(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.svc.CreateGame(project.ID, body.Name)
	if err == nil {
		json.Write(w, http.StatusCreated, p)
	} else {
		json.Error(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	project := MustLoadProjectFromContext(r)
	games := h.svc.LoadGamesByProjectID(project.ID)
	json.Write(w, http.StatusOK, games)
}

func (h *Handler) GameCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, GameID))
		if err != nil {
			json.Error(w, http.StatusBadRequest, err)
			return
		}
		game, err := h.svc.LoadGameByID(uint(id))
		if err != nil {
			msg := fmt.Sprintf("could not set game with id %d in context", id)
			json.Error(w, http.StatusNotFound, fmt.Errorf("%s: %v", msg, err))
			return
		}
		ctx := context.WithValue(r.Context(), GameKey, game)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func MustLoadGameFromContext(r *http.Request) db.Game {
	game, ok := r.Context().Value(GameKey).(db.Game)
	if !ok {
		panic("game must be set in context")
	}
	return game
}
