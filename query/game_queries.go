package query

import (
	"errors"
	"fmt"
)

var ErrGameNotFound = errors.New("game not found")

func (svc *Service) GetGamesBy(projectGUID string) ([]*Game, error) {
	svc.RLock()
	defer svc.RUnlock()

	project, err := svc.getProjectBy(projectGUID)
	if err != nil {
		return nil, fmt.Errorf("get games failed: %w", err)
	}

	list := make([]*Game, 0, len(project.games))
	for _, value := range svc.data.projects[projectGUID].games {
		list = append(list, value)
	}

	sortGamesByCreatedDate(list)
	return list, nil
}

func (svc *Service) GetGameBy(projectGUID string, gameGUID string) (*Game, error) {
	svc.RLock()
	defer svc.RUnlock()

	return svc.getGameBy(projectGUID, gameGUID)
}

func (svc *Service) getGameBy(projectGUID string, gameGUID string) (*Game, error) {
	project, err := svc.getProjectBy(projectGUID)
	if err != nil {
		return nil, fmt.Errorf("get game failed: %w", err)
	}

	if game, ok := project.games[gameGUID]; ok {
		return game, nil
	} else {
		return nil, fmt.Errorf("get game by guid (%s) failed: %w", gameGUID, ErrGameNotFound)
	}
}
