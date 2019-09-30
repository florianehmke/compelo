package compelo

import (
	"compelo/internal/db"
	"compelo/pkg/rating"
)

func (svc *Service) LoadRatingsByGameID(gameID uint) []db.Rating {
	return svc.db.LoadRatingsByGameID(gameID)
}

func (svc *Service) LoadRatingByPlayerIDAndGameID(playerID, gameID uint) (db.Rating, error) {
	r, err := svc.LoadRatingByPlayerIDAndGameID(playerID, gameID)
	if err == db.RecordNotFound {
		r = db.Rating{
			Rating:   rating.InitialRating,
			GameID:   gameID,
			PlayerID: playerID,
		}
		return svc.db.SaveRating(r)
	}
	return r, err
}

func (svc *Service) UpdateRating(playerID, gameID uint, delta int) (db.Rating, error) {
	r, err := svc.db.LoadRatingByPlayerIDAndGameID(playerID, gameID)
	if err != nil {
		return db.Rating{}, nil
	}

	r.Rating = r.Rating + delta
	return svc.db.SaveRating(r)
}
