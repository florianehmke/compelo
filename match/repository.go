package match

import (
	"compelo"
	"compelo/db"
)

type Repository interface {
	Create(
		compelo.Match,
		map[int]compelo.Team,
		map[int][]compelo.Appearance,
	) (compelo.Match, error)

	LoadByGameID(uint) ([]compelo.Match, error)
	LoadByID(uint) (compelo.Match, error)

	LoadTeamsByMatchID(uint) ([]compelo.Team, error)
	LoadAppearancesByMatchIDAndTeamID(matchID, teamID uint) ([]compelo.Appearance, error)
}

var _ Repository = repository{}

type repository struct {
	db *db.DB
}

func (r repository) Create(
	match compelo.Match,
	teamMap map[int]compelo.Team,
	playerMap map[int][]compelo.Appearance,
) (compelo.Match, error) {
	tx := r.db.Begin()
	tx.Create(&match)

	for i, team := range teamMap {
		team.MatchID = match.ID
		if err := tx.Create(&team).Error; err != nil {
			return match, err
		}

		for _, p := range playerMap[i] {
			p.MatchID = match.ID
			p.TeamID = team.ID
			if err := tx.Create(&p).Error; err != nil {
				return match, err
			}
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
