package query

import (
	"compelo/event"
	"log"
	"sort"
)

type gameStatsHandler struct {
	c *Compelo
}

func (h *gameStatsHandler) on(e event.Event) {
	switch e := e.(type) {
	case *event.MatchCreated:
		h.handleMatchCreated(e)
	}
}

func (h *gameStatsHandler) handleMatchCreated(e *event.MatchCreated) {
	log.Println("[query:game-stats] handling event", e.GetID(), e.EventType())

	game := h.c.projects[e.ProjectGUID].games[e.GameGUID]
	match := game.matches[e.GUID]

	game.gameStats.MaxScoreDiff = append(game.gameStats.MaxScoreDiff, match)
	sort.Slice(game.gameStats.MaxScoreDiff, func(i, j int) bool {
		return game.gameStats.MaxScoreDiff[i].scoreDifference() > (game.gameStats.MaxScoreDiff[j].scoreDifference())
	})
	if len(game.gameStats.MaxScoreDiff) > 3 {
		game.gameStats.MaxScoreDiff = game.gameStats.MaxScoreDiff[:len(game.gameStats.MaxScoreDiff)-1]
	}

	game.gameStats.MaxScoreSum = append(game.gameStats.MaxScoreSum, match)
	sort.Slice(game.gameStats.MaxScoreSum, func(i, j int) bool {
		return game.gameStats.MaxScoreSum[i].scoreSum() > (game.gameStats.MaxScoreSum[j].scoreSum())
	})

	if len(game.gameStats.MaxScoreSum) > 3 {
		game.gameStats.MaxScoreSum = game.gameStats.MaxScoreSum[:len(game.gameStats.MaxScoreSum)-1]
	}
}
