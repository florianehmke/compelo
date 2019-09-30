package compelo

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"compelo/internal/db"
)

func (svc *Service) CreateProject(name, pw string) (db.Project, error) {
	return svc.db.CreateProject(db.Project{
		Name:         name,
		PasswordHash: hashAndSalt([]byte(pw)),
	})
}

func (svc *Service) LoadAllProjects() []db.Project {
	return svc.db.LoadAllProjects()
}

func (svc *Service) LoadProjectByName(name string) (db.Project, error) {
	return svc.db.LoadProjectByName(name)
}

func (svc *Service) LoadProjectByID(id uint) (db.Project, error) {
	return svc.db.LoadProjectByID(id)
}

func (svc *Service) AuthorizeProject(name, pw string) (db.Project, error) {
	p, err := svc.LoadProjectByName(name)
	if err != nil {
		return db.Project{}, errors.New("unknown project")
	}
	err = bcrypt.CompareHashAndPassword(p.PasswordHash, []byte(pw))
	if err != nil {
		return db.Project{}, errors.New("wrong password for project")
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
