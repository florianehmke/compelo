package match

import (
	"compelo"
	"compelo/db"
	"compelo/game"
	"compelo/player"
	"sort"
	"time"
)

type Service struct {
	repository    Repository
	playerService *player.Service
	gameService   *game.Service
}

func NewService(db *db.DB, ps *player.Service, gs *game.Service) *Service {
	return &Service{
		repository:    repository{db},
		playerService: ps,
		gameService:   gs,
	}
}

func (s *Service) CreateMatch(param CreateMatchParameter, game compelo.Game) (compelo.Match, error) {
	param.GameID = game.ID
	param.Date = time.Now()
	param.determineWinner()

	return s.repository.Create(param)
}

func (p *CreateMatchParameter) determineWinner() {
	highScore := 0
	highScoreCount := 0
	for _, t := range p.Teams {
		if t.Score > highScore {
			highScore = t.Score
			highScoreCount = 1
		} else if t.Score == highScore {
			highScoreCount += 1
		}
	}
	if highScoreCount == 1 {
		for _, t := range p.Teams {
			if t.Score == highScore {
				t.Winner = true
			}
		}
	}
}

type MatchData struct {
	ID     uint       `json:"id"`
	Date   time.Time  `json:"date"`
	GameID uint       `json:"gameId"`
	Teams  []TeamData `json:"teams"`
}

type TeamData struct {
	ID      uint         `json:"id"`
	MatchID uint         `json:"matchId"`
	Score   int          `json:"score"`
	Winner  bool         `json:"winner"`
	Players []PlayerData `json:"players"`
}

type PlayerData struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}

func (s *Service) LoadByGameID(gameID uint) ([]MatchData, error) {
	var matchDataList []MatchData

	matches, err := s.repository.LoadByGameID(gameID)
	for _, match := range matches {
		matchData, err := s.LoadByID(match.ID)
		if err != nil {
			return matchDataList, err
		}
		matchDataList = append(matchDataList, matchData)
	}

	return matchDataList, err
}

func (s *Service) LoadByID(id uint) (MatchData, error) {
	// 1. Get basic match data.
	match, err := s.repository.LoadByID(id)
	if err != nil {
		return MatchData{}, err
	}
	matchData := MatchData{
		ID:     match.ID,
		Date:   match.Date,
		GameID: match.GameID,
	}

	// 2. Get data about teams.
	teams, err := s.repository.LoadTeamsByMatchID(id)
	if err != nil {
		return MatchData{}, err
	}
	for _, t := range teams {
		teamData := TeamData{
			ID:      t.ID,
			MatchID: t.MatchID,
			Score:   t.Score,
			Winner:  t.Winner,
		}

		// 3. Get data about players.
		players, err := s.LoadPlayersByMatchIDAndTeamID(id, t.ID)
		if err != nil {
			return MatchData{}, err
		}
		for _, p := range players {
			playerData := PlayerData{
				ID:        p.ID,
				Name:      p.Name,
				ProjectID: p.ProjectID,
			}
			teamData.Players = append(teamData.Players, playerData)
		}
		matchData.Teams = append(matchData.Teams, teamData)
	}

	// 4. Sort teams by score.
	sort.Slice(matchData.Teams, func(i, j int) bool {
		if matchData.Teams[i].Score == matchData.Teams[j].Score {
			return matchData.Teams[i].ID < matchData.Teams[j].ID
		}
		return matchData.Teams[i].Score > matchData.Teams[j].Score
	})

	return matchData, err
}

func (s *Service) LoadPlayersByMatchIDAndTeamID(matchID, teamID uint) ([]compelo.Player, error) {
	appearances, err := s.repository.LoadAppearancesByMatchIDAndTeamID(matchID, teamID)
	if err != nil {
		return nil, err
	}

	var players []compelo.Player
	for _, appearance := range appearances {
		p, err := s.playerService.LoadPlayerByID(appearance.PlayerID)
		if err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, err
}
