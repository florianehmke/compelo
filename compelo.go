package compelo

import (
	"time"
)

type Model struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type Project struct {
	Model

	Name         string `json:"name"`
	PasswordHash []byte `json:"-"`
}

type Game struct {
	Model

	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}

type Player struct {
	Model

	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}

type Rating struct {
	Model

	Rating   float64 `json:"rating"`
	GameID   uint    `json:"gameId"`
	PlayerID uint    `json:"playerId"`
}

type Match struct {
	Model

	Date   time.Time `json:"date"`
	GameID uint      `json:"gameId"`
}

type Team struct {
	Model

	MatchID uint `json:"matchId"`
	Score   int  `json:"score"`
	Winner  bool `json:"winner"`
}

type Appearance struct {
	Model

	MatchID  uint `json:"matchId"`
	TeamID   uint `json:"teamId"`
	PlayerID uint `json:"playerId"`
}
