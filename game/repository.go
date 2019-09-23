package game

import "compelo/db"

type Game struct {
	db.Model

	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}
