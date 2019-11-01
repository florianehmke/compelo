package compelo

import (
	"github.com/pkg/errors"
)

type GameStats struct {
	MaxScoreSum  MatchData `json:"maxScoreSum"`
	MaxScoreDiff MatchData `json:"maxScoreDiff"`
}

func (svc *Service) LoadGameStats(gameID uint) (GameStats, error) {
	var gameStats GameStats
	var err error

	diff, err := svc.db.LoadMaxScoreDiffByGameID(gameID)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load max score diff")
	}
	gameStats.MaxScoreDiff, err = svc.LoadMatchByID(diff.MatchID)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load match data for max score diff")
	}

	sum, err := svc.db.LoadMaxScoreSumByGameID(gameID)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load max score sum")
	}
	gameStats.MaxScoreSum, err = svc.LoadMatchByID(sum.MatchID)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load match data for max score sum")
	}

	return gameStats, nil
}
