package query

import (
	"fmt"
	"sort"
)

func (svc *Service) GetPlayerStatsBy(projectGUID string, gameGUID string) ([]*PlayerStats, error) {
	svc.RLock()
	defer svc.RUnlock()

	game, err := svc.getGameBy(projectGUID, gameGUID)
	if err != nil {
		return nil, fmt.Errorf("get matches failed: %w", err)
	}

	list := make([]*PlayerStats, 0, len(game.playerStats))
	for _, value := range game.playerStats {
		list = append(list, value)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Current.Rating > (list[j].Current.Rating)
	})
	return list, nil
}
