package compelo

import (
	"compelo/internal/db"
)

func (svc *Service) LoadRatingsByGameID(gameID uint) []db.Rating {
	return svc.db.LoadRatingsByGameID(gameID)
}

func (svc *Service) LoadOrCreateRatingByPlayerIDAndGameID(playerID, gameID uint) (db.Rating, error) {
	return svc.db.LoadOrCreateRatingByPlayerIDAndGameID(playerID, gameID)
}

func (svc *Service) UpdateRating(playerID, gameID uint, delta int) (db.Rating, error) {
	r, err := svc.db.LoadOrCreateRatingByPlayerIDAndGameID(playerID, gameID)
	if err != nil {
		return db.Rating{}, nil
	}

	r.Rating = r.Rating + delta
	return svc.db.SaveRating(r)
}
