package compelo

import (
	"github.com/pkg/errors"

	"compelo/internal/db"
)

func (svc *Service) LoadOrCreateRatingByPlayerIDAndGameID(playerID, gameID uint) (db.Rating, error) {
	return svc.db.LoadOrCreateRatingByPlayerIDAndGameID(playerID, gameID)
}

func (svc *Service) UpdateRating(playerID, gameID uint, delta int) (db.Rating, error) {
	r, err := svc.db.LoadOrCreateRatingByPlayerIDAndGameID(playerID, gameID)
	if err != nil {
		return db.Rating{}, errors.Wrap(err, "load/create player rating failed")
	}

	r.Rating = r.Rating + delta
	return svc.db.SaveRating(r)
}
