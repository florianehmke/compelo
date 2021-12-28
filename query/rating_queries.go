package query

import (
	"errors"
	"fmt"
)

var ErrRatingNotFound = errors.New("rating not found")

func (svc *Service) GetRatingBy(projectGUID string, playerGUID string, gameGUID string) (*Rating, error) {
	svc.RLock()
	defer svc.RUnlock()

	return svc.getRatingBy(projectGUID, playerGUID, gameGUID)
}

func (svc *Service) getRatingBy(projectGUID string, playerGUID string, gameGUID string) (*Rating, error) {
	player, err := svc.getPlayerBy(projectGUID, playerGUID)
	if err != nil {
		return nil, fmt.Errorf("get rating failed: %w", err)
	}

	if r, ok := player.ratings[gameGUID]; ok {
		return r, nil
	} else {
		return nil, fmt.Errorf("get rating by player guid (%s) failed: %w", playerGUID, ErrRatingNotFound)
	}
}
