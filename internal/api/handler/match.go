package handler

import (
	"net/http"
	"time"

	compelo "compelo/internal"
	"compelo/internal/db"
	"compelo/pkg/json"
)

type CreateMatchRequest struct {
	Teams []struct {
		PlayerIDs []int `json:"playerIds" `
		Score     int   `json:"score"`
	} `json:"teams"`
}

func (h *Handler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	game := MustLoadGameFromContext(r)

	var body CreateMatchRequest
	if err := json.Unmarshal(r.Body, &body); err != nil {
		json.Error(w, http.StatusBadRequest, err)
		return
	}

	m, err := h.svc.CreateMatch(createMatchParameter(game, body))
	if err == nil {
		json.Write(w, http.StatusCreated, m)
	} else {
		json.Error(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) GetAllMatches(w http.ResponseWriter, r *http.Request) {
	game := MustLoadGameFromContext(r)

	matches, err := h.svc.LoadMatchesByGameID(game.ID)
	if err == nil {
		json.Write(w, http.StatusOK, matches)
	} else {
		json.Error(w, http.StatusBadRequest, err)
	}
}

func createMatchParameter(game db.Game, body CreateMatchRequest) compelo.CreateMatchParameter {
	param := compelo.CreateMatchParameter{
		GameID: game.ID,
		Date:   time.Now(),
	}
	for _, t := range body.Teams {
		param.Teams = append(param.Teams, compelo.CreateMatchParameterTeam{
			PlayerIDs: t.PlayerIDs,
			Score:     t.Score,
		})
	}
	return param
}
