package db

import (
	"time"
)

// match_results view
type MatchResult struct {
	PlayerID    uint
	GameID      uint
	Date        time.Time
	RatingDelta int
	Result      Result
}

func (db *gormDB) LoadMatchResultsByGameID(gameID uint) ([]MatchResult, error) {
	var results []MatchResult
	err := db.gorm.Where(&MatchResult{GameID: gameID}).Find(&results).Error
	return results, err
}
