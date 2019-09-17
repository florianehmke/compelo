package match

import (
	"compelo"
	"compelo/db"
)

type Repository interface {
	Create(CreateMatchParameter) (compelo.Match, error)

	LoadByGameID(uint) ([]compelo.Match, error)
	LoadByID(uint) (compelo.Match, error)

	LoadTeamsByMatchID(uint) ([]compelo.Team, error)
	LoadAppearancesByMatchIDAndTeamID(matchID, teamID uint) ([]compelo.Appearance, error)
}

var _ Repository = repository{}

type repository struct {
	db *db.DB
}

func (r repository) Create(param CreateMatchParameter) (compelo.Match, error) {
	tx := r.db.Begin()

	// 1. Create match.
	match := compelo.Match{GameID: param.GameID, Date: param.Date}
	tx.Create(&match)

	// 2. Create teams.
	for _, team := range param.Teams {
		t := compelo.Team{
			MatchID: match.ID,
			Score:   team.Score,
			Winner:  team.Winner,
		}
		tx.Create(&t)

		// 3. Create appearances for players.
		for _, playerID := range team.PlayerIDs {
			c := compelo.Appearance{
				MatchID:  match.ID,
				TeamID:   t.ID,
				PlayerID: uint(playerID),
			}
			tx.Create(&c)
		}
	}

	return match, tx.Commit().Error
}

func (r repository) LoadByGameID(id uint) ([]compelo.Match, error) {
	var matches []compelo.Match
	err := r.db.Where(compelo.Match{GameID: id}).Find(&matches).Error
	return matches, err
}

func (r repository) LoadByID(id uint) (compelo.Match, error) {
	var match compelo.Match
	err := r.db.First(&match, id).Error
	return match, err
}

func (r repository) LoadTeamsByMatchID(id uint) ([]compelo.Team, error) {
	var teams []compelo.Team
	err := r.db.Where(compelo.Team{MatchID: id}).Find(&teams).Error
	return teams, err
}

func (r repository) LoadAppearancesByMatchIDAndTeamID(matchID, teamID uint) ([]compelo.Appearance, error) {
	var players []compelo.Appearance
	err := r.db.Where(compelo.Appearance{MatchID: matchID, TeamID: teamID}).Find(&players).Error
	return players, err
}
