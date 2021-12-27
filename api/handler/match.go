package handler

import (
	"compelo/api/json"
	"compelo/command"
	"net/http"
)

type CreateMatchRequest struct {
	Teams []CreateMatchRequestTeam `json:"teams"`
}

type CreateMatchRequestTeam struct {
	PlayerGUIDs []string `json:"playerGuids" `
	Score       int      `json:"score"`
}

func (h *Handler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	game := MustLoadGameFromContext(r)

	var request CreateMatchRequest
	if err := json.Unmarshal(r.Body, &request); err != nil {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	c := command.CreateNewMatchCommand{
		GameGUID:    game.GUID,
		ProjectGUID: game.ProjectGUID,
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

func (h *Handler) GetAllMatches(w http.ResponseWriter, r *http.Request) {
	game := MustLoadGameFromContext(r)

	matches, err := h.q.GetMatchesBy(game.ProjectGUID, game.GUID)
	if err == nil {
		json.WriteResponse(w, http.StatusOK, matches)
	} else {
		json.WriteErrorResponse(w, http.StatusBadRequest, err)
	}
}
