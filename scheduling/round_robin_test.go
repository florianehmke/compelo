package scheduling_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"compelo/scheduling"
)

func TestRoundRobinEven(t *testing.T) {
	rr := scheduling.NewRoundRobin()
	rr.AddPlayer(1)
	rr.AddPlayer(2)
	rr.AddPlayer(3)
	rr.AddPlayer(4)
	rr.AddPlayer(5)
	rr.AddPlayer(6)

	pairings := rr.Schedule()
	assert.NotNil(t, pairings)

	assert.Equal(t, pairings[0][0].A, 1)
	assert.Equal(t, pairings[0][0].B, 6)
	assert.Equal(t, pairings[1][0].A, 5)
	assert.Equal(t, pairings[1][0].B, 1)
}

func TestRoundRobinOdd(t *testing.T) {
	rr := scheduling.NewRoundRobin()
	rr.AddPlayer(1)
	rr.AddPlayer(2)
	rr.AddPlayer(3)
	rr.AddPlayer(4)
	rr.AddPlayer(5)

	pairings := rr.Schedule()
	assert.NotNil(t, pairings)

	assert.Equal(t, pairings[0][0].A, 2)
	assert.Equal(t, pairings[0][0].B, 5)
	assert.Equal(t, pairings[1][0].A, 5)
	assert.Equal(t, pairings[1][0].B, 1)
}

func TestRoundRobinRounds(t *testing.T) {
	rr := scheduling.NewRoundRobin()
	rr.AddPlayer(1)
	rr.AddPlayer(2)
	rr.AddPlayer(3)
	rr.AddPlayer(4)
	rr.AddPlayer(5)
	rr.AddPlayer(6)

	pairings := rr.ScheduleRounds(4)
	assert.NotNil(t, pairings)

	assert.Equal(t, pairings[0][0][0].A, 1)
	assert.Equal(t, pairings[0][0][0].B, 6)
	assert.Equal(t, pairings[0][1][0].A, 5)
	assert.Equal(t, pairings[0][1][0].B, 1)

	// in round 2 the sides are swapped
	assert.Equal(t, pairings[1][0][0].A, 6)
	assert.Equal(t, pairings[1][0][0].B, 1)
	assert.Equal(t, pairings[1][1][0].A, 1)
	assert.Equal(t, pairings[1][1][0].B, 5)
}
