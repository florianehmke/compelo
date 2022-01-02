package tests

import (
	"log"
	"os"
	"testing"

	"compelo/api/handler"
	"compelo/query"
)

func TestReferenceProject(t *testing.T) {
	suite := createAPITestSuite(t)
	defer os.Remove(suite.dbName)

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
		1: {
			Teams: []handler.CreateMatchRequestTeam{
				{
					PlayerGUIDs: []string{suite.playerGUIDs[0]},
					Score:       0,
				},
				{
					PlayerGUIDs: []string{suite.playerGUIDs[1]},
					Score:       3,
				},
			},
		},
	}
	suite.competitionRequest = handler.CreateCompetitionRequest{
		Rounds: 2,
		Name:   "Competition 1",
		Teams: []handler.CreateCompetitionRequestTeam{
			{PlayerGUIDs: []string{suite.playerGUIDs[0]}},
			{PlayerGUIDs: []string{suite.playerGUIDs[1]}},
			{PlayerGUIDs: []string{suite.playerGUIDs[2]}},
			{PlayerGUIDs: []string{suite.playerGUIDs[3]}},
		},
	}

	suite.createMatches()

	suite.expectedMatches = map[int]query.Match{
		0: {
			GUID: suite.matchGUIDs[0],
		},
		1: {
			GUID: suite.matchGUIDs[1],
		},
	}

	suite.listMatches()
	suite.loadGameStats()
	suite.loadPlayerStats()

	suite.createCompetition()
	suite.listCompetitions()

	log.Println("Done!")
}
