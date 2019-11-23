package db

import (
	"time"
)

// match_results view
type MatchResult struct {
	PlayerID    uint      `json:"playerId"`
	GameID      uint      `json:"gameId"`
	Date        time.Time `json:"date" ts_type:"string"`
	MatchID     uint      `json:"matchId"`
	Score       int       `json:"score"`
	RatingDelta int       `json:"ratingDelta"`
	Result      Result    `json:"result"`
}

func (db *gormDB) LoadMatchResultsByGameID(gameID uint) ([]MatchResult, error) {
	var results []MatchResult
	err := db.gorm.Where(&MatchResult{GameID: gameID}).Find(&results).Error
	return results, err
}
