package query

import "compelo/event"

func (h *defaultHandler) handlePlayerCreated(e *event.PlayerCreated) {
	h.data.projects[e.ProjectGUID].players[e.GUID] = &Player{
		MetaData: MetaData{
			ID:          e.GetID(),
			CreatedDate: e.Date,
			UpdatedDate: e.Date,
		},
		GUID:        e.GUID,
		ProjectGUID: e.ProjectGUID,
		Name:        e.Name,
		ratings:     make(map[string]*Rating),
	}
}
