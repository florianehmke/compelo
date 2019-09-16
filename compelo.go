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

type Match struct {
	Model

	Date   time.Time `json:"date"`
	GameID uint      `json:"gameId"`
}

type MatchTeam struct {
	Model

	MatchID uint `json:"matchId"`
	Score   int  `json:"score"`
	Winner  bool `json:"winner"`
}

type MatchPlayer struct {
	Model

	MatchID     uint `json:"matchId"`
	MatchTeamID uint `json:"matchTeamId"`
	PlayerID    uint `json:"playerId"`
}
