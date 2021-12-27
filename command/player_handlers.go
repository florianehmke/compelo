package command

import "compelo/event"

func (c *Compelo) handlePlayerCreated(e *event.PlayerCreated) {
	c.projects[e.ProjectGUID].players[e.GUID] = player{
		guid:        e.GUID,
		projectGUID: e.ProjectGUID,
		name:        e.Name,
	}
}
