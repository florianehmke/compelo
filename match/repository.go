package match

import (
	"compelo"
	"compelo/db"
)

type Repository interface {
	Create(
		compelo.Match,
		map[int]compelo.MatchTeam,
		map[int][]compelo.MatchPlayer,
	) (compelo.Match, error)

	LoadByGameID(uint) ([]compelo.Match, error)
	LoadByID(uint) (compelo.Match, error)

	LoadTeamsByMatchID(uint) ([]compelo.MatchTeam, error)
	LoadPlayersByMatchIDAndTeamID(matchID, teamID uint) ([]compelo.MatchPlayer, error)
}

var _ Repository = repository{}

type repository struct {
	db *db.DB
}

func (r repository) Create(
	match compelo.Match,
	teamMap map[int]compelo.MatchTeam,
	playerMap map[int][]compelo.MatchPlayer,
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
			p.MatchTeamID = team.ID
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

func (r repository) LoadTeamsByMatchID(id uint) ([]compelo.MatchTeam, error) {
	var teams []compelo.MatchTeam
	err := r.db.Where(compelo.MatchTeam{MatchID: id}).Find(&teams).Error
	return teams, err
}

func (r repository) LoadPlayersByMatchIDAndTeamID(matchID, teamID uint) ([]compelo.MatchPlayer, error) {
	var players []compelo.MatchPlayer
	err := r.db.Where(compelo.MatchPlayer{MatchID: matchID, MatchTeamID: teamID}).Find(&players).Error
	return players, err
}
