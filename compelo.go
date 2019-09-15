package compelo

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

	Name         string `json:"name" gorm:"unique;not null"`
	PasswordHash []byte `json:"-" gorm:"not null"`
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

	Date   time.Time `json:"date" gorm:"not null"`
	GameID uint      `json:"gameId" gorm:"type:int REFERENCES games(id) ON DELETE CASCADE"`
}

type MatchTeam struct {
	Model

	MatchID uint `json:"matchId" gorm:"type:int REFERENCES matches(id) ON DELETE CASCADE"`
	Score   int  `json:"score"`
	Winner  bool `json:"winner"`
}

type MatchPlayer struct {
	Model

	MatchID     uint `json:"matchId" gorm:"type:int REFERENCES match_teams(id) ON DELETE CASCADE"`
	MatchTeamID uint `json:"matchTeamId" gorm:"type:int REFERENCES match_teams(id) ON DELETE CASCADE"`
	PlayerID    uint `json:"playerId" gorm:"type:int REFERENCES players(id)"`
}
