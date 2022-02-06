package command

import (
	"compelo/event"
	"time"

	"github.com/google/uuid"
)

type CreateNewMatchCommand struct {
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`
	Teams       []struct {
		PlayerGUIDs []string
		Score       int
	} `json:"teams"`
}

func (svc *Service) CreateNewMatch(cmd CreateNewMatchCommand) (Response, error) {
	svc.Lock()
	defer svc.Unlock()

	// TODO: validate event

	guid := uuid.New().String()
	svc.raise(&event.MatchCreated{
		GUID:        guid,
		GameGUID:    cmd.GameGUID,
		ProjectGUID: cmd.ProjectGUID,
		Date:        time.Now(),
		Teams:       cmd.Teams,
	})
	return Response{GUID: guid}, nil
}

type DeleteMatchCommand struct {
	ProjectGUID string `json:"projectGuid"`
	GameGUID    string `json:"gameGuid"`
	GUID        string `json:"matchGuid"`
}

func (svc *Service) DeleteMatch(cmd DeleteMatchCommand) (Response, error) {
	svc.Lock()
	defer svc.Unlock()

	svc.raise(&event.MatchDeleted{
		ProjectGUID: cmd.ProjectGUID,
		GameGUID:    cmd.GameGUID,
		GUID:        cmd.GUID,
	})
	return Response{GUID: cmd.GUID}, nil
}
