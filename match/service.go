package match

import (
	"compelo"
	"compelo/db"
	"compelo/game"
	"compelo/player"
	"time"
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

	Teams         int
	WinningTeam   int
	PlayerTeamMap map[uint]int
	TeamScoreMap  map[int]int
}

// TODO wrap in txn
func (s *Service) CreateMatch(param CreateMatchParameter) (compelo.Match, error) {
	if err := s.validate(param); err != nil {
		return compelo.Match{}, err
	}

	m := compelo.Match{GameID: param.GameID, Date: time.Now()}

	teamMap := map[int]compelo.MatchTeam{}
	playerMap := map[int][]compelo.MatchPlayer{}

	for i := 1; i <= param.Teams; i++ {
		teamMap[i] = compelo.MatchTeam{
			Score:  param.TeamScoreMap[i],
			Winner: i == param.WinningTeam,
		}

		for playerID, team := range param.PlayerTeamMap {
			if team == i {
				playerMap[i] = append(playerMap[i], compelo.MatchPlayer{
					PlayerID: playerID,
				})
			}
		}
	}
	return s.repository.Create(m, teamMap, playerMap)
}

func (s *Service) validate(param CreateMatchParameter) error {
	if _, err := s.gameService.LoadGameByID(param.GameID); err != nil {
		return err
	}

	for playerID, _ := range param.PlayerTeamMap {
		if _, err := s.playerService.LoadPlayerByID(playerID); err != nil {
			return err
		}
	}

	return nil
}

type CompleteMatch struct {
	Match        compelo.Match         `json:"match"`
	MatchTeams   []compelo.MatchTeam   `json:"teams"`
	MatchPlayers []compelo.MatchPlayer `json:"players"`
}

func (s *Service) LoadByID(id uint) (CompleteMatch, error) {
	match := CompleteMatch{}
	var err error

	if match.Match, err = s.repository.LoadByID(id); err != nil {
		return match, err
	}
	if match.MatchTeams, err = s.repository.LoadTeamsByMatchID(id); err != nil {
		return match, err
	}
	if match.MatchPlayers, err = s.repository.LoadPlayersByMatchID(id); err != nil {
		return match, err
	}

	return match, nil
}
