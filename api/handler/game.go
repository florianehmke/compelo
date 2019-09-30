package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"compelo/internal/db"
)

const (
	GameID  = "gameID"
	GameKey = "game"
)

func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	project := mustLoadProjectFromContext(r)
	var body struct {
		Name string `json:"name"`
	}
	if err := unmarshal(r.Body, &body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.svc.CreateGame(project.ID, body.Name)
	if err == nil {
		writeJSON(w, http.StatusCreated, p)
	} else {
		writeError(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	project := mustLoadProjectFromContext(r)
	games := h.svc.LoadGamesByProjectID(project.ID)
	writeJSON(w, http.StatusOK, games)
}

func (h *Handler) GameCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, GameID))
		if err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}
		game, err := h.svc.LoadGameByID(uint(id))
		if err != nil {
			msg := fmt.Sprintf("could not set game with id %d in context", id)
			writeError(w, http.StatusNotFound, fmt.Errorf("%s: %v", msg, err))
			return
		}
		ctx := context.WithValue(r.Context(), GameKey, game)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func mustLoadGameFromContext(r *http.Request) db.Game {
	game, ok := r.Context().Value(GameKey).(db.Game)
	if !ok {
		panic("game must be set in context")
	}
	return game
}
