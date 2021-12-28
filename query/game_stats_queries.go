package query

import (
	"fmt"
)

func (c *Compelo) GetGameStatsBy(projectGUID string, gameGUID string) (*GameStats, error) {
	c.RLock()
	defer c.RUnlock()

	game, err := c.getGameBy(projectGUID, gameGUID)
	if err != nil {
		return nil, fmt.Errorf("get matches failed: %w", err)
	}

	return game.gameStats, nil
}
