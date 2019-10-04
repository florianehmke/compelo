package compelo

import (
	"sort"
	"time"

	"github.com/pkg/errors"

	"compelo/internal/db"
	"compelo/pkg/rating"
)

type PlayerStats struct {
	db.Player                  // embedded player
	Current   Stats            `json:"current"`
	History   map[string]Stats `json:"history"` // mapped by RFC3339 formatted date (eg. 2019-10-01T00:00:00Z)
}

type Stats struct {
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
			Player: p,
			Current: Stats{
				Rating:       r.Rating,
				PeakRating:   rating.InitialRating,
				LowestRating: rating.InitialRating,
			},
			History: make(map[string]Stats),
		}
		results, err := svc.db.LoadMatchResultsByPlayerIDAndGameID(p.ID, gameID)
		if err != nil {
			return nil, errors.Wrap(err, "could not load match results")
		}

		pws.applyStats(results)
		players = append(players, pws)
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Current.Rating > (players[j].Current.Rating)
	})
	return players, nil
}

func (p *PlayerStats) applyStats(results []db.MatchResult) {
	p.Current.Rating = rating.InitialRating

	for i, result := range results {
		// rating (peak, low, current)
		p.Current.Rating = p.Current.Rating + result.RatingDelta
		if p.Current.Rating > p.Current.PeakRating {
			p.Current.PeakRating = p.Current.Rating
		}
		if p.Current.Rating < p.Current.LowestRating {
			p.Current.LowestRating = p.Current.Rating
		}

		// game results
		switch result.Result {
		case db.Win:
			p.Current.WinCount += 1
		case db.Draw:
			p.Current.DrawCount += 1
		case db.Loss:
			p.Current.LossCount += 1
		}
		p.Current.GameCount = i + 1

		// store copy in history map
		p.History[beginningOfDay(result.Date).Format(time.RFC3339)] = p.Current
	}
}

func beginningOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}
