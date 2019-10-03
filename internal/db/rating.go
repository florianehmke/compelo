package db

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

func (db *DB) LoadRatingByPlayerIDAndGameID(playerID, gameID uint) (Rating, error) {
	var rating Rating
	ref := db.gorm.Where(Rating{GameID: gameID, PlayerID: playerID}).First(&rating)
	if ref.RecordNotFound() {
		return Rating{}, RecordNotFound
	}
	return rating, ref.Error
}

func (db *DB) SaveRating(rating Rating) (Rating, error) {
	err := db.gorm.Save(&rating).Error
	return rating, err
}
