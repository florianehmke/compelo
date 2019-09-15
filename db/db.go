package db

import (
	"compelo"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	*gorm.DB
}

func (db *DB) Close() {
	if err := db.DB.Close(); err != nil {
		log.Fatal(err)
	}
}

func New(dbPath string) *DB {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("failed to connect database")
	}

	db.Exec("PRAGMA foreign_keys = ON")
	db.DB().SetMaxOpenConns(1)
	db.LogMode(true)

	db.AutoMigrate(&compelo.Project{})
	db.AutoMigrate(&compelo.Player{})
	db.AutoMigrate(&compelo.Game{})
	db.AutoMigrate(&compelo.Match{})
	db.AutoMigrate(&compelo.MatchPlayer{})
	db.AutoMigrate(&compelo.MatchTeam{})

	return &DB{db}
}
