package tests

import (
	"compelo/command"
	"compelo/event"
	"compelo/query"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type basicProject struct {
	projectName string
	projectGUID string
	gameName    string
	gameGUID    string
	players     []*basicPlayer
	matchGUID   string
}

type basicPlayer struct {
	guid  string
	name  string
	score int
}

func Test(t *testing.T) {
	log.Println("Starting Test")
	defer os.Remove("cq_test.db")

	// Create event store.
	bus := event.NewBus()
	store := event.NewStore(bus, "cq_test.db")

	// Simulate some prior events to ensure re-hydration works.
	store.StoreEvent(&event.ProjectCreated{GUID: "guid", Name: "First Project"})

	// Setup query.
	query := query.NewService(bus)

	// Load all events from db (rehydrates queries).
	events, err := store.LoadEvents()
	assert.Nil(t, err)

	// Setup command (from existing events).
	command := command.NewService(store, events)

	// Simulate interaction with command.
	testBasicWorkflow(t, command, query)
}

func testBasicWorkflow(t *testing.T, c *command.Service, q *query.Service) {
	var testProject = basicProject{
		projectName: "Project 1",
		gameName:    "Game 1",
		players: []*basicPlayer{
			{name: "Player 1", score: 1},
			{name: "Player 2", score: 2},
		},
	}

	// 1. Create a project.
	response, err := c.CreateNewProject(command.CreateNewProjectCommand{
		Name: testProject.projectName,
	})
	assert.Nil(t, err)
	testProject.projectGUID = response.GUID

	// 2. Create two players.
	for _, p := range testProject.players {
		response, err := c.CreateNewPlayer(command.CreateNewPlayerCommand{
			Name:        p.name,
			ProjectGUID: testProject.projectGUID,
		})
		assert.Nil(t, err)
		p.guid = response.GUID
	}

	// 3. Create a game.
	response, err = c.CreateNewGame(command.CreateNewGameCommand{
		Name:        testProject.gameName,
		ProjectGUID: testProject.projectGUID,
	})
	assert.Nil(t, err)
	testProject.gameGUID = response.GUID

	// 4. Create a match.
	response, err = c.CreateNewMatch(command.CreateNewMatchCommand{
		GameGUID:    testProject.gameGUID,
		ProjectGUID: testProject.projectGUID,
		Teams: []struct {
			PlayerGUIDs []string
			Score       int
		}{
			{Score: testProject.players[0].score, PlayerGUIDs: []string{testProject.players[0].guid}},
			{Score: testProject.players[1].score, PlayerGUIDs: []string{testProject.players[1].guid}},
		},
	})
	assert.Nil(t, err)
	testProject.matchGUID = response.GUID

	checkCommandResults(t, testProject)
	checkQuery(t, q, testProject)

}

func checkCommandResults(t *testing.T, testProject basicProject) {
	assert.NotEmpty(t, testProject.projectGUID)
	assert.NotEmpty(t, testProject.players[0].guid)
	assert.NotEmpty(t, testProject.players[1].guid)
	assert.NotEmpty(t, testProject.gameGUID)
	assert.NotEmpty(t, testProject.matchGUID)
}

func checkQuery(t *testing.T, q *query.Service, testProject basicProject) {
	checkQueryGetProjects(t, q)
	checkQueryGetPlayersBy(t, q, testProject)
	checkQueryGetGamesBy(t, q, testProject)
	checkQueryGetMatchesBy(t, q, testProject)

	checkQueryGetProjectBy(t, q, testProject)
	checkQueryGetGameBy(t, q, testProject)
	checkQueryGetPlayerBy(t, q, testProject)
	checkQueryGetMatchBy(t, q, testProject)
	checkQueryGetRatingBy(t, q, testProject)
}

func checkQueryGetProjects(t *testing.T, q *query.Service) {
	assert.Len(t, q.GetProjects(), 2)
}

func checkQueryGetPlayersBy(t *testing.T, q *query.Service, testProject basicProject) {
	players, err := q.GetPlayersBy(testProject.projectGUID)
	assert.Len(t, players, 2)
	assert.Nil(t, err)

	players, err = q.GetPlayersBy("404")
	assert.Nil(t, players)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))
}

func checkQueryGetGamesBy(t *testing.T, q *query.Service, testProject basicProject) {
	games, err := q.GetGamesBy(testProject.projectGUID)
	assert.Len(t, games, 1)
	assert.Nil(t, err)

	games, err = q.GetGamesBy("404")
	assert.Nil(t, games)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))
}

func checkQueryGetMatchesBy(t *testing.T, q *query.Service, testProject basicProject) {
	matches, err := q.GetMatchesBy(testProject.projectGUID, testProject.gameGUID)
	assert.Len(t, matches, 1)
	assert.Nil(t, err)

	matches, err = q.GetMatchesBy("404", testProject.gameGUID)
	assert.Nil(t, matches)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))

	matches, err = q.GetMatchesBy(testProject.projectGUID, "404")
	assert.Nil(t, matches)
	assert.True(t, errors.Is(err, query.ErrGameNotFound))
}

func checkQueryGetProjectBy(t *testing.T, q *query.Service, testProject basicProject) {
	project, err := q.GetProjectBy(testProject.projectGUID)
	assert.NotNil(t, project)
	assert.Nil(t, err)
	assert.Equal(t, testProject.projectName, project.Name)

	project, err = q.GetProjectBy("404")
	assert.Nil(t, project)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))
}

func checkQueryGetGameBy(t *testing.T, q *query.Service, testProject basicProject) {
	game, err := q.GetGameBy(testProject.projectGUID, testProject.gameGUID)
	assert.NotNil(t, game)
	assert.Nil(t, err)
	assert.Equal(t, testProject.gameName, game.Name)
	assert.Equal(t, testProject.projectGUID, game.ProjectGUID)

	game, err = q.GetGameBy("404", testProject.gameGUID)
	assert.Nil(t, game)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))

	game, err = q.GetGameBy(testProject.projectGUID, "404")
	assert.Nil(t, game)
	assert.True(t, errors.Is(err, query.ErrGameNotFound))
}

func checkQueryGetPlayerBy(t *testing.T, q *query.Service, testProject basicProject) {
	for _, p := range testProject.players {
		player, err := q.GetPlayerBy(testProject.projectGUID, p.guid)
		assert.NotNil(t, player)
		assert.Nil(t, err)
		assert.Equal(t, p.name, player.Name)
		assert.Equal(t, testProject.projectGUID, player.ProjectGUID)

		player, err = q.GetPlayerBy("404", testProject.gameGUID)
		assert.Nil(t, player)
		assert.True(t, errors.Is(err, query.ErrProjectNotFound))

		player, err = q.GetPlayerBy(testProject.projectGUID, "404")
		assert.Nil(t, player)
		assert.True(t, errors.Is(err, query.ErrPlayerNotFound))
	}
}

func checkQueryGetMatchBy(t *testing.T, q *query.Service, testProject basicProject) {
	match, err := q.GetMatchBy(testProject.projectGUID, testProject.gameGUID, testProject.matchGUID)
	assert.NotNil(t, match)
	assert.Nil(t, err)
	assert.Equal(t, testProject.gameGUID, match.GameGUID)
	assert.Equal(t, testProject.projectGUID, match.ProjectGUID)
	assert.Len(t, match.Teams, 2)
	assert.Len(t, match.Teams[0].Players, 1)
	assert.Len(t, match.Teams[1].Players, 1)

	assert.Equal(t, 1, match.Teams[0].Score)
	assert.Equal(t, query.Loss, match.Teams[0].Result)
	assert.Equal(t, -16, match.Teams[0].RatingDelta)

	assert.Equal(t, 2, match.Teams[1].Score)
	assert.Equal(t, query.Win, match.Teams[1].Result)
	assert.Equal(t, 16, match.Teams[1].RatingDelta)

	match, err = q.GetMatchBy("404", testProject.gameGUID, testProject.matchGUID)
	assert.Nil(t, match)
	assert.True(t, errors.Is(err, query.ErrProjectNotFound))

	match, err = q.GetMatchBy(testProject.projectGUID, "404", testProject.matchGUID)
	assert.Nil(t, match)
	assert.True(t, errors.Is(err, query.ErrGameNotFound))

	match, err = q.GetMatchBy(testProject.projectGUID, testProject.gameGUID, "404")
	assert.Nil(t, match)
	assert.True(t, errors.Is(err, query.ErrMatchNotFound))
}

func checkQueryGetRatingBy(t *testing.T, q *query.Service, testProject basicProject) {
	for i, p := range testProject.players {
		rating, err := q.GetRatingBy(testProject.projectGUID, p.guid, testProject.gameGUID)
		assert.NotNil(t, rating)
		assert.Nil(t, err)

		assert.Equal(t, p.guid, rating.PlayerGUID)

		if i == 0 {
			assert.Equal(t, 1484, rating.Current)
		} else if i == 1 {
			assert.Equal(t, 1516, rating.Current)
		}

		rating, err = q.GetRatingBy("404", p.guid, testProject.gameGUID)
		assert.Nil(t, rating)
		assert.True(t, errors.Is(err, query.ErrProjectNotFound))

		rating, err = q.GetRatingBy(testProject.projectGUID, "404", testProject.gameGUID)
		assert.Nil(t, rating)
		assert.True(t, errors.Is(err, query.ErrPlayerNotFound))
	}
}
