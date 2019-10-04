package compelo

import (
	"sort"

	"github.com/pkg/errors"

	"compelo/internal/db"
	"compelo/pkg/rating"
)

type PlayerStats struct {
	db.Player        // embedded player
	Rating       int `json:"rating"`
	PeakRating   int `json:"peakRating"`
	LowestRating int `json:"lowestRating"`
	GameCount    int `json:"gameCount"`
	WinCount     int `json:"winCount"`
	DrawCount    int `json:"drawCount"`
	LossCount    int `json:"lossCount"`
}

func (svc *Service) LoadPlayerStatsByGameID(gameID uint) ([]PlayerStats, error) {
	ratings := svc.LoadRatingsByGameID(gameID)

	var players []PlayerStats
	for _, r := range ratings {
		p, err := svc.LoadPlayerByID(r.PlayerID)
		if err != nil {
			return nil, errors.Wrap(err, "could not load player by rating")
		}

		pws := PlayerStats{
			Player:       p,
			Rating:       r.Rating,
			PeakRating:   rating.InitialRating,
			LowestRating: rating.InitialRating,
		}
		results, err := svc.db.LoadMatchResultsByPlayerIDAndGameID(p.ID, gameID)
		if err != nil {
			return nil, errors.Wrap(err, "could not match results")
		}

		pws.applyRatingStats(results)
		pws.applyResultStats(results)

		players = append(players, pws)
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Rating > (players[j].Rating)
	})
	return players, nil
}

func (p *PlayerStats) applyRatingStats(results []db.MatchResult) {
	current := rating.InitialRating
	for _, result := range results {
		current = current + result.RatingDelta
		if current > p.PeakRating {
			p.PeakRating = current
		}
		if current < p.LowestRating {
			p.LowestRating = current
		}
	}
}

func (p *PlayerStats) applyResultStats(results []db.MatchResult) {
	for _, r := range results {
		switch r.Result {
		case db.Win:
			p.WinCount += 1
		case db.Draw:
			p.DrawCount += 1
		case db.Loss:
			p.LossCount += 1
		}
	}
	p.GameCount = len(results)
}
