package query

import "compelo/event"

func (h *defaultHandler) handleGameCreated(e *event.GameCreated) {
	h.data.projects[e.ProjectGUID].games[e.GUID] = &Game{
		MetaData: MetaData{
			ID:          e.GetID(),
			CreatedDate: e.Date,
			UpdatedDate: e.Date,
		},
		GUID:         e.GUID,
		ProjectGUID:  e.ProjectGUID,
		Name:         e.Name,
		eloMatchList: newEloMatchList(),
		playerStats:  make(map[string]*PlayerStats),
		gameStats:    &GameStats{},
		competitions: make(map[string]*Competition),
	}
}
