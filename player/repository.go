package player

import "compelo/db"

type Player struct {
	db.Model

	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}

type Rating struct {
	db.Model

	Rating   int  `json:"rating"`
	GameID   uint `json:"gameId"`
	PlayerID uint `json:"playerId"`
}
