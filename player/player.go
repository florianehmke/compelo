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

func (s *Service) LoadPlayersByProjectID(pid uint) ([]models.Player, error) {
	var players []models.Player
	err := s.db.Where(&models.Player{ProjectID: pid}).Find(&players).Error
	return players, err
}

func (s *Service) LoadPlayerByID(id uint) (models.Player, error) {
	var player models.Player
	err := s.db.First(&player, id).Error
	return player, err
}
