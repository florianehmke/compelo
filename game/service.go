package game

import (
	"compelo/db"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateGame(projectID uint, name string) (Game, error) {
	p := Game{
		Name:      name,
		ProjectID: projectID,
	}
	err := s.db.Create(&p).Error
	return p, err
}

func (s *Service) LoadGamesByProjectID(pid uint) ([]Game, error) {
	var games []Game
	err := s.db.Where(&Game{ProjectID: pid}).Find(&games).Error
	return games, err
}

func (s *Service) LoadGameByID(id uint) (Game, error) {
	var game Game
	err := s.db.First(&game, id).Error
	return game, err
}
