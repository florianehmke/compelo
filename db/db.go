package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"compelo/models"
)

type DB struct {
	*gorm.DB
}

func (db *DB) Close() {
	if err := db.DB.Close(); err != nil {
		log.Fatal(err)
	}
}

func New() *DB {
	db, err := gorm.Open("sqlite3", "file::memory:")
	db.Exec("PRAGMA foreign_keys = ON")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Player{})
	db.AutoMigrate(&models.Game{})
	db.AutoMigrate(&models.Match{})
	db.AutoMigrate(&models.MatchPlayer{})
	db.AutoMigrate(&models.MatchTeam{})

	return &DB{db}
}
