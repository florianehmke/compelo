package project

import (
	"compelo"
	"compelo/db"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateProject(name string, hash []byte) (*compelo.Project, error) {
	p := &compelo.Project{Name: name, PasswordHash: hash}
	err := s.db.Create(p).Error
	return p, err
}

func (s *Service) LoadByName(name string) (compelo.Project, error) {
	var project compelo.Project
	err := s.db.Where(&compelo.Project{Name: name}).Find(&project).Error
	return project, err
}

func (s *Service) LoadProjects() []compelo.Project {
	var projects []compelo.Project
	s.db.Find(&projects)
	return projects
}
