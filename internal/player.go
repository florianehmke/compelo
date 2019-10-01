package compelo

import (
	"compelo/internal/db"
)

func (svc *Service) CreatePlayer(projectID uint, name string) (db.Player, error) {
	return svc.db.CreatePlayer(db.Player{
		Name:      name,
		ProjectID: projectID,
	})
}

func (svc *Service) LoadPlayersByProjectID(projectID uint) []db.Player {
	return svc.db.LoadPlayersByProjectID(projectID)
}

func (svc *Service) LoadPlayerByID(id uint) (db.Player, error) {
	return svc.db.LoadPlayerByID(id)
}
