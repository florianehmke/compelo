package game

import (
	"compelo/db"
	"compelo/models"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateGame(projectID uint, name string) (*models.Game, error) {
	p := &models.Game{
		Name:      name,
		ProjectID: projectID,
	}
	err := s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadGames() []models.Game {
	var games []models.Game
	s.db.Find(&games)
	return games
}

func (s *Service) LoadGameByID(id uint) (models.Game, error) {
	var game models.Game
	err := s.db.First(&game, id).Error
	return game, err
}
