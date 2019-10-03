//go:generate go run scripts_generate.go
package db

import (
	"errors"
	"io/ioutil"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	gorm *gorm.DB
}

type Model struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

var RecordNotFound = errors.New("record not found")

func (db *DB) Close() {
	if err := db.gorm.Close(); err != nil {
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

func (db *DB) DoInTransaction(fn func(*DB) error) error {
	tx := db.gorm.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := fn(&DB{tx})
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
