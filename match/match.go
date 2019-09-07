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
	WinnerMatchTeam uint
}

func (s *Service) CreateMatch(param CreateMatchParameter) (*models.Match, error) {
	g, err := s.gameService.LoadGameByID(param.GameID)
	if err != nil {
		return nil, err // FIXME
	}

	p := &models.Match{
		Date:              param.Date,
		WinnerMatchTeamID: 0,
		GameID:            g.ID,
	}
	err = s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadMatches() []models.Match {
	var matches []models.Match
	s.db.Find(&matches)
	return matches
}
