package models

import (
	"time"
)

//
// https://github.com/jinzhu/gorm/issues/2006
//

type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type Project struct {
	Model

	Name string `json:"name" gorm:"unique;not null"`
}

type Game struct {
	Model

	Name      string `json:"name" gorm:"unique;not null"`
	ProjectID uint   `json:"projectId" gorm:"type:int REFERENCES projects(id) ON DELETE CASCADE"`
}

type Player struct {
	Model

	Name      string `json:"name" gorm:"unique;not null"`
	ProjectID uint   `json:"projectId" gorm:"type:int REFERENCES projects(id) ON DELETE CASCADE"`
}

type Match struct {
	Model

	Date              time.Time `json:"date" gorm:"not null"`
	GameID            uint      `json:"gameId" gorm:"not null"`
	WinnerMatchTeamID uint      `json:"winnerMatchTeamId" gorm:"not null"`
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
