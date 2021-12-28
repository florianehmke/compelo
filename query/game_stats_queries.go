package query

import (
	"fmt"
)

func (svc *Service) GetGameStatsBy(projectGUID string, gameGUID string) (*GameStats, error) {
	svc.RLock()
	defer svc.RUnlock()

	game, err := svc.getGameBy(projectGUID, gameGUID)
	if err != nil {
		return nil, fmt.Errorf("get matches failed: %w", err)
	}

	return game.gameStats, nil
}
