package tests

import (
	"os"
	"testing"

	"compelo/api/handler"
)

func TestReferenceProject(t *testing.T) {
	defer os.Remove("reference.db")
	suite := createAPITestSuite(t, "reference.db")

	suite.projectRequest = handler.CreateProjectRequest{
		Name:     "Reference Project",
		Password: "Secure Password",
	}
	suite.gameRequest = handler.CreateGameRequest{
		Name: "Reference Game",
	}
	suite.playerRequests = map[int]handler.CreatePlayerRequest{
		0: {Name: "Player 1"},
		1: {Name: "Player 2"},
		2: {Name: "Player 3"},
		3: {Name: "Player 4"},
	}

	suite.createProject()
	suite.listProjects()
	suite.selectProject()

	suite.createGame()
	suite.listGames()

	suite.createPlayers()
	suite.listGames()

	suite.matchRequests = map[int]handler.CreateMatchRequest{
		0: {
			Teams: []handler.CreateMatchRequestTeam{
				{
					PlayerGUIDs: []string{suite.playerGUIDs[0], suite.playerGUIDs[1]},
					Score:       1,
				},
				{
					PlayerGUIDs: []string{suite.playerGUIDs[2], suite.playerGUIDs[3]},
					Score:       2,
				},
			},
		},
	}

	suite.createMatches()
	suite.listMatches()
	suite.loadGameStats()
	suite.loadPlayerStats()
}
