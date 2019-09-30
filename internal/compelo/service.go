package compelo

import (
	"compelo/internal/db"
)

type Service struct {
	db *db.DB
}

func NewService(dbPath string) *Service {
	return &Service{db.New(dbPath)}
}
