package compelo

import (
	"compelo/internal/db"
)

func (svc *Service) CreateGame(projectID uint, name string) (db.Game, error) {
	return svc.db.CreateGame(db.Game{
		Name:      name,
		ProjectID: projectID,
	})
}

func (svc *Service) LoadGamesByProjectID(projectID uint) []db.Game {
	return svc.db.LoadGamesByProjectID(projectID)
}

func (svc *Service) LoadGameByID(id uint) (db.Game, error) {
	return svc.db.LoadGameByID(id)
}
