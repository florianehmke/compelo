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
	results, err := svc.db.LoadMatchResultsByGameID(gameID)
	if err != nil {
		return nil, errors.Wrapf(err, "could not load match results by game id %d", gameID)
	}

	statsMap := map[uint]*PlayerStats{}
	for _, result := range results {
		playerStats, exists := statsMap[result.PlayerID]

		// a new player appeared!
		if !exists {
			newPlayerStats, err := svc.newPlayerStats(result.PlayerID)
			if err != nil {
				return nil, errors.Wrap(err, "could not create initial player stats")
			}
			statsMap[result.PlayerID] = &newPlayerStats
			playerStats = &newPlayerStats
		}

		// add results of match to stats
		playerStats.addResult(result)

		// for all other players, copy their current stats into history for that day
		for _, otherPlayerStats := range statsMap {
			if otherPlayerStats.Player.ID != result.PlayerID {
				otherPlayerStats.copyResultToDate(result.Date)
			}
		}
	}

	// return a sorted slice
	var playerStats []PlayerStats
	for _, ps := range statsMap {
		playerStats = append(playerStats, *ps)
	}
	sort.Slice(playerStats, func(i, j int) bool {
		return playerStats[i].Current.Rating > (playerStats[j].Current.Rating)
	})

	return playerStats, nil
}

func (svc *Service) newPlayerStats(playerID uint) (PlayerStats, error) {
	player, err := svc.LoadPlayerByID(playerID)
	if err != nil {
		return PlayerStats{}, errors.Wrapf(err, "could not load player by id %d", playerID)
	}

	return PlayerStats{
		Player: player,
		Current: Stats{
			Rating:       rating.InitialRating,
			PeakRating:   rating.InitialRating,
			LowestRating: rating.InitialRating,
		},
		History: map[string]Stats{},
	}, nil
}

func (p *PlayerStats) addResult(result db.MatchResult) {
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
	p.Current.GameCount += 1
	p.History[formatDate(result.Date)] = p.Current
}

func (p *PlayerStats) copyResultToDate(date time.Time) {
	p.History[formatDate(date)] = p.Current
}

func formatDate(date time.Time) string {
	withoutTime := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	return withoutTime.Format("2006-01-02")
}
