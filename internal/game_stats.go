package compelo

import (
	"github.com/pkg/errors"
)

type GameStats struct {
	MaxScoreSum  []MatchData `json:"maxScoreSum"`
	MaxScoreDiff []MatchData `json:"maxScoreDiff"`
}

func (svc *Service) LoadGameStats(gameID uint) (GameStats, error) {
	var gameStats GameStats
	var err error

	stats, err := svc.db.LoadMaxScoreDiffByGameID(gameID, 5)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load max score diff")
	}
	for _, s := range stats {
		matchData, err := svc.LoadMatchByID(s.MatchID)
		if err != nil {
			return GameStats{}, errors.Wrap(err, "could not load match data for max score diff")
		}
		gameStats.MaxScoreDiff = append(gameStats.MaxScoreDiff, matchData)
	}

	stats, err = svc.db.LoadMaxScoreSumByGameID(gameID, 5)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load max score sum")
	}
	for _, s := range stats {
		matchData, err := svc.LoadMatchByID(s.MatchID)
		if err != nil {
			return GameStats{}, errors.Wrap(err, "could not load match data for max score sum")
		}
		gameStats.MaxScoreSum = append(gameStats.MaxScoreSum, matchData)

	}

	return gameStats, nil
}
