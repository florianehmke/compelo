package db

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"compelo/internal/db/scripts"
)

type gormDB struct {
	gorm *gorm.DB
}

type Database interface {
	DoInTransaction(fn func(Database) error) error
	Close()

	CreateGame(Game) (Game, error)
	LoadGamesByProjectID(projectID uint) []Game
	LoadGameByID(id uint) (Game, error)

	CreateMatch(match Match) (Match, error)
	LoadMatchByID(id uint) (Match, error)
	LoadMatchesByGameID(gameID uint) []Match

	CreateAppearance(appearance Appearance) (Appearance, error)
	CreateTeam(team Team) (Team, error)
	LoadTeamsByMatchID(matchID uint) ([]Team, error)

	CreatePlayer(player Player) (Player, error)
	LoadPlayerByID(id uint) (Player, error)
	LoadPlayersByProjectID(projectID uint) []Player
	LoadPlayersByMatchIDAndTeamID(matchID, teamID uint) ([]Player, error)

	CreateProject(project Project) (Project, error)
	LoadProjectByID(id uint) (Project, error)
	LoadProjectByName(name string) (Project, error)
	LoadAllProjects() []Project

	LoadOrCreateRatingByPlayerIDAndGameID(playerID, gameID uint) (Rating, error)
	SaveRating(Rating) (Rating, error)

	LoadMatchResultsByGameID(gameID uint) ([]MatchResult, error)
}

var _ Database = &gormDB{}

type Model struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (db *gormDB) Close() {
	if err := db.gorm.Close(); err != nil {
		log.Fatal(err)
	}
}

func New(dbPath string) *gormDB {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("failed to connect database")
	}

	db.DB().SetMaxOpenConns(1)
	db.Exec("PRAGMA foreign_keys = ON")
	db.LogMode(false)

	f, err := scripts.Scripts.Open("schema.sql")
	if err != nil {
		panic(err)
	}
	schema, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	db.Exec(string(schema))

	return &gormDB{db}
}

func (db *gormDB) DoInTransaction(fn func(Database) error) error {
	tx := db.gorm.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := fn(&gormDB{tx})
	if err != nil {
		xerr := tx.Rollback().Error
		if xerr != nil {
			return xerr
		}
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
