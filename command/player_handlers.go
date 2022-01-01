package command

import "compelo/event"

func (svc *Service) handlePlayerCreated(e *event.PlayerCreated) {
	svc.data.projects[e.ProjectGUID].players[e.GUID] = player{
		guid:        e.GUID,
		projectGUID: e.ProjectGUID,
		name:        e.Name,
	}
}
