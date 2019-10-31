package db

// match_score_stats view
type MatchScoreStats struct {
	MatchID   uint
	GameID    uint
	ScoreDiff int
	ScoreSum  int
}

func (db *gormDB) LoadMaxScoreDiffByGameID(gameID uint) (MatchScoreStats, error) {
	var result MatchScoreStats
	err := db.gorm.
		Where(&MatchScoreStats{GameID: gameID}).
		Order("score_diff desc").
		First(&result).Error
	return result, err
}

func (db *gormDB) LoadMaxScoreSumByGameID(gameID uint) (MatchScoreStats, error) {
	var result MatchScoreStats
	err := db.gorm.
		Where(&MatchScoreStats{GameID: gameID}).
		Order("score_sum desc").
		First(&result).Error
	return result, err
}
