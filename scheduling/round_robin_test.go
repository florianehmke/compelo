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
}
