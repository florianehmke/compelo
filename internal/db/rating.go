package db

import (
	"compelo/pkg/rating"
)

type Rating struct {
	Model

	Rating   int  `json:"rating"`
	GameID   uint `json:"gameId"`
	PlayerID uint `json:"playerId"`
}

func (db *DB) LoadRatingsByGameID(gameID uint) []Rating {
	var ratings []Rating
	db.gorm.Where(&Rating{GameID: gameID}).Find(&ratings)
	return ratings
}

func (db *DB) LoadOrCreateRatingByPlayerIDAndGameID(playerID, gameID uint) (Rating, error) {
	var r Rating

	where := Rating{GameID: gameID, PlayerID: playerID}
	attrs := Rating{Rating: rating.InitialRating}

	err := db.gorm.Where(where).Attrs(attrs).FirstOrCreate(&r).Error

	return r, err
}

func (db *DB) SaveRating(rating Rating) (Rating, error) {
	err := db.gorm.Save(&rating).Error
	return rating, err
}
