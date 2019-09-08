package player

import (
	"compelo"
	"compelo/db"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreatePlayer(projectID uint, name string) (*compelo.Player, error) {
	p := &compelo.Player{
		Name:      name,
		ProjectID: projectID,
	}
	err := s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadPlayersByProjectID(pid uint) ([]compelo.Player, error) {
	var players []compelo.Player
	err := s.db.Where(&compelo.Player{ProjectID: pid}).Find(&players).Error
	return players, err
}

func (s *Service) LoadPlayerByID(id uint) (compelo.Player, error) {
	var player compelo.Player
	err := s.db.First(&player, id).Error
	return player, err
}
