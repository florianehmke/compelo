package project

import (
	"compelo/db"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateProject(name string, hash []byte) (Project, error) {
	p := Project{Name: name, PasswordHash: hash}
	err := s.db.Create(&p).Error
	return p, err
}

func (s *Service) LoadByName(name string) (Project, error) {
	var project Project
	err := s.db.Where(&Project{Name: name}).Find(&project).Error
	return project, err
}

func (s *Service) LoadProjects() []Project {
	var projects []Project
	s.db.Find(&projects)
	return projects
}
