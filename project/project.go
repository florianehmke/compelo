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

func (s *Service) CreateProject(name string, hash []byte) (*models.Project, error) {
	p := &models.Project{Name: name, PasswordHash: hash}
	err := s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadByName(name string) (models.Project, error) {
	var project models.Project
	err := s.db.Where(&models.Project{Name: name}).Find(&project).Error
	return project, err
}

func (s *Service) LoadProjects() []models.Project {
	var projects []models.Project
	s.db.Find(&projects)
	return projects
}
