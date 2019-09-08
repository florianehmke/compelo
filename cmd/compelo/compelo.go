package main

import (
	"compelo/api"
	"compelo/db"
	"compelo/game"
	"compelo/match"
	"compelo/player"
	"compelo/project"
)

func main() {
	database := db.New()

	projectService := project.NewService(database)
	playerService := player.NewService(database)
	gameService := game.NewService(database)
	matchService := match.NewService(database, playerService, gameService)

	projectService.CreateProject("My Project", []byte("test"))
	playerService.CreatePlayer(1, "Player 1")
	playerService.CreatePlayer(1, "Player 2")
	gameService.CreateGame(1, "FIFA")

	api.Serve(
		project.NewRouter(projectService),
		player.NewRouter(playerService),
		match.NewRouter(matchService),
		game.NewRouter(gameService),
	)
}
