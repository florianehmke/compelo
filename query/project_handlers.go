package query

import "compelo/event"

func (c *Compelo) handleProjectCreated(e *event.ProjectCreated) {
	c.projects[e.GUID] = &Project{
		MetaData: MetaData{
			CreatedDate: e.Date,
			UpdatedDate: e.Date,
		},
		GUID:         e.GUID,
		Name:         e.Name,
		PasswordHash: e.PasswordHash,
		games:        make(map[string]*Game),
		players:      make(map[string]*Player),
	}
}
