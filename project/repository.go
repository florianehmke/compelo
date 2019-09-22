package project

import "compelo/db"

type Project struct {
	db.Model

	Name         string `json:"name"`
	PasswordHash []byte `json:"-"`
}
