package project

import (
	"compelo/db"
	"compelo/models"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateProject(name string) (*models.Project, error) {
	p := &models.Project{Name: name}
	err := s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadProjects() []models.Project {
	var projects []models.Project
	s.db.Find(&projects)
	return projects
}
