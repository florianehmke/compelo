package game

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

func (s *Service) CreateGame(projectID uint, name string) (*compelo.Game, error) {
	p := &compelo.Game{
		Name:      name,
		ProjectID: projectID,
	}
	err := s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadGamesByProjectID(pid uint) ([]compelo.Game, error) {
	var games []compelo.Game
	err := s.db.Where(&compelo.Game{ProjectID: pid}).Find(&games).Error
	return games, err
}

func (s *Service) LoadGameByID(id uint) (compelo.Game, error) {
	var game compelo.Game
	err := s.db.First(&game, id).Error
	return game, err
}
