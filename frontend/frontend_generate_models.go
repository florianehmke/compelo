//go:build ignore

package main

import (
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"

	compelo "compelo/internal"

	"compelo/internal/api/handler"
	"compelo/internal/api/security"
	"compelo/internal/db"
)

const modelPath = "compelo/src/generated/api/"

func main() {
	generateApiModels()
	generateAppModels()
	generateDatabaseModels()
}

func generateApiModels() {
	converter := newConverter()
	converter.Add(security.AuthRequest{})
	converter.Add(security.AuthResponse{})
	converter.Add(handler.CreateProjectRequest{})
	converter.Add(handler.CreateGameRequest{})
	converter.Add(handler.CreatePlayerRequest{})
	converter.Add(handler.CreateMatchRequest{})
	converter.Add(handler.CreateMatchRequestTeam{})
	err := converter.ConvertToFile(modelPath + "api.models.ts")
	if err != nil {
		panic(err.Error())
	}
}

func generateAppModels() {
	converter := newConverter()
	converter.Add(compelo.MatchData{})
	converter.Add(compelo.PlayerStats{})
	converter.Add(compelo.GameStats{})
	err := converter.ConvertToFile(modelPath + "app.models.ts")
	if err != nil {
		panic(err.Error())
	}
}

func generateDatabaseModels() {
	converter := newConverter()
	converter.Add(db.Game{})
	converter.Add(db.Match{})
	converter.Add(db.Team{})
	converter.Add(db.Appearance{})
	converter.Add(db.MatchResult{})
	converter.Add(db.MatchScoreStats{})
	converter.Add(db.Player{})
	converter.Add(db.Project{})
	converter.Add(db.Rating{})
	err := converter.ConvertToFile(modelPath + "db.models.ts")
	if err != nil {
		panic(err.Error())
	}
}

func newConverter() *typescriptify.TypeScriptify {
	converter := typescriptify.New()
	converter.CreateInterface = true
	converter.Indent = "  "
	converter.BackupDir = ""
	return converter
}
