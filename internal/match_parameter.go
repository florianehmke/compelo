package compelo

import (
	"errors"
	"time"

	"compelo/internal/db"
)

var (
	ErrTwoTeamsRequired      = errors.New("at least two teams required")
	ErrSameTeamSizeRequired  = errors.New("all teams need the same amount of players")
	ErrPlayerInMultipleTeams = errors.New("player can only be in one team")
)

type CreateMatchParameter struct {
	GameID uint
	Date   time.Time
	Teams  []Team
}

func (p *CreateMatchParameter) validate() error {
	if len(p.Teams) < 2 {
		return ErrTwoTeamsRequired
	}
	teamSize := len(p.Teams[0].PlayerIDs)
	for _, t := range p.Teams {
		if len(t.PlayerIDs) != teamSize {
			return ErrSameTeamSizeRequired
		}
	}
	playerMap := map[int]bool{}
	for _, t := range p.Teams {
		for _, pid := range t.PlayerIDs {
			if _, ok := playerMap[pid]; ok {
				return ErrPlayerInMultipleTeams
			}
			playerMap[pid] = true
		}
	}
	return nil
}

func (param *CreateMatchParameter) determineResult() {
	highScore := 0
	highScoreCount := 0
	for _, t := range param.Teams {
		if t.Score > highScore {
			highScore = t.Score
			highScoreCount = 1
		} else if t.Score == highScore {
			highScoreCount += 1
		}
	}
	if highScoreCount < len(param.Teams) {
		for i := range param.Teams {
			if param.Teams[i].Score == highScore {
				param.Teams[i].Result = db.Win
			} else {
				param.Teams[i].Result = db.Loss
			}
		}
	} else {
		for i := range param.Teams {
			param.Teams[i].Result = db.Draw
		}
	}
}
