package db

import (
	"io/ioutil"
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

	db.DB().SetMaxOpenConns(1)
	db.Exec("PRAGMA foreign_keys = ON")
	db.LogMode(true)

	query, err := ioutil.ReadFile("db/schema.sql")
	if err != nil {
		panic(err)
	}
	db.Exec(string(query))

	return &DB{db}
}
