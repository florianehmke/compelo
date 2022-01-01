package command

import (
	"compelo/event"
	"fmt"

	"github.com/google/uuid"
)

type CreateNewGameCommand struct {
	ProjectGUID string `json:"projectGuid"`
	Name        string `json:"name"`
}

func (svc *Service) CreateNewGame(cmd CreateNewGameCommand) (Response, error) {
	svc.Lock()
	defer svc.Unlock()

	if err := svc.checkUniqueConstraint(cmd.ProjectGUID + ":" + cmd.Name); err != nil {
		return Response{}, fmt.Errorf("game name is taken: %w", err)
	}

	guid := uuid.New().String()
	svc.raise(&event.GameCreated{
		GUID:        guid,
		ProjectGUID: cmd.ProjectGUID,
		Name:        cmd.Name,
	})
	return Response{GUID: guid}, nil
}
