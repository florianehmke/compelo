package main

import (
	"compelo/api"
	"compelo/db"
	"compelo/project"
)

func main() {
	database := db.New()
	api.Serve(project.NewRouter(project.NewService(database)))
}
