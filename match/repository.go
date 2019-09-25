package match

import (
	"time"

	"compelo/db"
)

type Match struct {
	db.Model

	Date   time.Time `json:"date"`
	GameID uint      `json:"gameId"`
}

type Team struct {
	db.Model

	MatchID     uint   `json:"matchId"`
	Score       int    `json:"score"`
	Result      string `json:"result"`
	RatingDelta int    `json:"ratingDelta"`
}

type Appearance struct {
	db.Model

	MatchID     uint `json:"matchId"`
	TeamID      uint `json:"teamId"`
	PlayerID    uint `json:"playerId"`
	RatingDelta int  `json:"ratingDelta"`
}

type Repository interface {
	create(createMatchParameter) (Match, error)
	loadByGameID(uint) ([]Match, error)
	loadByID(uint) (Match, error)

	loadTeamsByMatchID(uint) ([]Team, error)

	loadAppearancesByMatchIDAndTeamID(matchID, teamID uint) ([]Appearance, error)
}

type result int

const (
	Loss result = iota
	Draw
	Win
)

func (r result) String() string {
	return []string{"Loss", "Draw", "Win"}[r]
}

var _ Repository = repository{}

type repository struct {
	db *db.DB
}

func (r repository) create(param createMatchParameter) (Match, error) {
	tx := r.db.Begin()

	// 1. Create match.
	match := Match{GameID: param.gameID, Date: param.date}
	tx.Create(&match)

	// 2. Create teams.
	for _, team := range param.Teams {
		t := Team{
			MatchID:     match.ID,
			Score:       team.Score,
			Result:      team.result,
			RatingDelta: team.ratingDelta,
		}
		tx.Create(&t)

		// 3. Create appearances for players.
		for _, playerID := range team.PlayerIDs {
			c := Appearance{
				MatchID:     match.ID,
				TeamID:      t.ID,
				PlayerID:    uint(playerID),
				RatingDelta: team.ratingDelta,
			}
			tx.Create(&c)
		}
	}

	return match, tx.Commit().Error
}

func (r repository) loadByGameID(id uint) ([]Match, error) {
	var matches []Match
	err := r.db.Where(Match{GameID: id}).Find(&matches).Error
	return matches, err
}

func (r repository) loadByID(id uint) (Match, error) {
	var match Match
	err := r.db.First(&match, id).Error
	return match, err
}

func (r repository) loadTeamsByMatchID(id uint) ([]Team, error) {
	var teams []Team
	err := r.db.Where(Team{MatchID: id}).Find(&teams).Error
	return teams, err
}

func (r repository) loadAppearancesByMatchIDAndTeamID(matchID, teamID uint) ([]Appearance, error) {
	var players []Appearance
	err := r.db.Where(Appearance{MatchID: matchID, TeamID: teamID}).Find(&players).Error
	return players, err
}
