package query

import (
	"compelo/event"
	"log"
	"sort"
)

type gameStatsHandler struct {
	data *data
}

const matchesPerStats = 5

func (h *gameStatsHandler) handleMatchCreated(e *event.MatchCreated) {
	log.Println("[query:game-stats] handling event", e.GetID(), e.EventType())

	game := h.data.projects[e.ProjectGUID].games[e.GameGUID]
	match := game.matches[e.GUID]

	updateMaxScoreSumStats(game.gameStats, match)
	updateMaxScoreDiffStats(game.gameStats, match)
}

func updateMaxScoreSumStats(stats *GameStats, match *Match) {
	stats.MaxScoreSum = append(stats.MaxScoreSum, match)

	sort.Slice(stats.MaxScoreSum, func(i, j int) bool {
		return stats.MaxScoreSum[i].scoreSum() > (stats.MaxScoreSum[j].scoreSum())
	})

	if len(stats.MaxScoreSum) > matchesPerStats {
		stats.MaxScoreSum = stats.MaxScoreSum[:len(stats.MaxScoreSum)-1]
	}
}

func updateMaxScoreDiffStats(stats *GameStats, match *Match) {
	stats.MaxScoreDiff = append(stats.MaxScoreDiff, match)

	sort.Slice(stats.MaxScoreDiff, func(i, j int) bool {
		return stats.MaxScoreDiff[i].scoreDifference() > (stats.MaxScoreDiff[j].scoreDifference())
	})

	if len(stats.MaxScoreDiff) > matchesPerStats {
		stats.MaxScoreDiff = stats.MaxScoreDiff[:len(stats.MaxScoreDiff)-1]
	}
}
