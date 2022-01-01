package query

type GameStats struct {
	MaxScoreSum  []*Match `json:"maxScoreSum"`
	MaxScoreDiff []*Match `json:"maxScoreDiff"`
}
