package player

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

func (s *Service) CreatePlayer(projectID uint, name string) (*models.Player, error) {
	p := &models.Player{
		Name:      name,
		ProjectID: projectID,
	}
	err := s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadPlayers() []models.Player {
	var players []models.Player
	s.db.Find(&players)
	return players
}
