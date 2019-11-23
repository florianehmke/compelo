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

var (
	ErrProjectDoesNotExist  = errors.New("project does not exist")
	ErrWrongProjectPassword = errors.New("wrong password for project")
)

func (svc *Service) AuthenticateProject(id uint, pw string) (db.Project, error) {
	p, err := svc.LoadProjectByID(id)
	if err != nil {
		return db.Project{}, ErrProjectDoesNotExist
	}
	err = bcrypt.CompareHashAndPassword(p.PasswordHash, []byte(pw))
	if err != nil {
		return db.Project{}, ErrWrongProjectPassword
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
