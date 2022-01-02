package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"compelo/event"
	"compelo/scheduling"
)

var ErrAtLeastTwoTeams = errors.New("at least two teams required for competition")
var ErrAtLeastOneRound = errors.New("at one round required for competition")

type CreateNewCompetitionCommand struct {
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`
	Name        string `json:"name"`
	Rounds      int    `json:"rounds"`
	Teams       []struct {
		PlayerGUIDs []string
	} `json:"teams"`
}

func (c *Service) CreateNewCompetition(cmd CreateNewCompetitionCommand) (Response, error) {
	c.Lock()
	defer c.Unlock()

	if cmd.Rounds < 1 {
		return Response{}, ErrAtLeastOneRound
	}
	if len(cmd.Teams) < 2 {
		return Response{}, ErrAtLeastTwoTeams
	}
	if err := c.checkUniqueConstraint(cmd.ProjectGUID + ":" + cmd.Name); err != nil {
		return Response{}, fmt.Errorf("schedule name is taken: %w", err)
	}

	ev := &event.CompetitionCreated{
		GUID:        uuid.New().String(),
		GameGUID:    cmd.GameGUID,
		ProjectGUID: cmd.ProjectGUID,
		Date:        time.Now(),
		Name:        cmd.Name,
		Rounds:      cmd.Rounds,
		Teams:       cmd.Teams,
	}
	c.raise(ev)
	createMatchCommands(ev)

	return Response{GUID: ev.GUID}, nil
}

func createMatchCommands(competitionEvent *event.CompetitionCreated) {
	s := scheduling.NewRoundRobin()
	for i := range competitionEvent.Teams {
		s.AddPlayer(i)
	}
	schedules := s.ScheduleRounds(competitionEvent.Rounds)

	for competitionRound, round := range schedules {
		for competitionDay, pairings := range round {
			for _, pairing := range pairings {
				ev := event.MatchCreated{
					GUID:             uuid.New().String(),
					CompetitionGUID:  competitionEvent.GUID,
					GameGUID:         competitionEvent.GameGUID,
					ProjectGUID:      competitionEvent.ProjectGUID,
					CompetitionRound: competitionRound + 1,
					CompetitionDay:   competitionDay + 1,
				}
				msg, _ := json.MarshalIndent(ev, "", "  ")
				log.Println(string(msg))

				msg, _ = json.MarshalIndent(competitionEvent.Teams[pairing.A], "", "  ")
				log.Println(string(msg))

				msg, _ = json.MarshalIndent(competitionEvent.Teams[pairing.B], "", "  ")
				log.Println(string(msg))

				// TODO: add teams
			}

		}
	}

}
