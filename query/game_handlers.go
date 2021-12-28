package query

import "compelo/event"

func (c *Compelo) handleGameCreated(e *event.GameCreated) {
	c.projects[e.ProjectGUID].games[e.GUID] = &Game{
		MetaData: MetaData{
			CreatedDate: e.Date,
			UpdatedDate: e.Date,
		},
		GUID:        e.GUID,
		ProjectGUID: e.ProjectGUID,
		Name:        e.Name,
		matches:     make(map[string]*Match),
		playerStats: make(map[string]*PlayerStats),
		gameStats:   &GameStats{},
	}
}
