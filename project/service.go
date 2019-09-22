package project

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"compelo/db"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateProject(name, pw string) (Project, error) {
	p := Project{Name: name, PasswordHash: hashAndSalt([]byte(pw))}
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

func (s *Service) AuthorizeProject(name, pw string) (Project, error) {
	p, err := s.LoadByName(name)
	if err != nil {
		return Project{}, errors.New("unknown project")
	}
	err = bcrypt.CompareHashAndPassword(p.PasswordHash, []byte(pw))
	if err != nil {
		return Project{}, errors.New("wrong password for project")
	}
	return p, nil
}

func hashAndSalt(pwd []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return hash
}
