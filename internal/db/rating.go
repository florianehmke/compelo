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

	err := db.gorm.Where(Rating{
		GameID:   gameID,
		PlayerID: playerID,
	}).Attrs(Rating{
		Rating: rating.InitialRating,
	}).FirstOrCreate(&r).Error

	return r, err
}

func (db *DB) SaveRating(rating Rating) (Rating, error) {
	err := db.gorm.Save(&rating).Error
	return rating, err
}
