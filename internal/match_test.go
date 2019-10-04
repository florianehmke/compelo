package compelo_test

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"compelo/internal"
	"compelo/internal/db"
)

type basicProject struct {
	svc *compelo.Service

	project db.Project
	game    db.Game
	player1 db.Player
	player2 db.Player
	player3 db.Player
}

func (bp basicProject) create(t *testing.T) basicProject {
	var err error
	bp.svc = compelo.NewService("file::memory:")

	bp.project, err = bp.svc.CreateProject("Test Project", "secret")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), bp.project.ID)

	bp.game, err = bp.svc.CreateGame(bp.project.ID, "Test Game")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), bp.game.ID)

	bp.player1, err = bp.svc.CreatePlayer(bp.project.ID, "Test Player 1")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), bp.player1.ID)

	bp.player2, err = bp.svc.CreatePlayer(bp.project.ID, "Test Player 2")
	assert.NoError(t, err)
	assert.Equal(t, uint(2), bp.player2.ID)

	bp.player3, err = bp.svc.CreatePlayer(bp.project.ID, "Test Player 3")
	assert.NoError(t, err)
	assert.Equal(t, uint(3), bp.player3.ID)
	return bp
}

func TestMatch(t *testing.T) {
	bp := basicProject{}.create(t)
	bp.testTwoTeamsRequired(t)
	bp.testPlayerInMultipleTeams(t)
	bp.testSameTeamSizeRequired(t)
}

func (bp basicProject) testTwoTeamsRequired(t *testing.T) {
	p := compelo.CreateMatchParameter{
		GameID: bp.game.ID,
		Date:   time.Now(),
		Teams:  nil,
	}

	_, err := bp.svc.CreateMatch(p)
	assert.Equal(t, compelo.ErrTwoTeamsRequired, errors.Cause(err))
}

func (bp basicProject) testPlayerInMultipleTeams(t *testing.T) {
	p := compelo.CreateMatchParameter{
		GameID: bp.game.ID,
		Date:   time.Now(),
		Teams: []compelo.CreateMatchParameterTeam{
			{
				PlayerIDs: []int{
					int(bp.player1.ID),
				},
				Score: 3,
			},
			{
				PlayerIDs: []int{
					int(bp.player1.ID),
				},
				Score: 3,
			},
		},
	}

	_, err := bp.svc.CreateMatch(p)
	assert.Equal(t, compelo.ErrPlayerInMultipleTeams, errors.Cause(err))
}

func (bp basicProject) testSameTeamSizeRequired(t *testing.T) {
	p := compelo.CreateMatchParameter{
		GameID: bp.game.ID,
		Date:   time.Now(),
		Teams: []compelo.CreateMatchParameterTeam{
			{
				PlayerIDs: []int{
					int(bp.player1.ID),
				},
				Score: 3,
			},
			{
				PlayerIDs: []int{
					int(bp.player1.ID),
					int(bp.player2.ID),
				},
				Score: 3,
			},
		},
	}

	_, err := bp.svc.CreateMatch(p)
	assert.Equal(t, compelo.ErrSameTeamSizeRequired, errors.Cause(err))
}
