package match

import (
	"sort"
	"time"

	"compelo"
	"compelo/db"
	"compelo/game"
	"compelo/player"
)

type Service struct {
	repository Repository

	playerService *player.Service
	gameService   *game.Service
}

func NewService(
	db *db.DB,
	playerService *player.Service,
	gameService *game.Service,
) *Service {
	return &Service{
		repository:    repository{db},
		playerService: playerService,
		gameService:   gameService,
	}
}

type CreateMatchParameter struct {
	Date   time.Time
	GameID uint

	Teams []struct {
		PlayerIDs []int `json:"playerIds" binding:"required"`
		Score     int   `json:"score" binding:"required"`
		Winner    bool  `json:"winner" binding:"required"`
	} `json:"teams" binding:"required"`
}

// TODO wrap in txn
func (s *Service) CreateMatch(param CreateMatchParameter) (compelo.Match, error) {
	m := compelo.Match{GameID: param.GameID, Date: time.Now()}

	teamMap := map[int]compelo.MatchTeam{}
	playerMap := map[int][]compelo.MatchPlayer{}

	for i, t := range param.Teams {
		teamMap[i] = compelo.MatchTeam{
			Score:  t.Score,
			Winner: t.Winner,
		}

		for _, pid := range t.PlayerIDs {
			playerMap[i] = append(playerMap[i], compelo.MatchPlayer{
				PlayerID: uint(pid),
			})

		}
	}
	return s.repository.Create(m, teamMap, playerMap)
}

type Match struct {
	compelo.Match
	Teams []Team `json:"teams"`
}

type Team struct {
	compelo.MatchTeam
	Players []compelo.Player `json:"players"`
}

func (s *Service) LoadByGameID(gameID uint) ([]Match, error) {
	var matches []Match

	ms, err := s.repository.LoadByGameID(gameID)
	for _, m := range ms {
		match, err := s.LoadByID(m.ID)
		if err != nil {
			return matches, err
		}
		matches = append(matches, match)
	}

	return matches, err
}

func (s *Service) LoadByID(id uint) (Match, error) {
	var match = Match{}
	var err error

	if match.Match, err = s.repository.LoadByID(id); err != nil {
		return match, err
	}

	var teams []compelo.MatchTeam
	if teams, err = s.repository.LoadTeamsByMatchID(id); err != nil {
		return match, err
	}

	for _, t := range teams {
		players, err := s.LoadPlayersByMatchIDAndTeamID(id, t.ID)
		if err != nil {
			return match, err
		}

		match.Teams = append(match.Teams, Team{
			MatchTeam: t,
			Players:   players,
		})
	}

	sort.Slice(match.Teams, func(i, j int) bool {
		if match.Teams[i].Score == match.Teams[j].Score {
			return match.Teams[i].ID < match.Teams[j].ID
		}
		return match.Teams[i].Score > match.Teams[j].Score
	})
	return match, err
}

func (s *Service) LoadPlayersByMatchIDAndTeamID(matchID, teamID uint) ([]compelo.Player, error) {
	var err error
	var players []compelo.Player

	var mps []compelo.MatchPlayer
	if mps, err = s.repository.LoadPlayersByMatchIDAndTeamID(matchID, teamID); err != nil {
		return players, err
	}

	for _, mp := range mps {
		var p compelo.Player
		if p, err = s.playerService.LoadPlayerByID(mp.PlayerID); err == nil {
			players = append(players, p)
		} else {
			return players, err
		}
	}

	return players, err
}
