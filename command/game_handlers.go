package command

import "compelo/event"

func (svc *Service) handleGameCreated(e *event.GameCreated) {
	svc.data.projects[e.ProjectGUID].games[e.GUID] = game{
		guid:         e.GUID,
		projectGUID:  e.ProjectGUID,
		name:         e.Name,
		matches:      make(map[string]match),
		competitions: make(map[string]competition),
	}
}
