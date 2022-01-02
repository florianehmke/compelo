package command

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"compelo/event"
	"compelo/scheduling"
)

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

	// TODO: validate event

	if err := c.checkUniqueConstraint(cmd.ProjectGUID + ":" + cmd.Name); err != nil {
		return Response{}, fmt.Errorf("schedule name is taken: %w", err)
	}

	// TODO: add match "index" for competition

	guid := uuid.New().String()
	ev := &event.CompetitionCreated{
		GUID:        guid,
		GameGUID:    cmd.GameGUID,
		ProjectGUID: cmd.ProjectGUID,
		Date:        time.Now(),
		Name:        cmd.Name,
		Rounds:      cmd.Rounds,
		Teams:       cmd.Teams,
	}
	c.raise(ev)
	createMatchCommands(ev)

	return Response{GUID: guid}, nil
}

func createMatchCommands(competitionEvent *event.CompetitionCreated) {
	s := scheduling.NewRoundRobin()
	for i := range competitionEvent.Teams {
		s.AddPlayer(i)
	}
	schedule := s.Schedule()

	for i := 0; i < competitionEvent.Rounds; i++ {
		for _, pairings := range schedule {
			for _, pairing := range pairings {
				ev := event.MatchCreated{
					GUID:            uuid.New().String(),
					CompetitionGUID: competitionEvent.GUID,
					GameGUID:        competitionEvent.GameGUID,
					ProjectGUID:     competitionEvent.ProjectGUID,
				}
				msg, _ := json.MarshalIndent(ev, "", "  ")
				log.Println(string(msg))

				msg, _ = json.MarshalIndent(competitionEvent.Teams[pairing.A], "", "  ")
				log.Println(string(msg))

				msg, _ = json.MarshalIndent(competitionEvent.Teams[pairing.B], "", "  ")
				log.Println(string(msg))

				// TODO: add teams
				// TODO: switch sides based on round odd / even, or move to scheduling?
			}
		}
	}
}
