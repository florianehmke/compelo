package command

import "compelo/event"

func (svc *Service) handleProjectCreated(e *event.ProjectCreated) {
	svc.data.projects[e.GUID] = project{
		guid:         e.GUID,
		name:         e.Name,
		passwordHash: e.PasswordHash,
		games:        make(map[string]game),
		players:      make(map[string]player),
	}
}
