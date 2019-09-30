package db

import (
	"time"
)

const (
	Win  = "Win"
	Draw = "Draw"
	Loss = "Loss"
)

type Match struct {
	Model

	Date   time.Time `json:"date"`
	GameID uint      `json:"gameId"`
}

type Team struct {
	Model

	MatchID     uint   `json:"matchId"`
	Score       int    `json:"score"`
	Result      string `json:"result"`
	RatingDelta int    `json:"ratingDelta"`
}

type Appearance struct {
	Model

	MatchID     uint `json:"matchId"`
	TeamID      uint `json:"teamId"`
	PlayerID    uint `json:"playerId"`
	RatingDelta int  `json:"ratingDelta"`
}

func (db *DB) CreateMatch(match Match) (Match, error) {
	err := db.Create(&match).Error
	return match, err
}

func (db *DB) CreateTeam(team Team) (Team, error) {
	err := db.Create(&team).Error
	return team, err
}

func (db *DB) CreateAppearance(appearance Appearance) (Appearance, error) {
	err := db.Create(&appearance).Error
	return appearance, err
}

func (db *DB) LoadMatchesByGameID(id uint) []Match {
	var matches []Match
	db.Where(Match{GameID: id}).Find(&matches)
	return matches
}

func (db *DB) LoadMatchByID(id uint) (Match, error) {
	var match Match
	err := db.First(&match, id).Error
	return match, err
}

func (db *DB) LoadTeamsByMatchID(id uint) ([]Team, error) {
	var teams []Team
	err := db.Where(Team{MatchID: id}).Find(&teams).Error
	return teams, err
}

func (db *DB) LoadAppearancesByMatchIDAndTeamID(matchID, teamID uint) ([]Appearance, error) {
	var players []Appearance
	err := db.Where(Appearance{MatchID: matchID, TeamID: teamID}).Find(&players).Error
	return players, err
}
