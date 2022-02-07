package query

import (
	"errors"
	"fmt"
)

var ErrPlayerNotFound = errors.New("player not found")

func (svc *Service) GetPlayersBy(projectGUID string) ([]*Player, error) {
	svc.RLock()
	defer svc.RUnlock()

	project, err := svc.getProjectBy(projectGUID)
	if err != nil {
		return nil, fmt.Errorf("get players failed: %w", err)
	}

	list := make([]*Player, 0, len(project.players))
	for _, value := range project.players {
		list = append(list, value)
	}

	sortPlayersByCreatedDate(list)
	return list, nil
}

func (svc *Service) GetPlayerBy(projectGUID string, playerGUID string) (*Player, error) {
	svc.RLock()
	defer svc.RUnlock()

	return svc.getPlayerBy(projectGUID, playerGUID)
}

func (svc *Service) getPlayerBy(projectGUID string, playerGUID string) (*Player, error) {
	project, err := svc.getProjectBy(projectGUID)
	if err != nil {
		return nil, fmt.Errorf("get player failed: %w", err)
	}

	if player, ok := project.players[playerGUID]; ok {
		return player, nil
	}
	return nil, fmt.Errorf("get player by guid (%s) failed: %w", playerGUID, ErrPlayerNotFound)
}
