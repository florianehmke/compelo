package db

import (
	"time"
)

type Result string

const (
	Win  Result = "Win"
	Draw Result = "Draw"
	Loss Result = "Loss"
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
	Result      Result `json:"result"`
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
	err := db.gorm.Create(&match).Error
	return match, err
}

func (db *DB) CreateTeam(team Team) (Team, error) {
	err := db.gorm.Create(&team).Error
	return team, err
}

func (db *DB) CreateAppearance(appearance Appearance) (Appearance, error) {
	err := db.gorm.Create(&appearance).Error
	return appearance, err
}

func (db *DB) LoadMatchesByGameID(id uint) []Match {
	var matches []Match
	db.gorm.Where(Match{GameID: id}).Find(&matches)
	return matches
}

func (db *DB) LoadMatchByID(id uint) (Match, error) {
	var match Match
	err := db.gorm.First(&match, id).Error
	return match, err
}

func (db *DB) LoadTeamsByMatchID(id uint) ([]Team, error) {
	var teams []Team
	err := db.gorm.Where(Team{MatchID: id}).Find(&teams).Error
	return teams, err
}

func (db *DB) LoadAppearancesByMatchIDAndTeamID(matchID, teamID uint) ([]Appearance, error) {
	var players []Appearance
	err := db.gorm.Where(Appearance{MatchID: matchID, TeamID: teamID}).Find(&players).Error
	return players, err
}

func (db *DB) LoadPlayersByMatchIDAndTeamID(matchID, teamID uint) ([]Player, error) {
	var players []Player
	err := db.gorm.
		Joins("left join appearances on appearances.player_id = players.id").
		Where("appearances.match_id = ? and appearances.team_id = ? ", matchID, teamID).
		Find(&players).Error

	return players, err
}
