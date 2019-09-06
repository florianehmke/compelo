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
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Project{})

	return &DB{db}
}
