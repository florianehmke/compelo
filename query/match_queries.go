package query

import (
	"errors"
	"fmt"
)

var ErrMatchNotFound = errors.New("match not found")

func (svc *Service) GetMatchesBy(projectGUID string, gameGUID string) ([]*Match, error) {
	svc.RLock()
	defer svc.RUnlock()

	game, err := svc.getGameBy(projectGUID, gameGUID)
	if err != nil {
		return nil, fmt.Errorf("get matches failed: %w", err)
	}

	list := make([]*Match, 0, len(game.eloMatchList.entries))
	for _, value := range game.eloMatchList.entries {
		list = append(list, value)
	}

	sortMatchesByCreatedDate(list)
	return list, nil
}

func (svc *Service) GetMatchBy(projectGUID string, gameGUID string, matchGUID string) (*Match, error) {
	svc.RLock()
	defer svc.RUnlock()

	game, err := svc.getGameBy(projectGUID, gameGUID)
	if err != nil {
		return nil, fmt.Errorf("get matches failed: %w", err)
	}

	if match, ok := game.eloMatchList.entries[matchGUID]; ok {
		return match, nil
	} else {
		return nil, fmt.Errorf("get match by guid (%s) failed: %w", matchGUID, ErrMatchNotFound)
	}
}
