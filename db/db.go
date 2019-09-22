//go:generate go run scripts_generate.go
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
	db.LogMode(false)

	f, err := Scripts.Open("schema.sql")
	if err != nil {
		panic(err)
	}
	schema, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	db.Exec(string(schema))

	return &DB{db}
}
