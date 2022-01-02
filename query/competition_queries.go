package query

import (
	"fmt"
)

func (svc *Service) GetCompetitionsBy(projectGUID string, gameGUID string) ([]*Competition, error) {
	svc.RLock()
	defer svc.RUnlock()

	game, err := svc.getGameBy(projectGUID, gameGUID)
	if err != nil {
		return nil, fmt.Errorf("get competitions failed: %w", err)
	}

	list := make([]*Competition, 0, len(game.competitions))
	for _, value := range game.competitions {
		list = append(list, value)
	}

	sortCompetitionsByCreatedDate(list)
	return list, nil
}
