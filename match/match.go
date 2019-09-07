package match

import (
	"compelo/db"
	"compelo/game"
	"compelo/models"
	"compelo/player"
	"time"
)

type Service struct {
	db *db.DB

	playerService *player.Service
	gameService   *game.Service
}

func NewService(
	db *db.DB,
	playerService *player.Service,
	gameService *game.Service,
) *Service {
	return &Service{
		db:            db,
		playerService: playerService,
		gameService:   gameService,
	}
}

type CreateMatchParameter struct {
	Date            time.Time
	GameID          uint
	Teams           uint
	PlayerTeamMap   map[uint]uint
	TeamScoreMap    map[uint]int
	WinnerMatchTeam uint
}

func (s *Service) CreateMatch(param CreateMatchParameter) (*models.Match, error) {
	// Verify that the game exists.
	g, err := s.gameService.LoadGameByID(param.GameID)
	if err != nil {
		return nil, err // FIXME
	}

	// Create the match.
	m, err := s.createMatch(param.Date, g.ID)
	if err != nil {
		return nil, err // FIXME
	}

	// Create the teams.
	teamMap := map[uint]*models.MatchTeam{}
	for i := 1; i <= int(param.Teams); i++ {
		t, err := s.createTeam(m.ID, param.TeamScoreMap[uint(i)]) // FIXME maybe validate scores
		if err != nil {
			return nil, err // FIXME
		}
		teamMap[uint(i)] = t
	}

	// Create the players.
	playerIdMatchPlayerMap := map[uint]*models.MatchPlayer{}
	for playerID, teamNumber := range param.PlayerTeamMap {
		p, err := s.playerService.LoadPlayerByID(playerID)
		if err != nil {
			return nil, err // FIXME
		}
		mp, err := s.createPlayer(m.ID, teamMap[teamNumber].ID, p.ID)
		if err != nil {
			return nil, err // FIXME
		}
		playerIdMatchPlayerMap[playerID] = mp
	}

	p := &models.Match{
		Date:              param.Date,
		WinnerMatchTeamID: 0,
		GameID:            g.ID,
	}
	err = s.db.Create(p).Error
	return p, err
}

func (s *Service) createMatch(date time.Time, gameID uint) (*models.Match, error) {
	m := &models.Match{Date: date, GameID: gameID}
	err := s.db.Create(m).Error
	return m, err
}

// FIXME (score)
func (s *Service) createTeam(matchID uint, score int) (*models.MatchTeam, error) {
	t := &models.MatchTeam{MatchID: matchID, Score: score}
	err := s.db.Create(t).Error
	return t, err
}

func (s *Service) createPlayer(matchID uint, matchTeamID uint, playerID uint) (*models.MatchPlayer, error) {
	p := &models.MatchPlayer{MatchID: matchID, MatchTeamID: matchTeamID, PlayerID: playerID}
	err := s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadMatches() []models.Match {
	var matches []models.Match
	s.db.Find(&matches)
	return matches
}

type CompleteMatch struct {
	Match        models.Match         `json:"match"`
	MatchTeams   []models.MatchTeam   `json:"teams"`
	MatchPlayers []models.MatchPlayer `json:"players"`
}

func (s *Service) LoadByID(id uint) (CompleteMatch, error) {
	var match models.Match
	err := s.db.First(&match, id).Error
	if err != nil {
		return CompleteMatch{}, err // FIXME
	}

	var players []models.MatchPlayer
	s.db.Find(&players) // FIXME by match id

	var teams []models.MatchTeam
	s.db.Find(&teams) // FIXME by match id

	return CompleteMatch{
		Match:        match,
		MatchTeams:   teams,
		MatchPlayers: players,
	}, nil
}
