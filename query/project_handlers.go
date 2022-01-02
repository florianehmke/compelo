package query

import "compelo/event"

func (h *defaultHandler) handleProjectCreated(e *event.ProjectCreated) {
	h.data.projects[e.GUID] = &Project{
		MetaData: MetaData{
			ID:          e.GetID(),
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
