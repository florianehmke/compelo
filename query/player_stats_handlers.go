package query

import (
	"compelo/event"
	"compelo/rating"
	"log"
)

type playerStatsHandler struct {
	data *data
}

func (h *playerStatsHandler) handleMatchCreated(e *event.MatchCreated) {
	log.Println("[query:player-stats] handling event", e.GetID(), e.EventType())

	game := h.data.projects[e.ProjectGUID].games[e.GameGUID]
	match := game.eloMatchList.entries[e.GUID]
	for _, t := range match.Teams {
		for _, p := range t.Players {
			stats, ok := game.playerStats[p.GUID]
			if !ok {
				stats = h.newPlayerStats(p)
			}
			stats.addResult(match, t)

			// FIXME: very inefficient, instead chart it better.
			// Results are just copied to provide a data point
			// for every player on every match day, even if the
			// player did not play that day.
			for _, otherPlayerStats := range game.playerStats {
				if otherPlayerStats.Player.GUID != p.GUID {
					otherPlayerStats.copyCurrentResultToHistory(e.Date)
				}
			}

			game.playerStats[p.GUID] = stats
		}
	}
}

func (h *playerStatsHandler) newPlayerStats(player *Player) *PlayerStats {
	return &PlayerStats{
		Player: player,
		Current: Stats{
			Rating:       rating.InitialRating,
			PeakRating:   rating.InitialRating,
			LowestRating: rating.InitialRating,
		},
		History: map[string]Stats{},
	}
}
