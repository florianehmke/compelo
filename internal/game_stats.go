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

	mss, err := svc.db.LoadMaxScoreDiffByGameID(gameID)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load max score diff")
	}
	msd, err := svc.db.LoadMaxScoreSumByGameID(gameID)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load max score sum")
	}

	gameStats.MaxScoreSum, err = svc.LoadMatchByID(mss.MatchID)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load match data for max score sum")
	}
	gameStats.MaxScoreDiff, err = svc.LoadMatchByID(msd.MatchID)
	if err != nil {
		return GameStats{}, errors.Wrap(err, "could not load match data for max score sum")
	}

	return gameStats, nil
}
