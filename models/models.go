package models

import (
	"time"
)

type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

type Project struct {
	Model

	Name string `json:"name" gorm:"unique;not null"`
}

type Game struct {
	Model

	Name      string
	ProjectID uint
}

type Player struct {
	Model

	Name      string
	ProjectID uint
}

type Match struct {
	Model

	GameID            uint
	WinnerMatchTeamID uint
}

type MatchTeam struct {
	Model

	MatchID uint
	Score   int
}

type MatchPlayer struct {
	Model

	MatchID  uint
	PlayerID uint
}
