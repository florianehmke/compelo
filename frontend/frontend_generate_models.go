//go:build ignore

package main

import (
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"

	"compelo/api/handler"
	"compelo/api/security"
	"compelo/query"
)

const modelPath = "compelo/src/generated/api/"

func main() {
	generateApiModels()
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

	converter.Add(query.Project{})
	converter.Add(query.Player{})
	converter.Add(query.Game{})
	converter.Add(query.Match{})

	converter.Add(query.PlayerStats{})
	converter.Add(query.GameStats{})

	err := converter.ConvertToFile(modelPath + "api.models.ts")
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
